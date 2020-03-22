package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		moons: []*moon{
			{
				name:     io,
				position: position{-1, 0, 2},
			},
			{
				name:     europa,
				position: position{2, -10, -7},
			},
			{
				name:     ganymede,
				position: position{4, -8, 8},
			},
			{
				name:     callisto,
				position: position{3, 5, -1},
			},
		},
	}
	input := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	stepsToSimulate = 100
	day := &Day{
		moons: []*moon{
			{
				name:     io,
				position: position{-8, -10, 0},
			},
			{
				name:     europa,
				position: position{5, 5, 10},
			},
			{
				name:     ganymede,
				position: position{2, -7, 3},
			},
			{
				name:     callisto,
				position: position{9, -8, -3},
			},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "1940", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		moons: []*moon{
			{
				name:     io,
				position: position{-8, -10, 0},
			},
			{
				name:     europa,
				position: position{5, 5, 10},
			},
			{
				name:     ganymede,
				position: position{2, -7, 3},
			},
			{
				name:     callisto,
				position: position{9, -8, -3},
			},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "4686774924", answer)
}

func TestSimulateMotion(t *testing.T) {
	moons := []*moon{
		{
			name:     io,
			position: position{-1, 0, 2},
		},
		{
			name:     europa,
			position: position{2, -10, -7},
		},
		{
			name:     ganymede,
			position: position{4, -8, 8},
		},
		{
			name:     callisto,
			position: position{3, 5, -1},
		},
	}

	expected := []*moon{
		{
			name:     io,
			position: position{2, 1, -3},
			velocity: velocity{-3, -2, 1},
		},
		{
			name:     europa,
			position: position{1, -8, 0},
			velocity: velocity{-1, 1, 3},
		},
		{
			name:     ganymede,
			position: position{3, -6, 1},
			velocity: velocity{3, 2, -3},
		},
		{
			name:     callisto,
			position: position{2, 0, 4},
			velocity: velocity{1, -1, -1},
		},
	}

	simulateMotion(moons, 10, dimensions...)

	assert.Equal(t, expected, moons)
}

func TestTotalEnergy(t *testing.T) {
	moons := []*moon{
		{
			name:     io,
			position: position{2, 1, -3},
			velocity: velocity{-3, -2, 1},
		},
		{
			name:     europa,
			position: position{1, -8, 0},
			velocity: velocity{-1, 1, 3},
		},
		{
			name:     ganymede,
			position: position{3, -6, 1},
			velocity: velocity{3, 2, -3},
		},
		{
			name:     callisto,
			position: position{2, 0, 4},
			velocity: velocity{1, -1, -1},
		},
	}
	assert.Equal(t, 179, totalEnergy(moons))
}
