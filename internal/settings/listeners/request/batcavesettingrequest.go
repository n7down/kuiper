package request

import (
	"github.com/n7down/kuiper/internal/settings/listeners/commands"
	"github.com/n7down/kuiper/internal/settings/persistence"
)

type BatCaveSettingRequest struct {
	DeviceID       string `json:"m"`
	DeepSleepDelay int32  `json:"s"`
}

func (s *BatCaveSettingRequest) IsEqual(settings persistence.BatCaveSetting) (bool, []string) {
	c := commands.BatCaveSettingCommands{}
	hasChanges := false

	if s.DeepSleepDelay != settings.DeepSleepDelay {
		hasChanges = true
		c.AddDeepSleepDelayCommand(settings.DeepSleepDelay)
	}

	return hasChanges, c.GetCommands()
}
