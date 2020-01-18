package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		passwordRange: passwordRange{
			from: 367479,
			to:   893698,
		},
	}
	input := `367479-893698`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		passwordRange: passwordRange{
			from: 900000,
			to:   999999,
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "1", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		passwordRange: passwordRange{
			from: 900000,
			to:   999999,
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "0", answer)
}

func TestIsSixDigitNumber(t *testing.T) {
	testCases := map[string]struct {
		input    password
		expected bool
	}{
		"test 1": {
			input:    password(0),
			expected: false,
		},
		"test 2": {
			input:    password(123123),
			expected: true,
		},
		"test 3": {
			input:    password(1231235),
			expected: false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.input.isSixDigitNumber())
		})
	}
}

func TestHasTwoAdjacentDigits(t *testing.T) {
	testCases := map[string]struct {
		input    password
		expected bool
	}{
		"test 1": {
			input:    password(0),
			expected: false,
		},
		"test 2": {
			input:    password(11),
			expected: true,
		},
		"test 3": {
			input:    password(1221),
			expected: true,
		},
		"test 4": {
			input:    password(1233),
			expected: true,
		},
		"test 5": {
			input:    password(1123),
			expected: true,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.input.hasTwoAdjacentDigits())
		})
	}
}

func TestHasNeverDecreasingDigits(t *testing.T) {
	testCases := map[string]struct {
		input    password
		expected bool
	}{
		"test 1": {
			input:    password(0),
			expected: true,
		},
		"test 2": {
			input:    password(112),
			expected: true,
		},
		"test 3": {
			input:    password(1231),
			expected: false,
		},
		"test 4": {
			input:    password(1233),
			expected: true,
		},
		"test 5": {
			input:    password(112359),
			expected: true,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.input.hasNeverDecreasingDigits())
		})
	}
}

func TestHasExactlyTwoAdjacentDigits(t *testing.T) {
	testCases := map[string]struct {
		input    password
		expected bool
	}{
		"test 1": {
			input:    password(0),
			expected: false,
		},
		"test 2": {
			input:    password(111122),
			expected: true,
		},
		"test 3": {
			input:    password(123444),
			expected: false,
		},
		"test 4": {
			input:    password(112233),
			expected: true,
		},
		"test 5": {
			input:    password(111222),
			expected: false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.input.hasExactlyTwoAdjacentDigits())
		})
	}
}
