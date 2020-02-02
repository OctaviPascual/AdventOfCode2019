package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLessThanOpcode(t *testing.T) {
	assert.Equal(t, lessThanOpcode, lessThan{}.opcode())
}
