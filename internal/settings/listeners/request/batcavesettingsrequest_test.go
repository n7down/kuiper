package request

import (
	"testing"

	"github.com/n7down/kuiper/internal/settings/listeners/response"
	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/stretchr/testify/assert"
)

func Test_Get_Bat_Cave_Setting_Request_Is_Equal(t *testing.T) {
	testCases := []struct {
		name            string
		req             BatCaveSettingRequest
		persistence     persistence.BatCaveSetting
		expectedValue   bool
		expectedSetting response.BatCaveSettingResponse
	}{
		{
			name: "Deep_Sleep_Delay_Are_Equal",
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
		{
			name: "Deep_Sleep_Delay_Has_Changes_In_Persistence",
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
