package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		orbitMap: orbitMap{
			orbitalRelationships: []orbitalRelationship{
				{
					parent:    "COM",
					satellite: "B",
				},
				{
					parent:    "B",
					satellite: "C",
				},
				{
					parent:    "C",
					satellite: "D",
				},
				{
					parent:    "D",
					satellite: "E",
				},
				{
					parent:    "E",
					satellite: "F",
				},
				{
					parent:    "B",
					satellite: "G",
				},
				{
					parent:    "G",
					satellite: "H",
				},
				{
					parent:    "D",
					satellite: "I",
				},
				{
					parent:    "E",
					satellite: "J",
				},
				{
					parent:    "J",
					satellite: "K",
				},
				{
					parent:    "K",
					satellite: "L",
				},
			},
		},
	}
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "42", answer)
}

func TestSolvePartTwo(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "4", answer)
}

func TestLowestCommonAncestor(t *testing.T) {
	testCases := map[string]struct {
		path1    []string
		path2    []string
		expected string
		err      bool
	}{
		"test 1": {
			path1:    []string{"B", "A", "COM"},
			path2:    []string{"D", "C", "A", "COM"},
			expected: "A",
		},
		"test 2": {
			path1:    []string{"B", "A", "COM"},
			path2:    []string{"B", "A", "COM"},
			expected: "B",
		},
		"test 3": {
			path1:    []string{"B", "A", "COM"},
			path2:    []string{"C", "B", "A", "COM"},
			expected: "B",
		},
		"test 4": {
			path1:    []string{"B", "A"},
			path2:    []string{"C", "B", "A", "COM"},
			expected: "B",
			err:      true,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			path1 := testCase.path1
			path2 := testCase.path2
			ancestor, err := lowestCommonAncestor(path1, path2)

			if testCase.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, testCase.expected, ancestor)
		})
	}

}
