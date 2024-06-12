package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormEnv(t *testing.T) {
	paramName := "test"
	cleanEnv := func(paramName string) {
		if _, exists := os.LookupEnv(paramName); exists {
			os.Unsetenv(paramName)
		}
	}
	t.Run("test_need_remove", func(t *testing.T) {
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

	t.Run("test_exists", func(t *testing.T) {
		err := os.Setenv(paramName, "old_value")
		require.NoError(t, err)
		test_not_remove(t)
	})

	t.Run("test_not_exists", func(t *testing.T) {
		test_not_remove(t)
	})
}

func TestRunCmd(t *testing.T) {
	// Place your code here
}
