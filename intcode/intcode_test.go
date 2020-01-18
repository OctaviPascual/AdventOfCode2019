package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIntcodeProgram(t *testing.T) {

}

func TestRun(t *testing.T) {
	testCases := map[string]struct {
		input    Program
		expected Program
	}{
		"test 1": {
			input: Program{
				memory: []int{1, 0, 0, 0, 99},
			},
			expected: Program{
				memory: []int{2, 0, 0, 0, 99},
			},
		},
		"test 2": {
			input: Program{
				memory: []int{2, 3, 0, 3, 99},
			},
			expected: Program{
				memory: []int{2, 3, 0, 6, 99},
			},
		},
		"test 3": {
			input: Program{
				memory: []int{2, 4, 4, 5, 99, 0},
			},
			expected: Program{
				memory: []int{2, 4, 4, 5, 99, 9801},
			},
		},
		"test 4": {
			input: Program{
				memory: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			},
			expected: Program{
				memory: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			},
		},
		"test 5": {
			input: Program{
				memory: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			},
			expected: Program{
				memory: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := testCase.input.run()
			require.NoError(t, err)

			assert.Equal(t, testCase.expected, testCase.input)
		})
	}
}
