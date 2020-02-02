package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpIfTrueOpcode(t *testing.T) {
	assert.Equal(t, jumpIfTrueOpcode, jumpIfTrue{}.opcode())
}
