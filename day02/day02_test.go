package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		program: "1,9,10,3,2,3,11,0,99,30,40,50",
	}
	input := `1,9,10,3,2,3,11,0,99,30,40,50`

	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		program: "1,0,0,0,99,99,99,99,99,99,99,99,99",
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "101", answer)
}

func TestSolvePartTwo(t *testing.T) {
	program := "1,0,0,0,99,19690720"
	for i := 0; i < 94; i++ {
		program += ",0"
	}
	day := &Day{
		program: program,
	}

	actual, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "305", actual)
}
