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
		{
			name:           "Deep_Sleep_Delay_Set_To_65535",
			deepSleepDelay: 65535,
			expectedValue: []string{
				"0000ffff",
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

func Test_Get_Bat_Cave_Setting_Commands_Get_Commands_Int(t *testing.T) {
	testCases := []struct {
		name           string
		deepSleepDelay int32
		expectedValue  []int
	}{
		{
			name:           "Deep_Sleep_Delay_Set_To_1",
			deepSleepDelay: 1,
			expectedValue: []int{
				1,
			},
		},
		{
			name:           "Deep_Sleep_Delay_Set_To_65535",
			deepSleepDelay: 65535,
			expectedValue: []int{
				65535,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := BatCaveSettingCommands{}
			c.AddDeepSleepDelayCommand(testCase.deepSleepDelay)
			commands := c.GetCommandsInt()
			assert.Equal(t, testCase.expectedValue, commands, "should have the same commands")
		})
	}
}
