package mysql

import "github.com/n7down/kuiper/internal/settings/persistence"

func (p *MysqlPersistence) CreateBatCaveSetting(settings persistence.BatCaveSetting) int64 {
	rowsAffected := p.db.Create(&settings).RowsAffected
	return rowsAffected
}

func (p *MysqlPersistence) GetBatCaveSetting(deviceID string) (bool, persistence.BatCaveSetting) {
	var settings persistence.BatCaveSetting
	recordNotFound := p.db.Where("device_id=?", deviceID).First(&settings).RecordNotFound()
	return recordNotFound, settings
}

func (p *MysqlPersistence) UpdateBatCaveSetting(settings persistence.BatCaveSetting) int64 {
	rowsAffected := p.db.Model(&settings).Where("device_id = ?", settings.DeviceID).Updates(persistence.BatCaveSetting{DeepSleepDelay: settings.DeepSleepDelay}).RowsAffected
	return rowsAffected
}
