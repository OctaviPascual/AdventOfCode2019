package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type jumpIfFalse struct {
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
}

func (jumpIfFalse) opcode() opcode {
	return jumpIfFalseOpcode
}

func (jf jumpIfFalse) Execute(program *program.Program) error {
	firstParameter, err := getFirstParameter(jf.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get first parameter: %w", err)
	}

	if firstParameter == 0 {
		secondParameter, err := getSecondParameter(jf.secondParameterMode, program)
		if err != nil {
			return fmt.Errorf("could not get second parameter: %w", err)
		}
		program.InstructionPointer = secondParameter
		return nil
	}

	program.InstructionPointer += 3
	return nil
}
