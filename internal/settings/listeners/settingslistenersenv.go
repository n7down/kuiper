package listeners

import "github.com/n7down/kuiper/internal/settings/persistence/mysql"

type SettingsListenersEnv struct {
	db *mysql.SettingsMySqlDB
}

func NewSettingsListenersEnv(db *mysql.SettingsMySqlDB) *SettingsListenersEnv {
	return &SettingsListenersEnv{
		db: db,
	}
}
