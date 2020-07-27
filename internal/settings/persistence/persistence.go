package persistence

type Persistence interface {
	CreateBatCaveSetting(settings BatCaveSetting)
	GetBatCaveSetting(deviceID string) (bool, BatCaveSetting)
	UpdateBatCaveSetting(settings BatCaveSetting)
}
