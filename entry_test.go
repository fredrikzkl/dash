package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntry(t *testing.T) {
	path := "/users/john/warlock"
	entry := newEntry(path)
	require.Equal(t, "warlock", entry.name)
	require.Equal(t, path, entry.path)
}
