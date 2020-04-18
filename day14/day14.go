package day14

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	reactions []reaction
}

type reaction struct {
	reactants []balancedChemical
	product   balancedChemical
}

type balancedChemical struct {
	chemical    chemical
	coefficient int
}

type chemical string

type state struct {
	reactionByChemical map[chemical]reaction
	leftoverByChemical map[chemical]int
}

const (
	rawMaterial chemical = "ORE"
	fuel        chemical = "FUEL"
)

var (
	// Regex matching a reaction of the form 71 AA, 5 BB, 2 CC => 11 DD
	reactionRe = regexp.MustCompile(`^(\d+ \w+(?:, \d+ \w+)*) => (\d+ \w+)$`)
	// Regex matching a balanced chemical of the form 42 AA
	chemicalRe = regexp.MustCompile(`(\d+) (\w+)`)
)

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	reactionsString := strings.Split(input, "\n")

	reactions, err := parseReactions(reactionsString)
	if err != nil {
		return nil, err
	}

	return &Day{
		reactions: reactions,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	state := state{
		reactionByChemical: createReactionByChemical(d.reactions),
		leftoverByChemical: make(map[chemical]int),
	}

	rawMaterialRequired, err := getRawMaterialRequired(fuel, 1, state)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", rawMaterialRequired), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	state := state{
		reactionByChemical: createReactionByChemical(d.reactions),
		leftoverByChemical: make(map[chemical]int),
	}

	maximumAmountOfFuel, err := getMaximumAmountOfFuel(1e12, state)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", maximumAmountOfFuel), nil
}

func parseReactions(reactionsString []string) ([]reaction, error) {
	reactions := make([]reaction, 0, len(reactionsString))

	for _, reactionString := range reactionsString {
		reaction, err := parseReaction(reactionString)
		if err != nil {
			return nil, err
		}

		reactions = append(reactions, reaction)
	}

	return reactions, nil
}

func parseReaction(reactionString string) (reaction, error) {
	matches := reactionRe.FindStringSubmatch(reactionString)
	if len(matches) != 3 {
		return reaction{}, fmt.Errorf("invalid format of reaction string: %s", reactionString)
	}
	// From now on, we know that the reaction has a valid format
	// We won't check for errors in the subsequent parsing steps as we know they can't occur
	return reaction{
		reactants: parseReactants(matches[1]),
		product:   parseProduct(matches[2]),
	}, nil
}

func parseReactants(reactantsString string) []balancedChemical {
	matches := chemicalRe.FindAllStringSubmatch(reactantsString, -1)

	reactants := make([]balancedChemical, 0, len(matches))
	for _, match := range matches {
		reactant := parseChemical(match)
		reactants = append(reactants, reactant)
	}
	return reactants
}

func parseProduct(productString string) balancedChemical {
	match := chemicalRe.FindStringSubmatch(productString)
	return parseChemical(match)
}

func parseChemical(match []string) balancedChemical {
	coefficient, _ := strconv.Atoi(match[1])
	return balancedChemical{
		chemical:    chemical(match[2]),
		coefficient: coefficient,
	}
}

func getRawMaterialRequired(chemical chemical, quantity int, state state) (int, error) {
	err := produce(chemical, quantity, state)
	if err != nil {
		return 0, fmt.Errorf("could not produce %d quantity of chemical %s: %w", quantity, chemical, err)
	}

	return state.leftoverByChemical[rawMaterial], nil
}

func createReactionByChemical(reactions []reaction) map[chemical]reaction {
	reactionByChemical := make(map[chemical]reaction, len(reactions))
	for _, reaction := range reactions {
		reactionByChemical[reaction.product.chemical] = reaction
	}
	return reactionByChemical
}

func produce(chemical chemical, quantity int, state state) error {
	if chemical == rawMaterial {
		state.leftoverByChemical[rawMaterial] += quantity
		return nil
	}

	reaction, ok := state.reactionByChemical[chemical]
	if !ok {
		return fmt.Errorf("there is no reaction to produce the chemical %s", chemical)
	}

	quantityNeeded := useLeftoversIfAny(quantity, chemical, state)
	reactionsNeeded := numberOfReactionsNeeded(quantityNeeded, reaction.product.coefficient)
	quantityProduced := reaction.product.coefficient * reactionsNeeded

	updateLeftovers(quantityProduced-quantityNeeded, chemical, state)

	for _, reactant := range reaction.reactants {
		err := produce(reactant.chemical, reactant.coefficient*reactionsNeeded, state)
		if err != nil {
			return err
		}
	}

	return nil
}

func useLeftoversIfAny(quantity int, chemical chemical, state state) int {
	newQuantity := util.Max(0, quantity-state.leftoverByChemical[chemical])
	leftOversUsed := quantity - newQuantity
	state.leftoverByChemical[chemical] -= leftOversUsed

	return newQuantity
}

func numberOfReactionsNeeded(quantity int, coefficient int) int {
	x := float64(quantity) / float64(coefficient)
	return int(math.Ceil(x))
}

func updateLeftovers(leftovers int, chemical chemical, state state) {
	state.leftoverByChemical[chemical] += leftovers
}

func getMaximumAmountOfFuel(rawMaterialQuantity int, state state) (int, error) {
	rawMaterialRequired, err := getRawMaterialRequired(fuel, 1, state)
	if err != nil {
		return 0, err
	}

	// This value is a lower bound: we will always be able to produce this amount of fuel
	lowAmountOfFuel := rawMaterialQuantity / rawMaterialRequired

	// This value is an upper bound: given that all reaction coefficients are integers
	// and that fuel coefficient is always 1, we will need at least 1 raw material to produce 1 fuel
	// Thus, at most we will be able to produce as much fuel as raw material
	highAmountOfFuel := rawMaterialQuantity

	return searchAmountOfFuel(lowAmountOfFuel, highAmountOfFuel, rawMaterialQuantity, state)
}

func searchAmountOfFuel(lowAmountOfFuel, highAmountOfFuel, rawMaterialQuantity int, state state) (int, error) {
	if lowAmountOfFuel >= highAmountOfFuel {
		return 0, fmt.Errorf("could not find a maximum amount of fuel, review inital low and high bounds")
	}

	amountOfFuel := (lowAmountOfFuel + highAmountOfFuel) / 2

	canProduceCurrent, err := canProduce(amountOfFuel, rawMaterialQuantity, state)
	if err != nil {
		return 0, err
	}

	canProduceNext, err := canProduce(amountOfFuel+1, rawMaterialQuantity, state)
	if err != nil {
		return 0, err
	}

	// We found the maximum amount of fuel, we know we won't manage to produce more fuel with the
	// same amount of raw materials
	if canProduceCurrent && !canProduceNext {
		return amountOfFuel, nil
	}

	if canProduceCurrent {
		return searchAmountOfFuel(amountOfFuel+1, highAmountOfFuel, rawMaterialQuantity, state)
	}

	return searchAmountOfFuel(lowAmountOfFuel, amountOfFuel, rawMaterialQuantity, state)
}

func canProduce(fuelQuantity, rawMaterialQuantity int, state state) (bool, error) {
	state.leftoverByChemical = make(map[chemical]int)

	rawMaterialRequired, err := getRawMaterialRequired(fuel, fuelQuantity, state)
	if err != nil {
		return false, err
	}

	return rawMaterialRequired <= rawMaterialQuantity, nil
}
