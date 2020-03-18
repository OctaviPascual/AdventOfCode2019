package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type lessThan struct {
	firstParameterMode  parameterMode
	secondParameterMode parameterMode
	thirdParameterMode  parameterMode
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

	value := 0
	if firstParameter < secondParameter {
		value = 1
	}

	err = storeWithThirdParameter(value, lt.thirdParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not store with third parameter: %w", err)
	}

	program.InstructionPointer += 4
	return nil
}
