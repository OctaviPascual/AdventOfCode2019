package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputOpcode(t *testing.T) {
	assert.Equal(t, outputOpcode, output{}.opcode())
}
