package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapabilityFlag_Has(t *testing.T) {
	flag := CapabilityFlag(1<<2)
	assert.True(t, flag.Has(clientLongFlag))
}
