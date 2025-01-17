package hw09structvalidator

import (
	"errors"
	"reflect"
	"testing"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
		validationResult error
	}{
		{name: "test_success", validationResult: nil},
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

			validation_errs, err := sv.Validate(fieldValue, fieldType, "tag")
			if tc.validationResult != nil {
				require.Len(t, validation_errs, len(testSlice))
			} else {
				require.Len(t, validation_errs, 0)
			}
			require.Equal(t, len(elementValidatorMock.Calls), len(testSlice))
			require.NoError(t, err)
		})
	}

	t.Run("test_error", func(t *testing.T) {
		elementValidatorMock := mocks.NewElementValidatorInterface(t)
		elementValidatorMock.
			On("Validate", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("Test error"))

		sv := SliceValidator{}
		sv.ElementValidator = elementValidatorMock
		validation_errs, err := sv.Validate(fieldValue, fieldType, "tag")
		require.Len(t, validation_errs, 0)
		require.Error(t, err)
	})
}
