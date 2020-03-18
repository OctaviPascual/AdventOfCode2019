package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type adjustRelativeBase struct {
	firstParameterMode parameterMode
}

func (adjustRelativeBase) opcode() opcode {
	return adjustRelativeBaseOpcode
}

func (a adjustRelativeBase) Execute(program *program.Program) error {
	firstParameter, err := getFirstParameter(a.firstParameterMode, program)
	if err != nil {
		return fmt.Errorf("could not get first parameter: %w", err)
	}

	program.RelativeBase += firstParameter

	program.InstructionPointer += 2
	return nil
}
