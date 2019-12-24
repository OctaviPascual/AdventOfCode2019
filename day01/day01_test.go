package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDay(t *testing.T) {
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
	input := `12
14
1969
100756`
	assert.Equal(t, day, NewDay(input))
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
	assert.Equal(t, fuel(34241), day.SolvePartOne())
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
	assert.Equal(t, fuel(51316), day.SolvePartTwo())
}
