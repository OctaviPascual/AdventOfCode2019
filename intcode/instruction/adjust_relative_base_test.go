package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeBaseOffsetOpcode(t *testing.T) {
	assert.Equal(t, adjustRelativeBaseOpcode, adjustRelativeBase{}.opcode())
}
