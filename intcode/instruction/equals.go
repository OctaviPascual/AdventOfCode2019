package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type equals struct {
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
}

func (equals) opcode() opcode {
	return equalsOpcode
}

func (eq equals) Execute(program *program.Program) error {
	firstParameter, err := getFirstParameter(eq.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get first parameter: %w", err)
	}

	secondParameter, err := getSecondParameter(eq.secondParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get second parameter: %w", err)
	}

	address, err := program.Fetch(program.InstructionPointer + 3)
	if err != nil {
		return fmt.Errorf("could not get third parameter: %w", err)
	}

	value := 0
	if firstParameter == secondParameter {
		value = 1
	}

	err = program.Store(address, value)
	if err != nil {
		return nil
	}

	program.InstructionPointer += 4
	return nil
}
