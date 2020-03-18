package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type multiply struct {
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
	thirdParameterMode  parameterMode
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

	err = storeWithThirdParameter(firstParameter*secondParameter, m.thirdParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not store with third parameter: %w", err)
	}

	program.InstructionPointer += 4
	return nil
}
