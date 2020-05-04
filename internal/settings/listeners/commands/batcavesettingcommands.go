package commands

import (
	"fmt"
	"strconv"
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

func (c *BatCaveSettingCommands) GetCommandsInt() []int {
	commandsInt := []int{}
	for _, command := range c.commands {
		commandInt, _ := strconv.ParseInt(command, 16, 64)
		commandsInt = append(commandsInt, int(commandInt))
	}
	return commandsInt
}

func (c *BatCaveSettingCommands) AddDeepSleepDelayCommand(d uint32) {
	h := fmt.Sprintf("%04x%04x", DEEP_SLEEP_DELAY_COMMAND, d)
	c.commands = append(c.commands, h)
}
