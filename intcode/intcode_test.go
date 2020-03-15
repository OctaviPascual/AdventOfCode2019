package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunWithNounAndVerb(t *testing.T) {
	testCases := map[string]struct {
		program  string
		noun     int
		verb     int
		expected int
	}{
		"test 1": {
			program:  "1,0,0,0,99",
			noun:     0,
			verb:     0,
			expected: 2,
		},
		"test 2": {
			program:  "2,3,0,3,99",
			noun:     3,
			verb:     0,
			expected: 2,
		},
		"test 3": {
			program:  "2,4,4,5,99,0",
			noun:     4,
			verb:     4,
			expected: 2,
		},
		"test 4": {
			program:  "1,1,1,4,99,5,6,0,99",
			noun:     1,
			verb:     1,
			expected: 30,
		},
		"test 5": {
			program:  "1,9,10,3,2,3,11,0,99,30,40,50",
			noun:     9,
			verb:     10,
			expected: 3500,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			program, err := NewIntcodeProgram(testCase.program)
			require.NoError(t, err)

			output, err := program.RunWithNounAndVerb(testCase.noun, testCase.verb)
			require.NoError(t, err)

			assert.Equal(t, testCase.expected, output)
		})
	}
}

func TestRunWithInput(t *testing.T) {
	testCases := map[string]struct {
		program  string
		input    int
		expected int
	}{
		"program that outputs 1 if input is equal to 8 using position mode": {
			program:  "3,9,8,9,10,9,4,9,99,-1,8",
			input:    8,
			expected: 1,
		},
		"program that outputs 1 if input is equal to 8 using immediate mode": {
			program:  "3,3,1108,-1,8,3,4,3,99",
			input:    8,
			expected: 1,
		},
		"program that outputs 0 if input is not equal to 8 using position mode": {
			program:  "3,9,8,9,10,9,4,9,99,-1,8",
			input:    7,
			expected: 0,
		},
		"program that outputs 0 if input is not equal to 8 using immediate mode": {
			program:  "3,3,1108,-1,8,3,4,3,99",
			input:    7,
			expected: 0,
		},
		"program that outputs 1 if input is less than 8 using position mode": {
			program:  "3,9,7,9,10,9,4,9,99,-1,8",
			input:    7,
			expected: 1,
		},
		"program that outputs 1 if input is less than 8 using immediate mode": {
			program:  "3,3,1107,-1,8,3,4,3,99",
			input:    7,
			expected: 1,
		},
		"program that outputs 0 if input is greater or equal than 8 using position mode": {
			program:  "3,9,7,9,10,9,4,9,99,-1,8",
			input:    10,
			expected: 0,
		},
		"program that outputs 0 if input is greater or equal than 8 using immediate mode": {
			program:  "3,3,1107,-1,8,3,4,3,99",
			input:    10,
			expected: 0,
		},
		"program that outputs 0 if input is 0 using position mode": {
			program:  "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			input:    0,
			expected: 0,
		},
		"program that outputs 0 if input is 0 using immediate mode": {
			program:  "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			input:    0,
			expected: 0,
		},
		"program that outputs 1 if input is not 0 using position mode": {
			program:  "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			input:    12,
			expected: 1,
		},
		"program that outputs 1 if input is not 0 using immediate mode": {
			program:  "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			input:    12,
			expected: 1,
		},
		"program that outputs 999 if input is less than 8": {
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31," +
				"1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104," +
				"999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			input:    7,
			expected: 999,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			program, err := NewIntcodeProgram(testCase.program)
			require.NoError(t, err)

			outputs, err := program.RunWithInput(testCase.input)
			require.NoError(t, err)

			assert.Equal(t, testCase.expected, outputs[0])
		})
	}
}
