package hw09structvalidator

import (
	"errors"
	"reflect"
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
)

func TestSliceValidator(t *testing.T) {
	type TestStruct struct {
		testField []int
	}

	testSlice := []int{1, 2, 3}

	testStruct := TestStruct{
		testField: testSlice,
	}

	fieldValue := reflect.ValueOf(testStruct).FieldByName("testField")
	fieldType, _ := reflect.TypeOf(testStruct).FieldByName("testField")

	tasks := []struct {
		name             string
		validationResult validationErrs.ValidationError
	}{
		// {name: "test_success", validationResult: validationErrs.ValidationError{}},
		{name: "test_errors", validationResult: validationErrs.ValidationError{
			Field: "test",
			Err:   errors.New("Test error"),
		}},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			elementValidatorMock := mocks.NewElementValidatorInterface(t)
			elementValidatorMock.
				On("Validate", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.validationResult)

			sv := SliceValidator{}
			sv.ElementValidator = elementValidatorMock

			errs := sv.Validate(fieldValue, fieldType, "tag")
			if tc.validationResult.IsErr() {
				require.Len(t, errs, len(testSlice))
			} else {
				require.Len(t, errs, 0)
			}
			require.Equal(t, len(elementValidatorMock.Calls), len(testSlice))
		})

	}
}
