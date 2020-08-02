package request

import (
	"testing"

	"github.com/n7down/kuiper/internal/settings/listeners/response"
	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/stretchr/testify/assert"
)

func Test_BatCaveSettingRequest_IsEqual_Should_Return_Changes_When_BatCaveSetting_And_Persistence_Are_Different(t *testing.T) {
	testCases := []struct {
		name            string
		req             BatCaveSettingRequest
		persistence     persistence.BatCaveSetting
		expectedValue   bool
		expectedSetting response.BatCaveSettingResponse
	}{
		{
			name: "DeepSleepDelay_Has_Changes_In_Persistence",
			req: BatCaveSettingRequest{
				DeepSleepDelay: 15,
			},
			persistence: persistence.BatCaveSetting{
				DeepSleepDelay: 20,
			},
			expectedValue: false,
			expectedSetting: response.BatCaveSettingResponse{
				DeepSleepDelay: 20,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			isEqual, res := testCase.req.IsEqual(testCase.persistence)
			assert.Equal(t, testCase.expectedValue, isEqual, "should have the same boolean value")
			assert.Equal(t, testCase.expectedSetting, res, "should have the same setting")
		})
	}
}

func Test_BatCaveSettingRequest_IsEqual_Should_Return_Empty_String_When_BatCaveSetting_And_Persistence_Are_The_Same(t *testing.T) {
	testCases := []struct {
		name            string
		req             BatCaveSettingRequest
		persistence     persistence.BatCaveSetting
		expectedValue   bool
		expectedSetting response.BatCaveSettingResponse
	}{
		{
			name: "DeepSleepDelay_Are_Equal",
			req: BatCaveSettingRequest{
				DeepSleepDelay: 15,
			},
			persistence: persistence.BatCaveSetting{
				DeepSleepDelay: 15,
			},
			expectedValue: true,
			expectedSetting: response.BatCaveSettingResponse{
				DeepSleepDelay: 15,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			isEqual, res := testCase.req.IsEqual(testCase.persistence)
			assert.Equal(t, testCase.expectedValue, isEqual, "should have the same boolean value")
			assert.Equal(t, testCase.expectedSetting, res, "should have the same setting")
		})
	}
}

func Test_BatCaveSettingRequest_IsEqualAndGetCommands_Should_Return_Empty_String_When_BatCaveSetting_And_Persistence_Are_The_Same(t *testing.T) {
	testCases := []struct {
		name               string
		req                BatCaveSettingRequest
		persistence        persistence.BatCaveSetting
		expectedHasChanges bool
		expectedCommands   []string
	}{
		{
			name: "DeepSleepDelay_Has_Changes_In_Persistence",
			req: BatCaveSettingRequest{
				DeepSleepDelay: 15,
			},
			persistence: persistence.BatCaveSetting{
				DeepSleepDelay: 1,
			},
			expectedHasChanges: true,
			expectedCommands: []string{
				"00000001",
			},
		},
		{
			name: "DeepSleepDelay_Has_Changes_And_Is_Max_Value",
			req: BatCaveSettingRequest{
				DeepSleepDelay: 15,
			},
			persistence: persistence.BatCaveSetting{
				DeepSleepDelay: 65535,
			},
			expectedHasChanges: true,
			expectedCommands: []string{
				"0000ffff",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			hasChanges, commands := testCase.req.IsEqualAndGetCommands(testCase.persistence)
			assert.Equal(t, testCase.expectedHasChanges, hasChanges, "should have the same boolean value")
			assert.Equal(t, testCase.expectedCommands, commands, "should have the same setting")
		})
	}
}

func Test_BatCaveSettingRequest_IsEqualAndGetCommands_Should_Return_Command_String_When_BatCaveSetting_And_Persistence_Are_Different(t *testing.T) {
	testCases := []struct {
		name               string
		req                BatCaveSettingRequest
		persistence        persistence.BatCaveSetting
		expectedHasChanges bool
		expectedCommands   []string
	}{
		{
			name: "DeepSleepDelay_Are_Equal",
			req: BatCaveSettingRequest{
				DeepSleepDelay: 15,
			},
			persistence: persistence.BatCaveSetting{
				DeepSleepDelay: 15,
			},
			expectedHasChanges: false,
			expectedCommands:   []string(nil),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			hasChanges, commands := testCase.req.IsEqualAndGetCommands(testCase.persistence)
			assert.Equal(t, testCase.expectedHasChanges, hasChanges, "should have the same boolean value")
			assert.Equal(t, testCase.expectedCommands, commands, "should have the same setting")
		})
	}
}
