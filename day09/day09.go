package day09

import (
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2019/intcode"
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
	return runWithInput(d.program, 1)
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return runWithInput(d.program, 2)
}

func runWithInput(program string, input int) (string, error) {
	var outputs []string
	onOutput := func(output int) {
		outputs = append(outputs, strconv.Itoa(output))
	}

	intcodeProgram, err := intcode.NewIntcodeProgram(
		program, func() int { return input }, onOutput,
	)
	if err != nil {
		return "", err
	}

	err = intcodeProgram.Run()
	if err != nil {
		return "", err
	}

	return strings.Join(outputs, ","), nil
}
