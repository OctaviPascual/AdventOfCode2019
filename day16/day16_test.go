package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		signal: signal{
			digits: []digit{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	input := `12345678`
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
			input:    "80871224585914546619083218645595",
			expected: "24176176",
		},
		"test 2": {
			input:    "19617804207202209144916044189917",
			expected: "73745418",
		},
		"test 3": {
			input:    "69317163492948606335995924319873",
			expected: "52432133",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			day, err := NewDay(testCase.input)
			require.NoError(t, err)

			answer, err := day.SolvePartOne()
			require.NoError(t, err)

			assert.Equal(t, testCase.expected, answer)
		})
	}
}

func TestSolvePartTwo(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"test 1": {
			input:    "03036732577212944063491565474664",
			expected: "84462026",
		},
		"test 2": {
			input:    "02935109699940807407585447034323",
			expected: "78725270",
		},
		"test 3": {
			input:    "03081770884921959731165446850517",
			expected: "53553731",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			day, err := NewDay(testCase.input)
			require.NoError(t, err)

			answer, err := day.SolvePartTwo()
			require.NoError(t, err)

			assert.Equal(t, testCase.expected, answer)
		})
	}
}

func TestSequenceNext(t *testing.T) {
	initialSequence := signal{
		digits: []digit{1, 2, 3, 4, 5, 6, 7, 8},
	}

	after1Phase := signal{
		digits: []digit{4, 8, 2, 2, 6, 1, 5, 8},
	}
	assert.Equal(t, after1Phase, initialSequence.next())

	after2Phases := signal{
		digits: []digit{3, 4, 0, 4, 0, 4, 3, 8},
	}
	assert.Equal(t, after2Phases, after1Phase.next())

	after3Phases := signal{
		digits: []digit{0, 3, 4, 1, 5, 5, 1, 8},
	}
	assert.Equal(t, after3Phases, after2Phases.next())

	after4Phases := signal{
		digits: []digit{0, 1, 0, 2, 9, 4, 9, 8},
	}
	assert.Equal(t, after4Phases, after3Phases.next())
}
