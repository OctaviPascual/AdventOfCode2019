package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHaltOpcode(t *testing.T) {
	assert.Equal(t, haltOpcode, halt{}.opcode())
}
