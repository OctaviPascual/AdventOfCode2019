package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type jumpIfTrue struct {
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
}

func (jumpIfTrue) opcode() opcode {
	return jumpIfTrueOpcode
}

func (jt jumpIfTrue) Execute(program *program.Program) error {
	firstParameter, err := getFirstParameter(jt.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get first parameter: %w", err)
	}

	if firstParameter != 0 {
		secondParameter, err := getSecondParameter(jt.secondParameterMode, program)
		if err != nil {
			return fmt.Errorf("could not get second parameter: %w", err)
		}
		program.InstructionPointer = secondParameter
		return nil
	}

	program.InstructionPointer += 3
	return nil
}
