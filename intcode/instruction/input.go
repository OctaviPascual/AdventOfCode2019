package instruction

import "github.com/OctaviPascual/AdventOfCode2019/intcode/program"

type input struct{}

func (input) opcode() opcode {
	return inputOpcode
}

func (i input) Execute(program *program.Program) error {
	value, err := program.ReadInput()
	if err != nil {
		return err
	}

	address, err := program.Fetch(program.InstructionPointer + 1)
	if err != nil {
		return err
	}

	err = program.Store(address, value)
	if err != nil {
		return err
	}

	program.InstructionPointer += 2
	return nil
}
