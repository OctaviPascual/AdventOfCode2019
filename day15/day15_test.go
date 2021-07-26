package day15

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

func TestBuildCommands(t *testing.T) {
	current := position{x: 3, y: -1}
	parent := map[position]position{
		{x: 1, y: 0}:  {x: 0, y: 0},
		{x: 2, y: 0}:  {x: 1, y: 0},
		{x: 2, y: -1}: {x: 2, y: 0},
		{x: 2, y: -2}: {x: 2, y: -1},
		{x: 3, y: -2}: {x: 2, y: -2},
		{x: 3, y: -1}: {x: 3, y: -2},
	}

	assert.Equal(t, []command{east, east, south, south, east, north}, buildCommands(current, parent))
}
