package listeners

import (
	"github.com/n7down/kuiper/internal/settings/persistence"

	logger "github.com/n7down/kuiper/internal/logger"
)

type SettingsListenersEnv struct {
	persistence persistence.Persistence
	logger      logger.Logger
}

func NewSettingsListenersEnv(persistence persistence.Persistence, logger logger.Logger) *SettingsListenersEnv {
	return &SettingsListenersEnv{
		persistence: persistence,
		logger:      logger,
	}
}
