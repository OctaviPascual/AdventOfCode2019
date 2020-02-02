package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type output struct {
	firstParameterMode parameterMode
}

func (output) opcode() opcode {
	return outputOpcode
}

func (o output) Execute(program *program.Program) error {
	firstParameter, err := getFirstParameter(o.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get first parameter: %w", err)
	}

	program.WriteOutput(firstParameter)

	program.InstructionPointer += 2
	return nil
}
