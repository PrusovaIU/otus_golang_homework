package hw09structvalidator

import (
	"fmt"
	"testing"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:1"`
		Name   string
		Age    int      `validate:"min:18"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:12"`
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		name        string
		in          interface{}
		expectedErr validationErrs.ValidationErrors
	}{
		{
			"User",
			User{
				ID:     "1",
				Name:   "John",
				Age:    18,
				Email:  "john@doe.com",
				Role:   UserRole("admin"),
				Phones: []string{"111-111-1111", "222-222-2222"},
			},
			[]validationErrs.ValidationError{},
		},
		{
			"App",
			App{
				Version: "1.0.0",
			},
			[]validationErrs.ValidationError{},
		},
		{
			"Token",
			Token{
				Header:    []byte("header"),
				Payload:   []byte("payload"),
				Signature: []byte("signature"),
			},
			[]validationErrs.ValidationError{},
		},
		{
			"Response",
			Response{
				Code: 200,
			},
			[]validationErrs.ValidationError{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case %s", tt.name), func(t *testing.T) {
			// t.Parallel()

			result := Validate(tt.in)
			require.Equal(t, tt.expectedErr, result)
		})
	}
}

func TestErrors(t *testing.T) {
	user := User{
		ID:     "1",
		Name:   "John",
		Age:    10,
		Email:  "john@doe.com",
		Role:   UserRole("user"),
		Phones: []string{"111-111-1111", "222-222-2222"},
	}
	result := Validate(user)
	require.Len(t, result, 2)
}
