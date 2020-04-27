package listeners

import "github.com/n7down/kuiper/internal/settings/persistence/mysql"

type SettingsListenerEnv struct {
	db *mysql.SettingsMySqlDB
}

func NewSettingsListenerEnv(db *mysql.SettingsMySqlDB) *SettingsListenerEnv {
	return &SettingsListenerEnv{
		db: db,
	}
}
