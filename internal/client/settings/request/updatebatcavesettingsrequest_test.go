package request

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Update_Bat_Cave_Settings_Request_Validate(t *testing.T) {
	testCases := []struct {
		name           string
		req            UpdateBatCaveSettingsRequest
		expectedErrors int
	}{
		{
			name: "Valid_Fields_In_Request",
			req: UpdateBatCaveSettingsRequest{
				DeviceID:       "34ee5c9a4411",
				DeepSleepDelay: 10,
			},
			expectedErrors: 0,
		},
		{
			name: "Deep_Sleep_Is_Fields_Is_1",
			req: UpdateBatCaveSettingsRequest{
				DeviceID:       "34ee5c9a4411",
				DeepSleepDelay: 1,
			},
			expectedErrors: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			validationErrors := testCase.req.Validate()
			errs := map[string]interface{}{"validationError": validationErrors}
			errorMessage := fmt.Sprintf("should have no errors: %v", errs)
			assert.Equal(t, testCase.expectedErrors, len(validationErrors), errorMessage)
		})
	}
}

func Test_Update_Bat_Cave_Settings_Request_Validate_When_Device_ID_Is_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            UpdateBatCaveSettingsRequest
		expectedErrors map[string]interface{}
	}{
		{
			name: "DeviceID_Length_Is_Not_12_Characters_Long",
			req: UpdateBatCaveSettingsRequest{
				DeviceID:       "34e5c9a4411",
				DeepSleepDelay: 10,
			},
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"deviceID": []string{
						"The deviceID field needs to be a valid mac!",
					},
				},
			},
		},
		{
			name: "DeviceID_Is_Not_A_Valid_Mac_Address",
			req: UpdateBatCaveSettingsRequest{
				DeviceID:       "44cbagbe2e4f",
				DeepSleepDelay: 15,
			},
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"deviceID": []string{
						"The deviceID field needs to be a valid mac!",
					},
				},
			},
		},
		{
			name: "DeviceID_Is_Empty",
			req: UpdateBatCaveSettingsRequest{
				DeviceID:       "",
				DeepSleepDelay: 20,
			},
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"deviceID": []string{
						"The deviceID field needs to be a valid mac!",
					},
				},
			},
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

func Test_Update_Bat_Cave_Settings_Request_Validate_When_Deep_Sleep_Delay_Is_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            UpdateBatCaveSettingsRequest
		expectedErrors map[string]interface{}
	}{
		{
			name: "Deep_Sleep_Delay_Equals 0",
			req: UpdateBatCaveSettingsRequest{
				DeviceID:       "123456789aae",
				DeepSleepDelay: 0,
			},
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"deepSleepDelay": []string{
						"The deepSleepDelay field should be a positive non-zero value!",
					},
				},
			},
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
