package types_validators

import (
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/types_validators/mocks"
	"github.com/stretchr/testify/require"
)

func TestIntValidatorSuccess(t *testing.T) {
	tasks := []struct {
		name      string
		value     int64
		condName  string
		condValue string
	}{
		{name: "min", value: 10, condName: "min", condValue: "5"},
		{name: "boundary_min", value: 5, condName: "min", condValue: "5"},
		{name: "max", value: 10, condName: "max", condValue: "15"},
		{name: "boundary_max", value: 15, condName: "max", condValue: "15"},
		{name: "unknown_cond_name", value: 10, condName: "unknonw", condValue: "wrong_format"},
		{name: "in_without_any_spaces", value: 10, condName: "in", condValue: "5,15"},
		{name: "in_with_spaces", value: 10, condName: "in", condValue: "5 , 15"},
		{name: "in_min", value: 5, condName: "in", condValue: "5,15"},
		{name: "in_max", value: 15, condName: "in", condValue: "5,15"},
	}
	for _, tc := range tasks {
		tc := tc
		t.Run("success_"+tc.name, func(t *testing.T) {
			fieldValueMock := mocks.NewIntInterface(t)
			fieldValueMock.EXPECT().Int().Return(tc.value)

			err := IntValidator{}.Validate(fieldValueMock, tc.condName, tc.condValue)
			require.NoError(t, err)
		})
	}
}

func TestIntValidatorCondValueWrongFormat(t *testing.T) {
	tasks := []struct {
		name      string
		condName  string
		condValue string
	}{
		{name: "min", condName: "min", condValue: "wrong_format"},
		{name: "max", condName: "max", condValue: "wrong_format"},
		{name: "in", condName: "in", condValue: "wrong_format"},
		{name: "in_min", condName: "in", condValue: "qw,5"},
		{name: "in_max", condName: "in", condValue: "5,qw"},
	}

	for _, tc := range tasks {
		tc := tc
		t.Run("condValue_wrong_format_"+tc.name, func(t *testing.T) {
			fieldValueMock := mocks.NewIntInterface(t)
			fieldValueMock.EXPECT().Int().Return(1)

			err := IntValidator{}.Validate(fieldValueMock, tc.condName, tc.condValue)
			require.Error(t, err)
		})
	}
}

func TestIntValidatorInvalidValue(t *testing.T) {
	tasks := []struct {
		name      string
		value     int64
		condName  string
		condValue string
	}{
		{name: "min", value: 1, condName: "min", condValue: "5"},
		{name: "max", value: 10, condName: "max", condValue: "5"},
		{name: "in_lt_min", value: 1, condName: "in", condValue: "5,10"},
		{name: "in_gt_max", value: 11, condName: "in", condValue: "5,10"},
	}

	for _, tc := range tasks {
		tc := tc
		t.Run("invalid_value_"+tc.name, func(t *testing.T) {
			fieldValueMock := mocks.NewIntInterface(t)
			fieldValueMock.EXPECT().Int().Return(tc.value)

			err := IntValidator{}.Validate(fieldValueMock, tc.condName, tc.condValue)
			require.Error(t, err)
		})
	}
}
