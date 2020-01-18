package day02

import (
	"errors"
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
)

const (
	desiredIntcodeOutput = intcode.Output(19690720)
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

	intcodeOutput, err := intcodeProgram.Run(12, 2)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", intcodeOutput), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {

			intcodeProgram, err := intcode.NewIntcodeProgram(d.program)
			if err != nil {
				return "", err
			}

			intcodeOutput, err := intcodeProgram.Run(noun, verb)
			if err != nil {
				return "", err
			}

			if intcodeOutput == desiredIntcodeOutput {
				return fmt.Sprintf("%d", 100*noun+verb), nil
			}
		}
	}

	return "", errors.New("could not find combination to produce desired output")
}
