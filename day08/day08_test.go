package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		image: &encodedImage{
			width:  3,
			height: 2,
			layers: []layer{
				{
					pixels: []int{1, 2, 3, 4, 5, 6},
				},
				{
					pixels: []int{7, 8, 9, 0, 1, 2},
				},
			},
		},
	}
	imageWidth = 3
	imageHeight = 2
	input := `123456789012`

	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		image: &encodedImage{
			width:  3,
			height: 2,
			layers: []layer{
				{
					pixels: []int{1, 2, 3, 4, 5, 6},
				},
				{
					pixels: []int{7, 8, 9, 0, 1, 2},
				},
			},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "1", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		image: &encodedImage{
			width:  2,
			height: 2,
			layers: []layer{
				{
					pixels: []int{0, 2, 2, 2},
				},
				{
					pixels: []int{1, 1, 2, 2},
				},
				{
					pixels: []int{2, 2, 1, 2},
				},
				{
					pixels: []int{0, 0, 0, 0},
				},
			},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	expected := "\n" + blackEmoji + whiteEmoji + "\n" + whiteEmoji + blackEmoji + "\n"
	assert.Equal(t, expected, answer)
}
