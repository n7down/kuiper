package persistence

import "time"

type BatCaveSetting struct {
	DeviceID       string `gorm:"primary_key"`
	DeepSleepDelay int32
	CreatedAt      *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"index" json:"deleted_at"`
}
