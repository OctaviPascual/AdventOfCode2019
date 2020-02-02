package program

import (
	"errors"
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

	inputBuffer  []int
	outputBuffer []int

	memory []int
}

// NewProgram creates a new program from the program string
func NewProgram(programString string) (*Program, error) {
	memory, err := parse(programString)
	if err != nil {
		return nil, err
	}

	return &Program{
		memory: memory,
	}, nil
}

func parse(programString string) ([]int, error) {
	tokens := strings.Split(programString, ",")
	values := make([]int, 0, len(tokens))

	for _, token := range tokens {

		value, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("invalid value %s: %w", token, err)
		}

		values = append(values, value)
	}

	return values, nil
}

// Fetch fetches value at given position
func (p *Program) Fetch(position int) (int, error) {
	if position < 0 || position >= len(p.memory) {
		return 0, fmt.Errorf("fetch error: invalid memory position: %d", position)
	}
	return p.memory[position], nil
}

// Store stores value at given position
func (p *Program) Store(position int, value int) error {
	if position < 0 || position >= len(p.memory) {
		return fmt.Errorf("store error: invalid memory position: %d", position)
	}
	p.memory[position] = value
	return nil
}

// ReadInput reads the input value of the program
func (p *Program) ReadInput() (int, error) {
	if len(p.inputBuffer) == 0 {
		return 0, errors.New("no input in input buffer")
	}
	input := p.inputBuffer[0]
	p.inputBuffer = p.inputBuffer[1:]
	return input, nil
}

// SetInput sets the input of the program to the provided one
func (p *Program) SetInput(input []int) {
	p.inputBuffer = input
}

// GetOutput gets the output of the program
func (p *Program) GetOutput() []int {
	return p.outputBuffer
}

// WriteOutput writes the output value to the program
func (p *Program) WriteOutput(output int) {
	p.outputBuffer = append(p.outputBuffer, output)
}
