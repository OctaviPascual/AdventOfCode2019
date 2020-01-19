package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIntcodeProgram(t *testing.T) {
	testCases := map[string]struct {
		program  string
		expected *Program
	}{
		"test valid program 1": {
			program: "1,0,0,0,99",
			expected: &Program{
				memory: []int{1, 0, 0, 0, 99},
			},
		},
		"test valid program 2": {
			program: "2,3,0,3,99",
			expected: &Program{
				memory: []int{2, 3, 0, 3, 99},
			},
		},
		"test valid program 3": {
			program: "1,9,10,3,2,3,11,0,99,30,40,50",
			expected: &Program{
				memory: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			},
		},
		"test invalid program 1": {
			program:  "",
			expected: nil,
		},
		"test invalid program 2": {
			program:  "1,2,,3",
			expected: nil,
		},
		"test invalid program 3": {
			program:  "1,2,3,invalid",
			expected: nil,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			program, err := NewIntcodeProgram(testCase.program)
			assert.Equal(t, testCase.expected, program)
			if testCase.expected == nil {
				require.Error(t, err)
			}
		})
	}
}

func TestRun(t *testing.T) {
	testCases := map[string]struct {
		program  string
		noun     int
		verb     int
		expected Output
	}{
		"test 1": {
			program:  "1,0,0,0,99",
			noun:     0,
			verb:     0,
			expected: Output(2),
		},
		"test 2": {
			program:  "2,3,0,3,99",
			noun:     3,
			verb:     0,
			expected: Output(2),
		},
		"test 3": {
			program:  "2,4,4,5,99,0",
			noun:     4,
			verb:     4,
			expected: Output(2),
		},
		"test 4": {
			program:  "1,1,1,4,99,5,6,0,99",
			noun:     1,
			verb:     1,
			expected: Output(30),
		},
		"test 5": {
			program:  "1,9,10,3,2,3,11,0,99,30,40,50",
			noun:     9,
			verb:     10,
			expected: Output(3500),
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			program, err := NewIntcodeProgram(testCase.program)
			require.NoError(t, err)

			output, err := program.Run(testCase.noun, testCase.verb)

			assert.Equal(t, testCase.expected, output)
		})
	}
}
