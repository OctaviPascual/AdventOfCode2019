package instruction

import (
	"fmt"

	"github.com/OctaviPascual/AdventOfCode2019/intcode/program"
)

type opcode uint8

const (
	addOpcode                opcode = 1
	multiplyOpcode           opcode = 2
	inputOpcode              opcode = 3
	outputOpcode             opcode = 4
	jumpIfTrueOpcode         opcode = 5
	jumpIfFalseOpcode        opcode = 6
	lessThanOpcode           opcode = 7
	equalsOpcode             opcode = 8
	adjustRelativeBaseOpcode opcode = 9
	haltOpcode               opcode = 99
)

type parameterMode int

const (
	positionMode  parameterMode = 0
	immediateMode parameterMode = 1
	relativeMode  parameterMode = 2
)

// Instruction represent an instruction of the program
type Instruction interface {
	// Execute executes an instruction and modifies the program accordingly
	Execute(program *program.Program) error

	opcode() opcode
}

func getParameter(
	position int,
	parameterMode parameterMode,
	program *program.Program,
) (int, error) {
	switch parameterMode {
	case positionMode:
		address, err := program.Fetch(program.InstructionPointer + position)
		if err != nil {
			return 0, err
		}
		return program.Fetch(address)
	case immediateMode:
		return program.Fetch(program.InstructionPointer + position)
	case relativeMode:
		address, err := program.Fetch(program.InstructionPointer + position)
		if err != nil {
			return 0, err
		}
		return program.Fetch(address + program.RelativeBase)
	default:
		return 0, fmt.Errorf("invalid parameter mode: %d", parameterMode)
	}
}

func getFirstParameter(parameterMode parameterMode, program *program.Program) (int, error) {
	return getParameter(1, parameterMode, program)
}

func getSecondParameter(parameterMode parameterMode, program *program.Program) (int, error) {
	return getParameter(2, parameterMode, program)
}

func storeWithParameter(
	position, value int,
	parameterMode parameterMode,
	program *program.Program,
) error {
	address, err := program.Fetch(program.InstructionPointer + position)
	if err != nil {
		return err
	}
	switch parameterMode {
	case positionMode:
		return program.Store(address, value)
	case immediateMode:
		return program.Store(address, value)
	case relativeMode:
		return program.Store(address+program.RelativeBase, value)
	default:
		return fmt.Errorf("invalid parameter mode: %d", parameterMode)
	}
}

func storeWithFirstParameter(value int, parameterMode parameterMode, program *program.Program) error {
	return storeWithParameter(1, value, parameterMode, program)
}

func storeWithThirdParameter(value int, parameterMode parameterMode, program *program.Program) error {
	return storeWithParameter(3, value, parameterMode, program)
}

// ParseInstruction parses a value n to an instruction
func ParseInstruction(n int) (Instruction, error) {
	switch opcode(n % 100) {
	case addOpcode:
		return add{
			firstParameterMode:  parameterMode((n / 100) % 10),
			secondParameterMode: parameterMode((n / 1000) % 10),
			thirdParameterMode:  parameterMode((n / 10000) % 10),
		}, nil
	case multiplyOpcode:
		return multiply{
			firstParameterMode:  parameterMode((n / 100) % 10),
			secondParameterMode: parameterMode((n / 1000) % 10),
			thirdParameterMode:  parameterMode((n / 10000) % 10),
		}, nil
	case inputOpcode:
		return input{
			firstParameterMode: parameterMode((n / 100) % 10),
		}, nil
	case outputOpcode:
		return output{
			firstParameterMode: parameterMode((n / 100) % 10),
		}, nil
	case jumpIfTrueOpcode:
		return jumpIfTrue{
			firstParameterMode:  parameterMode((n / 100) % 10),
			secondParameterMode: parameterMode((n / 1000) % 10),
		}, nil
	case jumpIfFalseOpcode:
		return jumpIfFalse{
			firstParameterMode:  parameterMode((n / 100) % 10),
			secondParameterMode: parameterMode((n / 1000) % 10),
		}, nil
	case lessThanOpcode:
		return lessThan{
			firstParameterMode:  parameterMode((n / 100) % 10),
			secondParameterMode: parameterMode((n / 1000) % 10),
			thirdParameterMode:  parameterMode((n / 10000) % 10),
		}, nil
	case equalsOpcode:
		return equals{
			firstParameterMode:  parameterMode((n / 100) % 10),
			secondParameterMode: parameterMode((n / 1000) % 10),
			thirdParameterMode:  parameterMode((n / 10000) % 10),
		}, nil
	case adjustRelativeBaseOpcode:
		return adjustRelativeBase{
			firstParameterMode: parameterMode((n / 100) % 10),
		}, nil
	case haltOpcode:
		return halt{}, nil
	default:
		return nil, fmt.Errorf("unknown opcode %d", n)
	}
}
