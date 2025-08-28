package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntry(t *testing.T) {
	path := "/users/john/warlock"
	entry := NewEntry(path)
	require.Equal(t, "warlock", entry.Name)
	require.Equal(t, path, entry.Path)
}
