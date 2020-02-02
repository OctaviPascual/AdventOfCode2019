package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsOpcode(t *testing.T) {
	assert.Equal(t, equalsOpcode, equals{}.opcode())
}
