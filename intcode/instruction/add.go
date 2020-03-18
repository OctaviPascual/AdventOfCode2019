package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type add struct {
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
	thirdParameterMode  parameterMode
}

func (add) opcode() opcode {
	return addOpcode
}

func (a add) Execute(program *program.Program) error {
	firstParameter, err := getFirstParameter(a.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get first parameter: %w", err)
	}

	secondParameter, err := getSecondParameter(a.secondParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get second parameter: %w", err)
	}

	err = storeWithThirdParameter(firstParameter+secondParameter, a.thirdParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not store with third parameter: %w", err)
	}

	program.InstructionPointer += 4
	return nil
}
