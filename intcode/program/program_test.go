package program

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemory(t *testing.T) {
	testCases := map[string]struct {
		program  string
		expected map[int]int
	}{
		"test valid program 1": {
			program: "1,0,0,0,99",
			expected: map[int]int{
				0: 1,
				1: 0,
				2: 0,
				3: 0,
				4: 99,
			},
		},
		"test valid program 2": {
			program: "2,3,0,3,99",
			expected: map[int]int{
				0: 2,
				1: 3,
				2: 0,
				3: 3,
				4: 99,
			},
		},
		"test valid program 3": {
			program: "1,9,10,3,2,3,11,0,99,30,40,50",
			expected: map[int]int{
				0:  1,
				1:  9,
				2:  10,
				3:  3,
				4:  2,
				5:  3,
				6:  11,
				7:  0,
				8:  99,
				9:  30,
				10: 40,
				11: 50,
			},
		},
		"test invalid program 1": {
			program:  "",
			expected: nil,
		},
		"test invalid program 2": {
			program:  "1,2,,3",
			expected: nil,
		},
		"test invalid program 3": {
			program:  "1,2,3,invalid",
			expected: nil,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			memory, err := newMemory(testCase.program)
			if err != nil {
				assert.Nil(t, testCase.expected)
				return
			}

			assert.Equal(t, testCase.expected, memory)
		})
	}
}
