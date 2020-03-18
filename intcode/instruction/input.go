package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type input struct {
	firstParameterMode parameterMode
}

func (input) opcode() opcode {
	return inputOpcode
}

func (i input) Execute(program *program.Program) error {
	value := program.ReadInput()

	err := storeWithFirstParameter(value, i.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not store with first parameter: %w", err)
	}

	program.InstructionPointer += 2
	return nil
}
