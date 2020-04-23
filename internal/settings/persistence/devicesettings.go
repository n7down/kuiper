package persistence

import "time"

type BatCaveSettings struct {
	DeviceID       string
	DeepSleepDelay int32
	Updated        time.Time
}
