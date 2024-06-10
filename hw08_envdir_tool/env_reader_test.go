package main

import (
	"errors"
	"os"
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw08_envdir_tool/mocks"
	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	fileContent := "test_value"

	cases := []struct {
		name        string
		fileContent string
		exResult    EnvValue
	}{
		{
			name:        "with_file_content",
			fileContent: "test_value",
			exResult:    NewEnvValue(fileContent, false),
		},
		{
			name:        "without_file_content",
			fileContent: "",
			exResult:    NewEnvValue("", true),
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			fileMock := mocks.NewEnvFile(t)
			fileMock.EXPECT().ReadLine().Return([]byte(tc.fileContent), false, nil)

			result, err := readFile(fileMock)
			require.NoError(t, err)
			require.Equal(t, tc.exResult.Value, result.Value)
			require.Equal(t, tc.exResult.NeedRemove, result.NeedRemove)
		})
	}

	t.Run("test_error", func(t *testing.T) {
		exError := errors.New("Test error")

		fileMock := mocks.NewEnvFile(t)
		fileMock.EXPECT().ReadLine().Return([]byte{}, false, exError)

		_, err := readFile(fileMock)
		require.ErrorIs(t, err, exError)
	})
}

func TestReadDir(t *testing.T) {
	dir_current_path, err := os.Getwd()
	dir_path := dir_current_path + "\\testdata\\env"
	require.NoError(t, err)
	result, err := ReadDir(dir_path)
	require.NoError(t, err)

	checkValue := func(value EnvValue, exValue string, exNeesRemove bool) {
		require.Equal(t, exValue, value.Value)
		require.Equal(t, exNeesRemove, value.NeedRemove)
	}

	checkValue(result["BAR"], "bar", false)
	checkValue(result["EMPTY"], "", true)
	checkValue(result["FOO"], "foo\nwith new line", false)
	checkValue(result["HELLO"], `"hello"`, false)
	checkValue(result["UNSET"], "", true)
}
