package request

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBatCaveSettingsRequestValidate(t *testing.T) {
	testCases := []struct {
		name string
		req  CreateBatCaveSettingsRequest
	}{
		{
			name: "Valid fields in request",
			req: CreateBatCaveSettingsRequest{
				DeviceID:       "34ee5c9a4411",
				DeepSleepDelay: 10,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			validationErrors := testCase.req.Validate()
			errs := map[string]interface{}{"validationError": validationErrors}
			errorMessage := fmt.Sprintf("should have no errors: %v", errs)
			assert.Equal(t, 0, len(validationErrors), errorMessage)
		})
	}
}

func TestCreateBatCaveSettingsRequestValidateWhenDeviceIDIsNotValid(t *testing.T) {
	testCases := []struct {
		name           string
		req            CreateBatCaveSettingsRequest
		expectedErrors string
	}{
		{
			name: "DeviceID length is not 12 characters long",
			req: CreateBatCaveSettingsRequest{
				DeviceID:       "34e5c9a4411",
				DeepSleepDelay: 10,
			},

			// FIXME: add errors
			expectedErrors: "",
		},
		{
			name: "DeviceID is not a valid mac address",
			req: CreateBatCaveSettingsRequest{
				DeviceID:       "44cbagbe2e4f",
				DeepSleepDelay: 15,
			},

			// FIXME: add errors
			expectedErrors: "",
		},
		{
			name: "DeviceID is empty",
			req: CreateBatCaveSettingsRequest{
				DeviceID:       "",
				DeepSleepDelay: 20,
			},

			// FIXME: add errors
			expectedErrors: "",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			validationErrors := testCase.req.Validate()
			errs := map[string]interface{}{"validationError": validationErrors}
			errorMessage := fmt.Sprintf("should have errors: %s", testCase.expectedErrors)
			assert.Equal(t, testCase.expectedErrors, errs, errorMessage)
		})
	}
}

// FIXME: implement
func TestCreateBatCaveSettingsRequestValidateWhenDeepSleepDelayIsNotValid(t *testing.T) {
}
