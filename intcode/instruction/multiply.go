package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type multiply struct {
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
}

func (multiply) opcode() opcode {
	return multiplyOpcode
}

func (m multiply) Execute(program *program.Program) error {
	firstParameter, err := getFirstParameter(m.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get first parameter: %w", err)
	}

	secondParameter, err := getSecondParameter(m.secondParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get second parameter: %w", err)
	}

	address, err := program.Fetch(program.InstructionPointer + 3)
	if err != nil {
		return fmt.Errorf("could not get third parameter: %w", err)
	}

	err = program.Store(address, firstParameter*secondParameter)
	if err != nil {
		return err
	}

	program.InstructionPointer += 4
	return nil
}
