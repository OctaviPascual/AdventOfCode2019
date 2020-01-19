package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	outputPosition = 0
	nounPosition   = 1
	verbPosition   = 2
)

// Program represents an Intcode program
type Program struct {
	input        int
	outputBuffer []int

	memory []int
}

// Output is the output of an Intcode program
type Output int

type parameterMode int

const (
	positionMode  parameterMode = 0
	immediateMode parameterMode = 1
)

// NewIntcodeProgram creates a new Intcode Program
func NewIntcodeProgram(program string) (*Program, error) {
	memory, err := parse(program)
	if err != nil {
		return nil, err
	}

	return &Program{
		memory: memory,
	}, nil
}

func parse(program string) ([]int, error) {
	tokens := strings.Split(program, ",")
	values := make([]int, 0, len(tokens))

	for _, token := range tokens {

		value, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("invalid value %s: %w", token, err)
		}

		values = append(values, value)
	}

	return values, nil
}

// Run runs an Intcode program with the given noun and verb
func (p *Program) Run(noun, verb int) (Output, error) {
	p.memory[nounPosition] = noun
	p.memory[verbPosition] = verb

	err := p.run()
	if err != nil {
		return Output(0), fmt.Errorf("error running intcode program: %w", err)
	}

	return Output(p.memory[outputPosition]), nil
}

func (p *Program) run() error {
	instructionPointer := 0
	for {
		instruction := p.memory[instructionPointer]
		if instruction == 1 {
			address1 := p.memory[instructionPointer+1]
			address2 := p.memory[instructionPointer+2]
			address3 := p.memory[instructionPointer+3]
			p.memory[address3] = p.memory[address1] + p.memory[address2]
			instructionPointer += 4
		} else if instruction == 2 {
			address1 := p.memory[instructionPointer+1]
			address2 := p.memory[instructionPointer+2]
			address3 := p.memory[instructionPointer+3]
			p.memory[address3] = p.memory[address1] * p.memory[address2]
			instructionPointer += 4
		} else if instruction == 3 {
			p.memory[instructionPointer+1] = p.input
			instructionPointer += 2
		} else if instruction == 4 {
			p.outputBuffer = append(p.outputBuffer, p.memory[instructionPointer+1])
			instructionPointer += 2
		} else if instruction == 99 {
			return nil
		} else {
			return fmt.Errorf("found unknown instruction %d", instruction)
		}
	}
}
