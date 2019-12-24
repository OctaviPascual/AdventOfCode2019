package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &day{
		spacecraft: spacecraft{
			modules: []module{
				{mass: 12},
				{mass: 14},
				{mass: 1969},
				{mass: 100756},
			},
		},
	}
	input := `12
14
1969
100756`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestString(t *testing.T) {
	assert.Equal(t, "34241", fuel(34241).String())
}

func TestSolvePartOne(t *testing.T) {
	day := &day{
		spacecraft: spacecraft{
			modules: []module{
				{mass: 12},
				{mass: 14},
				{mass: 1969},
				{mass: 100756},
			},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, fuel(34241), answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &day{
		spacecraft: spacecraft{
			modules: []module{
				{mass: 12},
				{mass: 14},
				{mass: 1969},
				{mass: 100756},
			},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, fuel(51316), answer)
}
