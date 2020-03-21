package day11

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
