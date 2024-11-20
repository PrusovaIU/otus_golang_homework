package types_validators

import (
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/types_validators/mocks"
	"github.com/stretchr/testify/require"
)

func TestStringValidatorSuccess(t *testing.T) {
	tasks := []struct {
		name      string
		value     string
		condName  string
		condValue string
	}{
		{name: "len", value: "12345", condName: "len", condValue: "5"},
		{name: "len_with_spaces", value: "12345", condName: "len", condValue: " 5 "},
		{name: "regexp", value: "12345", condName: "regexp", condValue: "\\d{5}"},
		{name: "in", value: "12345", condName: "in", condValue: "qwerty,12345,asdfg"},
		{name: "in_with_spaces", value: "12345", condName: "in", condValue: "qwerty , 12345 , asdfg"},
		{name: "in_one_value", value: "12345", condName: "in", condValue: "12345"},
	}
	for _, tc := range tasks {
		tc := tc
		t.Run("success_"+tc.name, func(t *testing.T) {
			fieldValueMock := mocks.NewStringInterface(t)
			fieldValueMock.EXPECT().String().Return(tc.value)

			err := StringValidator{}.Validate(fieldValueMock, tc.condName, tc.condValue)
			require.NoError(t, err)
		})
	}
}

func TestStringValidatorCondValueWrongFormat(t *testing.T) {
	tasks := []struct {
		name      string
		condName  string
		condValue string
	}{
		{name: "len", condName: "len", condValue: "wrong_format"},
		{name: "regexp", condName: "regexp", condValue: "\\+"},
	}
	for _, tc := range tasks {
		tc := tc
		t.Run("condValue_wrong_format_"+tc.name, func(t *testing.T) {
			fieldValueMock := mocks.NewStringInterface(t)
			fieldValueMock.EXPECT().String().Return("12345")

			err := StringValidator{}.Validate(fieldValueMock, tc.condName, tc.condValue)
			require.Error(t, err)
		})
	}
}

func TestStringValidatorInvalidValue(t *testing.T) {
	tasks := []struct {
		name      string
		value     string
		condName  string
		condValue string
	}{
		{name: "len_lt", value: "1234", condName: "len", condValue: "5"},
		{name: "len_gt", value: "123456", condName: "len", condValue: "5"},
		{name: "regexp", value: "qwerty", condName: "regexp", condValue: "\\d+"},
		{name: "in", value: "12345", condName: "in", condValue: "qwerty,asdfg"},
	}
	for _, tc := range tasks {
		tc := tc
		t.Run("invalid_value_"+tc.name, func(t *testing.T) {
			fieldValueMock := mocks.NewStringInterface(t)
			fieldValueMock.EXPECT().String().Return(tc.value)

			err := StringValidator{}.Validate(fieldValueMock, tc.condName, tc.condValue)
			require.Error(t, err)
		})
	}
}
