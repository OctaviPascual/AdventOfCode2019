package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyOpcode(t *testing.T) {
	assert.Equal(t, multiplyOpcode, multiply{}.opcode())
}
