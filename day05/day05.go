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

	inputChannel := make(chan int, 1)
	inputChannel <- airConditionerUnitID
	outputChannel := make(chan int, 1000)

	err = intcodeProgram.Run(inputChannel, outputChannel)
	if err != nil {
		return "", err
	}

	var outputs []int
	for output := range outputChannel {
		outputs = append(outputs, output)
	}

	for i, output := range outputs {
		if i < len(outputs)-1 && output != 0 {
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

	inputChannel := make(chan int, 1)
	inputChannel <- thermalRadiatorControllerID
	outputChannel := make(chan int, 1)

	err = intcodeProgram.Run(inputChannel, outputChannel)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", <-outputChannel), nil
}
