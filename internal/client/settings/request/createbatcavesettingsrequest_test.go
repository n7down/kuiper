package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBatCaveSettingsRequestValidateWhenDeviceIDIsValid(t *testing.T) {
	testCases := []CreateBatCaveSettingsRequest{
		CreateBatCaveSettingsRequest{
			DeviceID:       "34ee5c9a4411",
			DeepSleepDelay: 10,
		},
	}

	for _, testCase := range testCases {
		validationErrors := testCase.Validate()
		errs := map[string]interface{}{"validationError": validationErrors}
		assert.Equal(t, len(testCase.Validate()), 0, errs)
	}
}

// func TestCreateBatCaveSettingsRequestValidateWhenDeviceIDIsNotValid(t *testing.T) {
// }
