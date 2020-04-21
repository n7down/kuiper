package persistence

import "time"

type GetBatCaveSettings struct {
	DeepSleepDelay int32
	Updated        time.Time
}

type UpdateBatCaveSettings struct {
	DeepSleepDelay int32 `json:"deepSleepDelay"`
}
