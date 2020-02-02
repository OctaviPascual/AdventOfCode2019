package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpIfFalseOpcode(t *testing.T) {
	assert.Equal(t, jumpIfFalseOpcode, jumpIfFalse{}.opcode())
}
