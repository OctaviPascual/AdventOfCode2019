package intcode

import (
	"fmt"
	"sync"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/instruction"
	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

const (
	outputPosition = 0
	nounPosition   = 1
	verbPosition   = 2
)

var (
	// MustNotInput panics if Intcode program expects an input
	MustNotInput = func() int {
		panic("intcode program expects an input")
	}

	// MustNotOutput panics if Intcode program produces an output
	MustNotOutput = func(output int) {
		panic(fmt.Sprintf("intcode program produced an output: %d", output))
	}
)

// Intcode represents an Intcode program
type Intcode struct {
	sync.RWMutex

	program    *program.Program
	shouldStop bool
}

// NewIntcodeProgram creates a new Intcode program from the following parameters:
// - programString is the string representation of the program
// - onInput is the function that will be called whenever the program expects an input
// - onOutput is the function that will be called whenever the program produces an output
func NewIntcodeProgram(
	programString string,
	onInput func() int,
	onOutput func(output int),
) (*Intcode, error) {
	p, err := program.NewProgram(programString, onInput, onOutput)
	if err != nil {
		return nil, fmt.Errorf("error creating program: %w", err)
	}

	return &Intcode{
		program: p,
	}, nil
}

// RunWithNounAndVerb runs an Intcode program with the given noun and verb
func (i *Intcode) RunWithNounAndVerb(noun, verb int) (int, error) {
	err := i.program.Store(nounPosition, noun)
	if err != nil {
		return 0, fmt.Errorf("error setting noun: %w", err)
	}

	err = i.program.Store(verbPosition, verb)
	if err != nil {
		return 0, fmt.Errorf("error setting verb: %w", err)
	}

	err = i.Run()
	if err != nil {
		return 0, err
	}

	output, err := i.program.Fetch(outputPosition)
	return output, err
}

// Run runs the Intcode program
func (i *Intcode) Run() error {
	for !i.program.Halted {

		n, err := i.program.Fetch(i.program.InstructionPointer)
		if err != nil {
			return fmt.Errorf("error fetching instruction: %w", err)
		}

		parsedInstruction, err := instruction.ParseInstruction(n)
		if err != nil {
			return fmt.Errorf("error parsing instruction: %w", err)
		}

		err = parsedInstruction.Execute(i.program)
		if err != nil {
			return fmt.Errorf("error executing instruction: %w", err)
		}

		i.RLock()
		if i.shouldStop {
			i.RUnlock()
			break
		}
		i.RUnlock()
	}
	return nil
}

// Stop stops the Intcode program
func (i *Intcode) Stop() {
	i.Lock()
	defer i.Unlock()
	i.shouldStop = true
}
