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
	intcodeProgram, err := intcode.NewIntcodeProgram(program)
	if err != nil {
		return "", err
	}

	inputChannel := make(chan int, 1)
	inputChannel <- input
	outputChannel := make(chan int, 1)
	errorChannel := make(chan error, 1)
	go func() {
		errorChannel <- intcodeProgram.Run(inputChannel, outputChannel)
	}()

	var outputs []string
	for output := range outputChannel {
		outputs = append(outputs, strconv.Itoa(output))
	}

	err = <-errorChannel
	if err != nil {
		return "", err
	}

	return strings.Join(outputs, ","), nil
}
