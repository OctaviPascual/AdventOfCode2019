package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOpcode(t *testing.T) {
	assert.Equal(t, addOpcode, add{}.opcode())
}
