package main

import (
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw08_envdir_tool/mocks"
	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fileContent := "test_value"
		fileMock := mocks.NewEnvFile(t)
		fileMock.EXPECT().ReadLine().Return([]byte(fileContent), false, nil)

		result, err := readFile(fileMock)
		require.NoError(t, err)
		require.Equal(t, fileContent, result.Value)
		require.False(t, result.NeedRemove)
	})

}

func TestReadDir(t *testing.T) {
	// Place your code here
}
