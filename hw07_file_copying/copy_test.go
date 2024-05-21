package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	// Place your code here.
}

func TestOpenFromFile(t *testing.T) {
	t.Run("success_opening", func(t *testing.T) {
		file_name := "test_file"
		_, err := os.Create(file_name)
		require.Nil(t, err)
	})
}
