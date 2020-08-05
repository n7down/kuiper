// +build unit,!integration

package request

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBatCaveSettingRequest_Should_Return_Error_When_DeviceID_Field_Is_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            GetBatCaveSettingRequest
		expectedErrors map[string]interface{}
	}{
		{
			name: "DeviceID_Length_Is_Greater_Then_12_Characters_Long",
			req: GetBatCaveSettingRequest{
				DeviceID: "34e5c9a441111",
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
			name: "DeviceID_Contains_An_Invalid_Mac_Address_Character",
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
