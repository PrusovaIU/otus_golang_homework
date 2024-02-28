package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestCheckData(t *testing.T) {
	tests := []struct {
		input string
		err   error
	}{
		{input: "a4bc2d5e", err: nil},
		{input: "abccd", err: nil},
		{input: "", err: nil},
		{input: "aaa0b", err: nil},
		{input: `qwe\4\5`, err: nil},
		{input: `qwe\45`, err: nil},
		{input: `qwe\\5`, err: nil},
		{input: `qwe\\\3`, err: nil},
		{input: "3abc", err: ErrInvalidString},
		{input: "45", err: ErrInvalidString},
		{input: "aaa10b", err: ErrInvalidString},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			err := check_data(tc.input)
			require.Equal(t, tc.err, err)
		})
	}
}

func TestRepetition(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: `a\\2`, expected: `a\\\\`},
		{input: "\t2", expected: "\t\t"},
		{input: "a2", expected: "aa"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := repetition([]byte(tc.input))
			require.Equal(t, tc.expected, string(result))
		})
	}
}
