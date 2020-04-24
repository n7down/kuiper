package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBatCaveSettingsRequestValidate(t *testing.T) {
	testCases := []CreateBatCaveSettingsRequest{
		CreateBatCaveSettingsRequest{
			DeviceID:       "34ee5c9a4411",
			DeepSleepDelay: 10,
		},
		CreateBatCaveSettingsRequest{
			DeviceID:       "44cba9be2e4f",
			DeepSleepDelay: 15,
		},
		CreateBatCaveSettingsRequest{
			DeviceID:       "c0b3a5ee334b",
			DeepSleepDelay: 15,
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
