//+build unit

package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_Bat_Cave_Setting_Commands_Get_Commands(t *testing.T) {
	testCases := []struct {
		name           string
		deepSleepDelay int32
		expectedValue  []string
	}{
		{
			name:           "Deep_Sleep_Delay_Set_To_1",
			deepSleepDelay: 1,
			expectedValue: []string{
				"00000001",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := BatCaveSettingCommands{}
			c.AddDeepSleepDelayCommand(testCase.deepSleepDelay)
			commands := c.GetCommands()
			assert.Equal(t, testCase.expectedValue, commands, "should have the same commands")
		})
	}
}
