package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type lessThan struct {
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
}

func (lessThan) opcode() opcode {
	return lessThanOpcode
}

func (lt lessThan) Execute(program *program.Program) error {
	firstParameter, err := getFirstParameter(lt.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get first parameter: %w", err)
	}

	secondParameter, err := getSecondParameter(lt.secondParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get second parameter: %w", err)
	}

	address, err := program.Fetch(program.InstructionPointer + 3)
	if err != nil {
		return fmt.Errorf("could not get third parameter: %w", err)
	}

	value := 0
	if firstParameter < secondParameter {
		value = 1
	}

	err = program.Store(address, value)
	if err != nil {
		return nil
	}

	program.InstructionPointer += 4
	return nil
}
