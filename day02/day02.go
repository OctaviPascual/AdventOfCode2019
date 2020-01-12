package day02

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	Output = 0
	Noun   = 1
	Verb   = 2

	DesiredIntcodeOutput = intcodeOutput(19690720)
)

type intcodeOutput int

type intcodeProgram struct {
	memory []int
}

type Day struct {
	initialState   []int
	intcodeProgram intcodeProgram
}

func NewDay(input string) (*Day, error) {
	values := strings.Split(input, ",")
	initialState := make([]int, 0, len(values))
	for _, value := range values {
		val, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("invalid value %s: %v", value, err)
		}
		initialState = append(initialState, val)
	}
	return &Day{
		initialState: initialState,
		intcodeProgram: intcodeProgram{
			memory: make([]int, len(initialState)),
		},
	}, nil
}

func (i intcodeProgram) run() error {
	instructionPointer := 0
	for {
		instruction := i.memory[instructionPointer]
		if instruction == 1 {
			address1 := i.memory[instructionPointer+1]
			address2 := i.memory[instructionPointer+2]
			address3 := i.memory[instructionPointer+3]
			i.memory[address3] = i.memory[address1] + i.memory[address2]
			instructionPointer += 4
		} else if instruction == 2 {
			address1 := i.memory[instructionPointer+1]
			address2 := i.memory[instructionPointer+2]
			address3 := i.memory[instructionPointer+3]
			i.memory[address3] = i.memory[address1] * i.memory[address2]
			instructionPointer += 4
		} else if instruction == 99 {
			return nil
		} else {
			return fmt.Errorf("found unknown instruction %d", instruction)
		}
	}
}

func (d Day) SolvePartOne() (string, error) {
	intcodeOutput, err := d.runIntcodeProgram(12, 2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", intcodeOutput), nil
}

func (d Day) SolvePartTwo() (string, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			intcodeOutput, err := d.runIntcodeProgram(noun, verb)
			if err != nil {
				return "", err
			}
			if intcodeOutput == DesiredIntcodeOutput {
				return fmt.Sprintf("%d", 100*noun+verb), nil
			}
		}
	}
	return "", errors.New("could not find combination to produce desired output")
}

func (d Day) runIntcodeProgram(noun, verb int) (intcodeOutput, error) {
	copy(d.intcodeProgram.memory, d.initialState)
	d.intcodeProgram.memory[Noun] = noun
	d.intcodeProgram.memory[Verb] = verb
	err := d.intcodeProgram.run()
	if err != nil {
		return intcodeOutput(0), err
	}
	return intcodeOutput(d.intcodeProgram.memory[Output]), nil
}
