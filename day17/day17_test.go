package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		program: "1,330,331,332,109,3132,1102,1,1182,16,1101,1467,0,24,101",
	}
	input := `1,330,331,332,109,3132,1102,1,1182,16,1101,1467,0,24,101`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePart(t *testing.T) {
	camera := newCamera(`..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..
`)

	assert.Equal(t, 76, camera.calibrate())
}
