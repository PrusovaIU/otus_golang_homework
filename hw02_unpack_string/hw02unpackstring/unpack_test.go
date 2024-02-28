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

// func TestCheckDataSuccess(t *testing.T) {
// 	validStrings := []string{
// 		"a4bc2d5e",
// 		"abccd",
// 		"",
// 		"aaa0b",
// 		`qwe\4\5`,
// 		`qwe\45`,
// 		`qwe\\5`,
// 		`qwe\\\3`,
// 	}
// 	for _, tc := range validStrings {
// 		tc := tc
// 		t.Run(tc, func(t *testing.T) {
// 			err := check_data(tc)
// 			require.NoError(t, err)
// 		})
// 	}
// }

func TestCheckData(t *testing.T) {
	tests := []struct {
		input string
		// is_error	bool
		err error
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
	// {
	// 	{input: "a4bc2d5e", is_error: false},
	// 	{input: "abccd", is_error: false},
	// 	{input: "", is_error: false},
	// 	{input: "aaa0b", is_error: false},
	// 	{input: `qwe\4\5`, is_error: false},
	// 	{input: `qwe\45`, is_error: false},
	// 	{input: `qwe\\5`, is_error: false},
	// 	{input: `qwe\\\3`, is_error: false},
	// 	{input: "3abc", is_error: true},
	// 	{input: "45", is_error: true},
	// 	{input: "aaa10b", is_error: true},
	// }
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			err := check_data(tc.input)
			require.Equal(t, tc.err, err)
		})
	}
}
