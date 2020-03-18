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

	inputChannel  <-chan int
	outputChannel chan<- int

	memory map[int]int
}

// NewProgram creates a new program from the program string
func NewProgram(programString string) (*Program, error) {
	memory, err := parseProgram(programString)
	if err != nil {
		return nil, err
	}

	return &Program{
		memory: memory,
	}, nil
}

func parseProgram(programString string) (map[int]int, error) {
	memory := make(map[int]int)

	tokens := strings.Split(programString, ",")
	for i, token := range tokens {

		value, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("invalid value %s: %w", token, err)
		}

		memory[i] = value
	}

	return memory, nil
}

// Fetch fetches value at given position
func (p *Program) Fetch(position int) (int, error) {
	if position < 0 {
		return 0, fmt.Errorf("fetch error: invalid memory position: %d", position)
	}
	return p.memory[position], nil
}

// Store stores value at given position
func (p *Program) Store(position int, value int) error {
	if position < 0 {
		return fmt.Errorf("store error: invalid memory position: %d", position)
	}
	p.memory[position] = value
	return nil
}

// SetChannels sets the input and output channels of the program to the provided ones
func (p *Program) SetChannels(input <-chan int, output chan<- int) {
	p.inputChannel = input
	p.outputChannel = output
}

// ReadInput reads an input value of the program from the input channel
func (p *Program) ReadInput() int {
	return <-p.inputChannel
}

// WriteOutput writes the output value to the program
func (p *Program) WriteOutput(output int) {
	p.outputChannel <- output
}
