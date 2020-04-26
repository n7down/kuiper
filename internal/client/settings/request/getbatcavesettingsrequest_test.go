// +build unit

package request

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_Bat_Cave_Setting_Request_Validate(t *testing.T) {
	testCases := []struct {
		name           string
		req            GetBatCaveSettingRequest
		expectedErrors int
	}{
		{
			name: "Valid_Fields_In_Request",
			req: GetBatCaveSettingRequest{
				DeviceID: "34ee5c9a4411",
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

func Test_Get_Bat_Cave_Setting_Request_Validate_When_Device_ID_Is_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            GetBatCaveSettingRequest
		expectedErrors map[string]interface{}
	}{
		{
			name: "DeviceID_Length_Is_Not_12_Characters_Long",
			req: GetBatCaveSettingRequest{
				DeviceID: "34e5c9a4411",
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
			req: GetBatCaveSettingRequest{
				DeviceID: "44cbagbe2e4f",
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
			req: GetBatCaveSettingRequest{
				DeviceID: "",
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
