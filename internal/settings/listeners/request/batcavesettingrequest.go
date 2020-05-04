package request

import (
	"github.com/n7down/kuiper/internal/settings/listeners/commands"
	"github.com/n7down/kuiper/internal/settings/listeners/response"
	"github.com/n7down/kuiper/internal/settings/persistence"
)

type BatCaveSettingRequest struct {
	DeviceID       string `json:"m"`
	DeepSleepDelay uint32 `json:"s"`
}

func (s *BatCaveSettingRequest) IsEqual(settings persistence.BatCaveSetting) (bool, response.BatCaveSettingResponse) {
	res := response.BatCaveSettingResponse{}
	isEqual := true

	if s.DeepSleepDelay != settings.DeepSleepDelay {
		isEqual = false
		res.DeepSleepDelay = settings.DeepSleepDelay
	} else {
		res.DeepSleepDelay = s.DeepSleepDelay
	}

	return isEqual, res
}

func (s *BatCaveSettingRequest) IsEqualAndGetCommands(settings persistence.BatCaveSetting) (bool, []string) {
	commands := commands.BatCaveSettingCommands{}
	hasChanges := false

	if s.DeepSleepDelay != settings.DeepSleepDelay {
		hasChanges = true
		commands.AddDeepSleepDelayCommand(settings.DeepSleepDelay)
	}

	return hasChanges, commands.GetCommands()
}
