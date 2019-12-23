package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var day = &Day{
	spacecraft: spacecraft{
		modules: []module{
			{mass: 12},
			{mass: 14},
			{mass: 1969},
			{mass: 100756},
		},
	},
}

func TestNewDay(t *testing.T) {
	input := `12
14
1969
100756`
	assert.Equal(t, day, NewDay(input))
}

func TestFormat(t *testing.T) {
	assert.Equal(t, "34241", fuel(34241).Format())
}

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, fuel(34241), day.SolvePartOne())
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, fuel(51316), day.SolvePartTwo())
}
