package program

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProgram(t *testing.T) {
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
			program, err := NewProgram(testCase.program)
			assert.Equal(t, testCase.expected, program)
			if testCase.expected == nil {
				require.Error(t, err)
			}
		})
	}
}
