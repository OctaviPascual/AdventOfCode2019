package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		program: "3,3,1107,-1,3,3,4,3,99",
	}
	input := `3,3,1107,-1,3,3,4,3,99`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		// program that outputs 1 if input is less than 3 using immediate mode
		program: "3,3,1107,-1,3,3,4,3,99",
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "1", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		// program that outputs 1 if input is less than 3 using immediate mode
		program: "3,3,1107,-1,3,3,4,3,99",
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "0", answer)
}
