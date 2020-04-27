package request

import "github.com/n7down/kuiper/internal/settings/persistence"

type BatCaveSettingRequest struct {
	DeviceID       string `json:"m"`
	DeepSleepDelay int32  `json:"s"`
}

func (s *BatCaveSettingRequest) IsEqual(settings persistence.BatCaveSetting) bool {
	if s.DeepSleepDelay == settings.DeepSleepDelay {
		return true
	}
	return false
}
