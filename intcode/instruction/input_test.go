package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputOpcode(t *testing.T) {
	assert.Equal(t, inputOpcode, input{}.opcode())
}
