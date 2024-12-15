package errors

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("\tfield %s %s\n", v.Field, v.Err)
}

func (v ValidationError) IsErr() bool {
	return v.Err != nil
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for _, err := range v {
		sb.WriteString(err.Error())
	}
	return sb.String()
}
