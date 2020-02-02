package intcode

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/instruction"
	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

const (
	outputPosition = 0
	nounPosition   = 1
	verbPosition   = 2
)

// Intcode represents an Intcode program
type Intcode struct {
	program *program.Program
}

// Output is the output of an Intcode program
type Output int

// NewIntcodeProgram creates a new Intcode Program
func NewIntcodeProgram(programString string) (*Intcode, error) {
	p, err := program.NewProgram(programString)
	if err != nil {
		return nil, fmt.Errorf("error creating program: %w", err)
	}

	return &Intcode{
		program: p,
	}, nil
}

// RunWithNounAndVerb runs an Intcode program with the given noun and verb
func (i *Intcode) RunWithNounAndVerb(noun, verb int) (Output, error) {
	err := i.program.Store(nounPosition, noun)
	if err != nil {
		return Output(0), fmt.Errorf("error setting noun: %w", err)
	}

	err = i.program.Store(verbPosition, verb)
	if err != nil {
		return Output(0), fmt.Errorf("error setting verb: %w", err)
	}

	err = run(i.program)
	if err != nil {
		return Output(0), fmt.Errorf("runtime error: %w", err)
	}

	output, err := i.program.Fetch(outputPosition)
	return Output(output), err
}

// RunWithInput runs an Intcode program with the given input
func (i *Intcode) RunWithInput(input int) ([]Output, error) {
	i.program.SetInput([]int{input})

	err := run(i.program)
	if err != nil {
		return nil, fmt.Errorf("runtime error: %w", err)
	}

	var outputs []Output
	for _, output := range i.program.GetOutput() {
		outputs = append(outputs, Output(output))
	}
	return outputs, nil
}

func run(program *program.Program) error {
	for !program.Halted {

		n, err := program.Fetch(program.InstructionPointer)
		if err != nil {
			return fmt.Errorf("error fetching instruction: %w", err)
		}

		parsedInstruction, err := instruction.ParseInstruction(n)
		if err != nil {
			return fmt.Errorf("error parsing instruction: %w", err)
		}

		err = parsedInstruction.Execute(program)
		if err != nil {
			return fmt.Errorf("error executing instruction: %w", err)
		}
	}
	return nil
}
