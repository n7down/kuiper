package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BatCaveSettingCommands_GetCommands_Should_Set_Commands_When_DeepSleepDelay_Values_Are_Added(t *testing.T) {
	testCases := []struct {
		name           string
		deepSleepDelay uint32
		expectedValue  []string
	}{
		{
			name:           "DeepSleepDelay_Set_To_1",
			deepSleepDelay: 1,
			expectedValue: []string{
				"00000001",
			},
		},
		{
			name:           "DeepSleepDelay_Set_To_65535",
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

func Test_BatCaveSettingCommands_GetCommandsInt_Should_Set_Commands_When_DeepSleepDelay_Values_Are_Added(t *testing.T) {
	testCases := []struct {
		name           string
		deepSleepDelay uint32
		expectedValue  []int
	}{
		{
			name:           "DeepSleepDelay_Set_To_1",
			deepSleepDelay: 1,
			expectedValue: []int{
				1,
			},
		},
		{
			name:           "DeepSleepDelay_Set_To_65535",
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
