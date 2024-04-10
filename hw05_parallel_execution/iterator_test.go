package hw05parallelexecution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterator(t *testing.T) {
	n := 3
	i := NewIterator(n)
	require.True(t, i.Check())
	for j := 0; j < n; j++ {
		k, ok := i.Get()
		require.Equal(t, j, k)
		require.True(t, ok)
	}
	k, ok := i.Get()
	require.Equal(t, n, k)
	require.False(t, ok)
}
