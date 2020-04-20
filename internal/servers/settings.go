package servers

import (
	"context"

	_ "github.com/go-sql-driver/mysql"

	settings_pb "github.com/n7down/iota/internal/pb/settings"
	"github.com/n7down/iota/internal/persistence"
	"github.com/n7down/iota/internal/persistence/mysql"
)

type SettingsServer struct {
	db *mysql.SettingsMySqlDB
}

func NewSettingsServer(db *mysql.SettingsMySqlDB) *SettingsServer {
	return &SettingsServer{
		db: db,
	}
}

func (s *SettingsServer) SetBatCaveSettings(ctx context.Context, req *settings_pb.SetBatCaveSettingsRequest) (*settings_pb.SetBatCaveSettingsResponse, error) {
	settings := persistence.UpdateBatCaveSettings{
		DeepSleepDelay: req.DeepSleepDelay,
	}

	err := s.db.UpdateBatCaveSettings(req.DeviceID, settings)
	if err != nil {
		return &settings_pb.SetBatCaveSettingsResponse{}, err
	}

	return &settings_pb.SetBatCaveSettingsResponse{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}, nil
}

func (s *SettingsServer) GetBatCaveSettings(ctx context.Context, req *settings_pb.GetBatCaveSettingsRequest) (*settings_pb.GetBatCaveSettingsResponse, error) {
	settings, err := s.db.GetBatCaveSettings(req.DeviceID)
	if err != nil {
		return &settings_pb.GetBatCaveSettingsResponse{}, err
	}

	return &settings_pb.GetBatCaveSettingsResponse{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: settings.DeepSleepDelay,
	}, nil
}
