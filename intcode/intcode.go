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

type opcode uint8

const (
	AddInstruction      opcode = 1
	MultiplyInstruction opcode = 2
	InputInstruction    opcode = 3
	OutputInstruction   opcode = 4
	HaltInstruction     opcode = 99
)

type instruction struct {
	opcode              opcode
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
	thirdParameterMode  parameterMode
}

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
		instruction := parseIntruction(p.memory[instructionPointer])
		switch instruction.opcode {
		case AddInstruction:
			var firstParameter int
			if instruction.firstParameterMode == positionMode {
				address1 := p.memory[instructionPointer+1]
				firstParameter = p.memory[address1]
			} else {
				firstParameter = p.memory[instructionPointer+1]
			}

			var secondParameter int
			if instruction.firstParameterMode == positionMode {
				address2 := p.memory[instructionPointer+2]
				secondParameter = p.memory[address2]
			} else {
				secondParameter = p.memory[instructionPointer+2]
			}

			address3 := p.memory[instructionPointer+3]
			p.memory[address3] = firstParameter + secondParameter
			instructionPointer += 4
		case MultiplyInstruction:
			var firstParameter int
			if instruction.firstParameterMode == positionMode {
				address1 := p.memory[instructionPointer+1]
				firstParameter = p.memory[address1]
			} else {
				firstParameter = p.memory[instructionPointer+1]
			}

			var secondParameter int
			if instruction.firstParameterMode == positionMode {
				address2 := p.memory[instructionPointer+2]
				secondParameter = p.memory[address2]
			} else {
				secondParameter = p.memory[instructionPointer+2]
			}

			address3 := p.memory[instructionPointer+3]
			p.memory[address3] = firstParameter * secondParameter
			instructionPointer += 4
		case InputInstruction:
			address := p.memory[instructionPointer+1]
			p.memory[address] = p.input
			instructionPointer += 2
		case OutputInstruction:
			address := p.memory[instructionPointer+1]
			p.outputBuffer = append(p.outputBuffer, p.memory[address])
			instructionPointer += 2
		case HaltInstruction:
			return nil
		default:
			return fmt.Errorf("found unknown instruction opcode %d", instruction.opcode)
		}
	}
}

func parseIntruction(n int) instruction {
	return instruction{
		opcode:              opcode(n % 100),
		firstParameterMode:  parameterMode((n / 100) % 10),
		secondParameterMode: parameterMode((n / 1000) % 10),
		thirdParameterMode:  parameterMode((n / 10000) % 10),
	}
}