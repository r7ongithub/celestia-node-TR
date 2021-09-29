package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	dir := t.TempDir()
	err := Init(dir, DefaultFullConfig())
	require.NoError(t, err)
	ok := IsInit(dir)
	assert.True(t, ok)
}
