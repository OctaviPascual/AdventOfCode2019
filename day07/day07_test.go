package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		program: "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
	}
	input := `3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`
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
			input:    "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			expected: "43210",
		},
		"test 2": {
			input: "3,23,3,24,1002,24,10,24,1002,23,-1,23," +
				"101,5,23,23,1,24,23,23,4,23,99,0,0",
			expected: "54321",
		},
		"test 3": {
			input: "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33," +
				"1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			expected: "65210",
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

func TestSolvePartTwo(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"test 1": {
			input: "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26," +
				"27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5",
			expected: "139629729",
		},
		"test 2": {
			input: "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54," +
				"-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4," +
				"53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10",
			expected: "18216",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			day, err := NewDay(testCase.input)
			require.NoError(t, err)

			actual, err := day.SolvePartTwo()
			require.NoError(t, err)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestGenerateAllPhaseSettingsCombinations(t *testing.T) {
	phaseSettings := []int{0, 1, 2}
	phaseSettingsCombinations := generateAllPhaseSettingsCombinations(phaseSettings)

	assert.Len(t, phaseSettingsCombinations, 6)

	assert.Contains(t, phaseSettingsCombinations, []int{0, 1, 2})
	assert.Contains(t, phaseSettingsCombinations, []int{0, 2, 1})

	assert.Contains(t, phaseSettingsCombinations, []int{1, 2, 0})
	assert.Contains(t, phaseSettingsCombinations, []int{1, 0, 2})

	assert.Contains(t, phaseSettingsCombinations, []int{2, 0, 1})
	assert.Contains(t, phaseSettingsCombinations, []int{2, 1, 0})
}
