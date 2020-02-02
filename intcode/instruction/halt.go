package instruction

import (
	"errors"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type halt struct{}

func (halt) opcode() opcode {
	return haltOpcode
}

func (halt) Execute(program *program.Program) error {
	if program.Halted {
		return errors.New("cannot halt an already halted program")
	}

	program.Halted = true
	return nil
}
