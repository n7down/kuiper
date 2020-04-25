package servers

import (
	"context"
	"time"

	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/n7down/kuiper/internal/settings/persistence/mysql"
)

type SettingsServer struct {
	db *mysql.SettingsMySqlDB
}

func NewSettingsServer(db *mysql.SettingsMySqlDB) *SettingsServer {
	return &SettingsServer{
		db: db,
	}
}

func (s *SettingsServer) CreateBatCaveSettings(ctx context.Context, req *settings_pb.CreateBatCaveSettingsRequest) (*settings_pb.CreateBatCaveSettingsResponse, error) {
	settings := persistence.BatCaveSettings{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
		Updated:        time.Now(),
	}

	err := s.db.CreateBatCaveSettings(settings)
	if err != nil {
		return &settings_pb.CreateBatCaveSettingsResponse{}, err
	}

	return &settings_pb.CreateBatCaveSettingsResponse{}, nil
}

func (s *SettingsServer) UpdateBatCaveSettings(ctx context.Context, req *settings_pb.UpdateBatCaveSettingsRequest) (*settings_pb.UpdateBatCaveSettingsResponse, error) {
	settings := persistence.BatCaveSettings{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
		Updated:        time.Now(),
	}

	s.db.UpdateBatCaveSettings(settings)

	return &settings_pb.UpdateBatCaveSettingsResponse{
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
		DeviceID:       settings.DeviceID,
		DeepSleepDelay: settings.DeepSleepDelay,
	}, nil
}
