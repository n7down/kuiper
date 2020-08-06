//go:generate mockgen -source persistence.go -destination=mock/mockpersistence.go -package=mock
package persistence

type Persistence interface {
	CreateBatCaveSetting(settings BatCaveSetting) int64
	GetBatCaveSetting(deviceID string) (bool, BatCaveSetting)
	UpdateBatCaveSetting(settings BatCaveSetting) int64
}
