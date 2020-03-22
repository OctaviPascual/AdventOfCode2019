package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/util"
)

const (
	io       = "Io"
	europa   = "Europa"
	ganymede = "Ganymede"
	callisto = "Callisto"

	xDimension = 0
	yDimension = 1
	zDimension = 2
)

var (
	stepsToSimulate = 1000

	moonNames = []string{io, europa, ganymede, callisto}
	moonRe    = regexp.MustCompile(`^<x=(-?\d+), y=(-?\d+), z=(-?\d+)>$`)

	dimensions = []int{xDimension, yDimension, zDimension}
)

// Day holds the data needed to solve part one and part two
type Day struct {
	moons []*moon
}

type moon struct {
	name     string
	position position
	velocity velocity
}

type position [3]int
type velocity [3]int

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	moonsStrings := strings.Split(input, "\n")

	moons, err := parseMoons(moonsStrings)
	if err != nil {
		return nil, err
	}

	return &Day{
		moons: moons,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	simulateMotion(d.moons, stepsToSimulate, dimensions...)
	return fmt.Sprintf("%d", totalEnergy(d.moons)), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	xCycleLength := findCycleLength(d.moons, xDimension)
	yCycleLength := findCycleLength(d.moons, yDimension)
	zCycleLength := findCycleLength(d.moons, zDimension)
	return fmt.Sprintf("%d", util.LCM3(xCycleLength, yCycleLength, zCycleLength)), nil
}

func parseMoons(moonsStrings []string) ([]*moon, error) {
	if len(moonsStrings) != 4 {
		return nil, fmt.Errorf("invalid number of moons: %d", len(moonsStrings))
	}

	moons := make([]*moon, 0, 4)
	for i, moonName := range moonNames {
		moon, err := parseMoon(moonsStrings[i], moonName)
		if err != nil {
			return nil, fmt.Errorf("could not parse moon %s: %w", moonName, err)
		}
		moons = append(moons, moon)
	}
	return moons, nil

}

func parseMoon(moonString, moonName string) (*moon, error) {
	matches := moonRe.FindStringSubmatch(moonString)
	if len(matches) != 4 {
		return nil, fmt.Errorf("invalid format of moon string: %s", moonString)
	}

	x, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, fmt.Errorf("invalid x position: %w", err)
	}

	y, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("invalid y position: %w", err)
	}

	z, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid z position: %w", err)
	}

	return &moon{
		name:     moonName,
		position: position{x, y, z},
	}, nil
}

func simulateMotion(moons []*moon, steps int, dimensions ...int) {
	for i := 0; i < steps; i++ {
		for _, dimension := range dimensions {
			applyTimeStep(moons, dimension)
		}
	}
}

func applyTimeStep(moons []*moon, dimension int) {
	applyGravity(moons, dimension)
	applyVelocity(moons, dimension)
}

func applyGravity(moons []*moon, dimension int) {
	for i := 0; i < len(moons); i++ {
		for j := i + 1; j < len(moons); j++ {
			updateVelocities(moons[i], moons[j], dimension)
		}
	}
}

func updateVelocities(moon1, moon2 *moon, dimension int) {
	if moon1.position[dimension] < moon2.position[dimension] {
		moon1.velocity[dimension]++
		moon2.velocity[dimension]--
	}
	if moon1.position[dimension] > moon2.position[dimension] {
		moon1.velocity[dimension]--
		moon2.velocity[dimension]++
	}
}

func applyVelocity(moons []*moon, dimension int) {
	for _, moon := range moons {
		moon.updatePosition(dimension)
	}
}

func (m *moon) updatePosition(dimension int) {
	m.position[dimension] += m.velocity[dimension]
}

func totalEnergy(moons []*moon) int {
	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += moon.totalEnergy()
	}
	return totalEnergy
}

func (m moon) totalEnergy() int {
	return m.potentialEnergy() * m.kineticEnergy()
}

func (m moon) potentialEnergy() int {
	potentialEnergy := 0
	for i := 0; i < len(m.position); i++ {
		potentialEnergy += util.Abs(m.position[i])
	}
	return potentialEnergy
}

func (m moon) kineticEnergy() int {
	kineticEnergy := 0
	for i := 0; i < len(m.position); i++ {
		kineticEnergy += util.Abs(m.velocity[i])
	}
	return kineticEnergy
}

func findCycleLength(moons []*moon, dimension int) int {
	steps := 0
	var initialPositions []int
	var initialVelocities []int
	for _, moon := range moons {
		initialPositions = append(initialPositions, moon.position[dimension])
		initialVelocities = append(initialVelocities, moon.velocity[dimension])
	}

	for {
		applyTimeStep(moons, dimension)
		steps++
		if equalPositionsAndVelocities(moons, dimension, initialPositions, initialVelocities) {
			return steps
		}
	}
}

func equalPositionsAndVelocities(moons []*moon, dimension int, positions, velocities []int) bool {
	for i := 0; i < len(moons); i++ {
		if moons[i].position[dimension] != positions[i] {
			return false
		}
		if moons[i].velocity[dimension] != velocities[i] {
			return false
		}
	}
	return true
}
