package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		program: "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99",
	}
	input := `109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"test 1": {
			input:    "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99",
			expected: "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99",
		},
		"test 2": {
			input:    "1102,34915192,34915192,7,4,7,99,0",
			expected: "1219070632396864",
		},
		"test 3": {
			input:    "104,1125899906842624,99",
			expected: "1125899906842624",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			day, err := NewDay(testCase.input)
			require.NoError(t, err)

			actual, err := day.SolvePartOne()
			require.NoError(t, err)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
