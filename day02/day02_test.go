package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDay(t *testing.T) {
	day := &day{
		initialState: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		intcodeProgram: intcodeProgram{
			memory: make([]int, 12),
		},
	}
	input := `1,9,10,3,2,3,11,0,99,30,40,50`
	assert.Equal(t, day, NewDay(input))
}

func TestString(t *testing.T) {
	assert.Equal(t, "1", intcodeOutput(1).String())
}

func TestRun(t *testing.T) {
	testCases := map[string]struct {
		input    intcodeProgram
		expected intcodeProgram
	}{
		"test 1": {
			input: intcodeProgram{
				memory: []int{1, 0, 0, 0, 99},
			},
			expected: intcodeProgram{
				memory: []int{2, 0, 0, 0, 99},
			},
		},
		"test 2": {
			input: intcodeProgram{
				memory: []int{2, 3, 0, 3, 99},
			},
			expected: intcodeProgram{
				memory: []int{2, 3, 0, 6, 99},
			},
		},
		"test 3": {
			input: intcodeProgram{
				memory: []int{2, 4, 4, 5, 99, 0},
			},
			expected: intcodeProgram{
				memory: []int{2, 4, 4, 5, 99, 9801},
			},
		},
		"test 4": {
			input: intcodeProgram{
				memory: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			},
			expected: intcodeProgram{
				memory: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			},
		},
		"test 5": {
			input: intcodeProgram{
				memory: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			},
			expected: intcodeProgram{
				memory: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testCase.input.run()
			assert.Equal(t, testCase.expected, testCase.input)
		})
	}
}

func TestSolvePartOne(t *testing.T) {
	day := &day{
		initialState: []int{1, 0, 0, 0, 99, 99, 99, 99, 99, 99, 99, 99, 99},
		intcodeProgram: intcodeProgram{
			memory: make([]int, 13),
		},
	}
	assert.Equal(t, intcodeOutput(101), day.SolvePartOne())
}

func TestSolvePartTwo(t *testing.T) {
	day := &day{
		initialState: append([]int{1, 0, 0, 0, 99, 19690720}, make([]int, 94)...),
		intcodeProgram: intcodeProgram{
			memory: make([]int, 100),
		},
	}
	assert.Equal(t, answer(305), day.SolvePartTwo())
}
