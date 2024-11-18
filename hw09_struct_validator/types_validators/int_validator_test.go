package types_validators

import (
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/types_validators/mocks"
	"github.com/stretchr/testify/require"
)

func TestSuccessValidateMinMax(t *testing.T) {
	tasks := []struct {
		name      string
		value     int64
		condName  string
		condValue string
	}{
		{name: "min", value: 10, condName: "min", condValue: "5"},
		{name: "max", value: 10, condName: "max", condValue: "15"},
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

func TestCondValueWrongFormat(t *testing.T) {
	tasks := []struct {
		name     string
		condName string
	}{
		{name: "min", condName: "min"},
		{name: "max", condName: "max"},
	}

	for _, tc := range tasks {
		tc := tc
		t.Run("condValue_wrong_format_"+tc.name, func(t *testing.T) {
			err := IntValidator{}.validateMinMax(1, tc.condName, "wrong_format")
			require.Error(t, err)
		})
	}

}
