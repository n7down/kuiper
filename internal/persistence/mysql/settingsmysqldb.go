package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/n7down/iota/internal/persistence"
)

type SettingsMySqlDB struct {
	db *sql.DB
}

func NewSettingsMySqlDB(dbUser, dbPass, dbSocket, dbHost, dbName string) (*SettingsMySqlDB, error) {
	dbInfo := fmt.Sprintf("%s:%s@%s(%s)/%s", dbUser, dbPass, dbSocket, dbHost, dbName)
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		return &SettingsMySqlDB{}, err
	}

	err = db.Ping()
	if err != nil {
		return &SettingsMySqlDB{}, err
	}

	return &SettingsMySqlDB{
		db: db,
	}, nil
}

func (s *SettingsMySqlDB) UpdateBatCaveSettings(deviceID string, settings persistence.UpdateBatCaveSettings) error {
	// query := `INSERT INTO settings SET device_id=?, deep_sleep_delay=?, updated=? ON DUPLICATE UPDATE settings SET deep_sleep_delay=?, updated=? WHERE device_id=?`
	query := `INSERT INTO bat_cave_settings (device_id, deep_sleep_delay, updated) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE SET deep_sleep_delay=? updated=? WHERE device_id=?`
	_, err := s.db.Exec(query, deviceID, settings.DeepSleepDelay, time.Now(), settings.DeepSleepDelay, time.Now(), deviceID)
	if err != nil {
		return err
	}
	return nil
}

func (s *SettingsMySqlDB) GetBatCaveSettings(deviceID string) (persistence.GetBatCaveSettings, error) {
	settings := persistence.GetBatCaveSettings{}
	query := `SELECT deep_sleep_delay, updated FROM bat_cave_settings WHERE device_id=?`
	row := s.db.QueryRow(query)
	err := row.Scan(&settings.DeepSleepDelay, &settings.Updated)
	if err != nil {
		return settings, err
	}
	return settings, nil
}
