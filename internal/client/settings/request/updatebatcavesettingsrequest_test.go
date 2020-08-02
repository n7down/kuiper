package request

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateBatCaveSettingRequest_Validate_Should_Return_Error_When_DeviceID_Field_Is_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            UpdateBatCaveSettingRequest
		expectedErrors map[string]interface{}
	}{
		{
			name: "DeviceID_Length_Is_Greater_Then_12_Characters_Long",
			req: UpdateBatCaveSettingRequest{
				DeviceID:       "34e5c9a441111",
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
			name: "DeviceID_Length_Is_Less_Then_12_Characters_Long",
			req: UpdateBatCaveSettingRequest{
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
			name: "DeviceID_Contains_An_Invalid_Mac_Address_Charater",
			req: UpdateBatCaveSettingRequest{
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
			req: UpdateBatCaveSettingRequest{
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

func Test_UpdateBatCaveSettingRequest_Validate_Should_Return_Error_When_DeepSleepDelayIs_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            UpdateBatCaveSettingRequest
		expectedErrors map[string]interface{}
	}{
		{
			name: "Deep_Sleep_Delay_Equals 0",
			req: UpdateBatCaveSettingRequest{
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
