package mysql

import (
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	grom "github.com/jinzhu/gorm"

	"github.com/n7down/kuiper/internal/settings/persistence"
)

type SettingsMySqlDB struct {
	db *grom.DB
}

// func NewSettingsMySqlDB(dbConn string) (*SettingsMySqlDB, error) {
// db, err := sql.Open("mysql", dbConn)
// if err != nil {
// 	return &SettingsMySqlDB{}, err
// }

// err = db.Ping()
// if err != nil {
// 	return &SettingsMySqlDB{}, err
// }

// return &SettingsMySqlDB{
// 	db: db,
// }, nil
// }

// FIXME: this is untested
// func NewSettingsMySqlDBWithURL(url *url.URL) (*SettingsMySqlDB, error) {
// dbUser := url.User.Username()
// dbPass, _ := url.User.Password()

// dbName := url.Path[1:len(url.Path)]
// if dbName == "" {
// 	dbName = "test"
// }

// dbConn := fmt.Sprintf("%s:%s@%s(%s)/%s", dbUser, dbPass, url.Scheme, url.Host, dbName)

// db, err := sql.Open("mysql", dbConn)
// if err != nil {
// 	return &SettingsMySqlDB{}, err
// }

// err = db.Ping()
// if err != nil {
// 	return &SettingsMySqlDB{}, err
// }

// return &SettingsMySqlDB{
// 	db: db,
// }, nil
// }

// func (s *SettingsMySqlDB) UpdateBatCaveSettings(deviceID string, settings persistence.UpdateBatCaveSettings) error {
// 	query := `INSERT INTO bat_cave_settings (device_id, deep_sleep_delay, updated) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE SET deep_sleep_delay=? updated=? WHERE device_id=?`
// 	_, err := s.db.Exec(query, deviceID, settings.DeepSleepDelay, time.Now(), settings.DeepSleepDelay, time.Now(), deviceID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *SettingsMySqlDB) GetBatCaveSettings(deviceID string) (persistence.GetBatCaveSettings, error) {
// 	settings := persistence.GetBatCaveSettings{}
// 	query := `SELECT deep_sleep_delay FROM bat_cave_settings WHERE device_id=?`
// 	row := s.db.QueryRow(query, deviceID)
// 	err := row.Scan(&settings.DeepSleepDelay)
// 	if err != nil {
// 		return settings, err
// 	}
// 	return settings, nil
// }

func NewSettingsMySqlDB(dbConn string) (*SettingsMySqlDB, error) {
	db, err := grom.Open("mysql", dbConn)
	if err != nil {
		return &SettingsMySqlDB{}, err
	}
	return &SettingsMySqlDB{
		db: db,
	}, nil
}

// FIXME: this is untested
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

	return &SettingsMySqlDB{
		db: db,
	}, nil
}

func (s *SettingsMySqlDB) CreateBatCaveSetting(settings persistence.BatCaveSetting) error {
	s.db.Create(settings)
	return nil
}

func (s *SettingsMySqlDB) GetBatCaveSetting(deviceID string) (persistence.BatCaveSetting, error) {
	var settings persistence.BatCaveSetting
	s.db.Where("device_id=?", deviceID).First(&settings)
	return settings, nil
}

func (s *SettingsMySqlDB) UpdateBatCaveSetting(settings persistence.BatCaveSetting) persistence.BatCaveSetting {
	s.db.Model(&settings).Where("device_id = ?", settings.DeviceID).Updates(persistence.BatCaveSetting{DeepSleepDelay: settings.DeepSleepDelay})
	return settings
}
