package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewURLSafeString(t *testing.T) {
	got, err := NewURLSafeString()
	print(got)
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
}
