package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		wire1: wire{
			{
				direction: right,
				length:    1,
			},
			{
				direction: up,
				length:    2,
			},
		},
		wire2: wire{
			{
				direction: left,
				length:    2,
			},
			{
				direction: down,
				length:    1,
			},
		},
	}
	input := `R1,U2
L2,D1`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	testCases := map[string]struct {
		input  string
		answer string
	}{
		"test 1": {
			input: `R8,U5,L5,D3
U7,R6,D4,L4`,
			answer: "6",
		},
		"test 2": {
			input: `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			answer: "159",
		},
		"test 3": {
			input: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			answer: "135",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			day, err := NewDay(testCase.input)
			require.NoError(t, err)

			answer, err := day.SolvePartOne()
			require.NoError(t, err)

			assert.Equal(t, testCase.answer, answer)
		})
	}
}

func TestSolvePartTwo(t *testing.T) {
	testCases := map[string]struct {
		input  string
		answer string
	}{
		"test 1": {
			input: `R8,U5,L5,D3
U7,R6,D4,L4`,
			answer: "30",
		},
		"test 2": {
			input: `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			answer: "610",
		},
		"test 3": {
			input: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			answer: "410",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			day, err := NewDay(testCase.input)
			require.NoError(t, err)

			answer, err := day.SolvePartTwo()
			require.NoError(t, err)

			assert.Equal(t, testCase.answer, answer)
		})
	}
}
