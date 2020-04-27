package mysql

import (
	"fmt"
	"net/url"

	grom "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/n7down/kuiper/internal/settings/persistence"
)

type SettingsMySqlDB struct {
	db *grom.DB
}

func NewSettingsMySqlDB(dbConn string) (*SettingsMySqlDB, error) {
	db, err := grom.Open("mysql", dbConn)
	if err != nil {
		return &SettingsMySqlDB{}, err
	}

	err = db.DB().Ping()
	if err != nil {
		return &SettingsMySqlDB{}, err
	}

	return &SettingsMySqlDB{
		db: db,
	}, nil
}

// !!!: this is untested
func NewSettingsMySqlDBWithURL(url *url.URL) (*SettingsMySqlDB, error) {
	dbUser := url.User.Username()
	dbPass, _ := url.User.Password()

	dbName := url.Path[1:len(url.Path)]
	if dbName == "" {
		dbName = "test"
	}

	dbConn := fmt.Sprintf("%s:%s@%s(%s)/%s", dbUser, dbPass, url.Scheme, url.Host, dbName)

	db, err := grom.Open("mysql", dbConn)
	if err != nil {
		return &SettingsMySqlDB{}, err
	}

	err = db.DB().Ping()
	if err != nil {
		return &SettingsMySqlDB{}, err
	}

	return &SettingsMySqlDB{
		db: db,
	}, nil
}

func (s *SettingsMySqlDB) CreateBatCaveSetting(settings persistence.BatCaveSetting) {
	s.db.Create(&settings)
}

func (s *SettingsMySqlDB) GetBatCaveSetting(deviceID string) (bool, persistence.BatCaveSetting) {
	var settings persistence.BatCaveSetting
	recordNotFound := s.db.Where("device_id=?", deviceID).First(&settings).RecordNotFound()
	return recordNotFound, settings
}

func (s *SettingsMySqlDB) UpdateBatCaveSetting(settings persistence.BatCaveSetting) {
	s.db.Model(&settings).Where("device_id = ?", settings.DeviceID).Updates(persistence.BatCaveSetting{DeepSleepDelay: settings.DeepSleepDelay})
}
