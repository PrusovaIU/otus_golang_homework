package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// Write me
	})
}

func TestPushOut(t *testing.T) {
	t.Run("push_out_by_len", func(t *testing.T) {
		c := NewCache(2)
		c.Set("aaa", 1)
		c.Set("bbb", 2)
		wasInCache := c.Set("ccc", 3)
		require.False(t, wasInCache)
		_, ok := c.Get("aaa")
		require.False(t, ok)
	})

	t.Run("push_out_old_elements", func(t *testing.T) {
		c := NewCache(2)
		c.Set("aaa", 1)
		c.Set("bbb", 2)
		c.Get("aaa")
		c.Set("ccc", 3)
		_, ok := c.Get("bbb")
		require.False(t, ok)
	})
}
