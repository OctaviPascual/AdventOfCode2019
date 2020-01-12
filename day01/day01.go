package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type fuel int

type module struct {
	mass int
}

type spacecraft struct {
	modules []module
}

// Day holds the data needed to solve part one and part two
type Day struct {
	spacecraft spacecraft
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	spacecraftString := strings.Split(input, "\n")

	spacecraft, err := parseSpacecraft(spacecraftString)
	if err != nil {
		return nil, fmt.Errorf("invalid spacecraft %s: %w", spacecraftString, err)
	}

	return &Day{
		spacecraft: *spacecraft,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	fuel := d.spacecraft.totalFuelRequired()
	return fmt.Sprintf("%d", fuel), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	fuel := d.spacecraft.totalFuelRequiredWithAddedFuel()
	return fmt.Sprintf("%d", fuel), nil
}

func parseSpacecraft(spacecraftString []string) (*spacecraft, error) {
	modules := make([]module, 0, len(spacecraftString))

	for _, moduleString := range spacecraftString {

		module, err := parseModule(moduleString)
		if err != nil {
			return nil, fmt.Errorf("invalid module %s: %w", moduleString, err)
		}

		modules = append(modules, *module)
	}

	return &spacecraft{
		modules: modules,
	}, nil
}

func parseModule(moduleString string) (*module, error) {
	mass, err := strconv.Atoi(moduleString)
	if err != nil {
		return nil, fmt.Errorf("invalid mass %s: %w", moduleString, err)
	}

	return &module{
		mass: mass,
	}, nil
}

func (s spacecraft) totalFuelRequired() fuel {
	totalFuelRequired := fuel(0)
	for _, module := range s.modules {
		totalFuelRequired += module.fuelRequired()
	}
	return totalFuelRequired
}

func (m module) fuelRequired() fuel {
	return fuel(m.mass/3 - 2)
}

func (s spacecraft) totalFuelRequiredWithAddedFuel() fuel {
	totalFuelRequired := fuel(0)
	for _, module := range s.modules {
		moduleFuelRequired := module.fuelRequired()
		addedFuelRequired := moduleFuelRequired.fuelRequired()

		totalFuelRequired += moduleFuelRequired + addedFuelRequired
	}
	return totalFuelRequired
}

func (f fuel) fuelRequired() fuel {
	fuelRequired := f/3 - 2
	if fuelRequired <= 0 {
		return 0
	}
	return fuelRequired + fuelRequired.fuelRequired()
}
