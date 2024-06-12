package main

import (
	"errors"
	"os"
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw08_envdir_tool/mocks"
	"github.com/stretchr/testify/require"
)

func TestFormEnv(t *testing.T) {
	paramName := "test"
	cleanEnv := func(paramName string) {
		if _, exists := os.LookupEnv(paramName); exists {
			os.Unsetenv(paramName)
		}
	}
	t.Run("need_remove", func(t *testing.T) {
		err := os.Setenv(paramName, "test_value")
		require.NoError(t, err)
		defer cleanEnv(paramName)
		envValue := NewEnvValue("", true)
		env := Environment{paramName: envValue}
		err = formEnv(env)
		require.NoError(t, err)
		_, exists := os.LookupEnv(paramName)
		require.False(t, exists)
	})

	test_not_remove := func(t *testing.T) {
		defer cleanEnv(paramName)
		exValue := "test_value"
		envValue := NewEnvValue(exValue, false)
		env := Environment{paramName: envValue}
		err := formEnv(env)
		require.NoError(t, err)
		value, exists := os.LookupEnv(paramName)
		require.True(t, exists)
		require.Equal(t, exValue, value)
	}

	t.Run("exists", func(t *testing.T) {
		err := os.Setenv(paramName, "old_value")
		require.NoError(t, err)
		test_not_remove(t)
	})

	t.Run("not_exists", func(t *testing.T) {
		test_not_remove(t)
	})
}

func TestProcessManage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		processMock := mocks.NewProcess(t)
		processMock.EXPECT().Start().Return(nil)
		processMock.EXPECT().Wait().Return(nil)

		err := processManage(processMock)
		require.NoError(t, err)
	})

	exErr := errors.New("Test error")

	t.Run("start_err", func(t *testing.T) {
		processMock := mocks.NewProcess(t)
		processMock.EXPECT().Start().Return(exErr)

		err := processManage(processMock)
		require.ErrorIs(t, err, exErr)
	})

	t.Run("wait_err", func(t *testing.T) {
		processMock := mocks.NewProcess(t)
		processMock.EXPECT().Start().Return(nil)
		processMock.EXPECT().Wait().Return(exErr)

		err := processManage(processMock)
		require.ErrorIs(t, err, exErr)
	})
}

func TestRunCmd(t *testing.T) {
	// Place your code here
}
