package program

import (
	"fmt"
	"strconv"
	"strings"
)

// Program represents an Intcode program
type Program struct {
	// InstructionPointer is the current position of the instruction pointer
	InstructionPointer int
	// Halted indicates if the current program has been halted
	Halted bool
	// RelativeBase is the current position of the relative base
	RelativeBase int

	onInput  func() int
	onOutput func(output int)

	memory map[int]int
}

// NewProgram creates a new program from the program string
func NewProgram(
	programString string,
	onInput func() int,
	onOutput func(output int),
) (*Program, error) {
	memory, err := newMemory(programString)
	if err != nil {
		return nil, err
	}

	return &Program{
		memory:   memory,
		onInput:  onInput,
		onOutput: onOutput,
	}, nil
}

func newMemory(programString string) (map[int]int, error) {
	tokens := strings.Split(programString, ",")

	memory := make(map[int]int, len(tokens))
	for i, token := range tokens {

		value, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("invalid value %s: %w", token, err)
		}

		memory[i] = value
	}

	return memory, nil
}

// Fetch fetches Value at given position
func (p *Program) Fetch(position int) (int, error) {
	if position < 0 {
		return 0, fmt.Errorf("fetch error: invalid memory position: %d", position)
	}

	return p.memory[position], nil
}

// Store stores Value at given position
func (p *Program) Store(position int, value int) error {
	if position < 0 {
		return fmt.Errorf("store error: invalid memory position: %d", position)
	}

	p.memory[position] = value
	return nil
}

// ReadInput reads an input value from onInput function
func (p *Program) ReadInput() int {
	return p.onInput()
}

// WriteOutput writes an output value to onOutput function
func (p *Program) WriteOutput(output int) {
	p.onOutput(output)
}
