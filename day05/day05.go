package day05

import (
	"errors"
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
)

const (
	airConditionerUnitID        = 1
	thermalRadiatorControllerID = 5
)

// Day holds the data needed to solve part one and part two
type Day struct {
	program string
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	return &Day{
		program: input,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	intcodeProgram, err := intcode.NewIntcodeProgram(d.program)
	if err != nil {
		return "", err
	}

	outputs, err := intcodeProgram.RunWithInput(airConditionerUnitID)
	if err != nil {
		return "", err
	}

	for i, output := range outputs {
		if i < len(outputs)-1 && output != intcode.Output(0) {
			return "", errors.New("all outputs but last should be equal to 0")
		}
	}

	return fmt.Sprintf("%d", outputs[len(outputs)-1]), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	intcodeProgram, err := intcode.NewIntcodeProgram(d.program)
	if err != nil {
		return "", err
	}

	outputs, err := intcodeProgram.RunWithInput(thermalRadiatorControllerID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", outputs[len(outputs)-1]), nil
}
