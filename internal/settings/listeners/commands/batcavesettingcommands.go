package commands

import (
	"fmt"
)

const (
	DEEP_SLEEP_DELAY_COMMAND = 0x00
)

type BatCaveSettingCommands struct {
	commands []string
}

func (c *BatCaveSettingCommands) GetCommands() []string {
	return c.commands
}

func (c *BatCaveSettingCommands) AddDeepSleepDelayCommand(d int32) {
	h := fmt.Sprintf("%04x%04x", DEEP_SLEEP_DELAY_COMMAND, d)
	c.commands = append(c.commands, h)
}
