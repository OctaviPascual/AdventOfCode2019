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
	var outputs []int
	onOutput := func(output int) {
		outputs = append(outputs, output)
	}

	intcodeProgram, err := intcode.NewIntcodeProgram(
		d.program, func() int { return airConditionerUnitID }, onOutput,
	)
	if err != nil {
		return "", err
	}

	err = intcodeProgram.Run()
	if err != nil {
		return "", err
	}

	for i, output := range outputs {
		if i < len(outputs)-1 && output != 0 {
			return "", errors.New("all outputs but last should be equal to 0")
		}
	}

	diagnosticCode := outputs[len(outputs)-1]
	return fmt.Sprintf("%d", diagnosticCode), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	var diagnosticCode int
	onOutput := func(output int) {
		diagnosticCode = output
	}

	intcodeProgram, err := intcode.NewIntcodeProgram(
		d.program, func() int { return thermalRadiatorControllerID }, onOutput,
	)
	if err != nil {
		return "", err
	}

	err = intcodeProgram.Run()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", diagnosticCode), nil
}
