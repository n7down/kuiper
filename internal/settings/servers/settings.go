package servers

import (
	"context"

	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/n7down/kuiper/internal/settings/persistence/mysql"
)

type SettingServer struct {
	db *mysql.SettingsMySqlDB
}

func NewSettingsServer(db *mysql.SettingsMySqlDB) *SettingServer {
	return &SettingServer{
		db: db,
	}
}

func (s *SettingServer) CreateBatCaveSetting(ctx context.Context, req *settings_pb.CreateBatCaveSettingRequest) (*settings_pb.CreateBatCaveSettingResponse, error) {
	settings := persistence.BatCaveSetting{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.db.CreateBatCaveSetting(settings)

	return &settings_pb.CreateBatCaveSettingResponse{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}, nil
}

func (s *SettingServer) UpdateBatCaveSetting(ctx context.Context, req *settings_pb.UpdateBatCaveSettingRequest) (*settings_pb.UpdateBatCaveSettingResponse, error) {
	setting := persistence.BatCaveSetting{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.db.UpdateBatCaveSetting(setting)

	return &settings_pb.UpdateBatCaveSettingResponse{
		DeviceID:       setting.DeviceID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}

func (s *SettingServer) GetBatCaveSetting(ctx context.Context, req *settings_pb.GetBatCaveSettingRequest) (*settings_pb.GetBatCaveSettingResponse, error) {
	setting, err := s.db.GetBatCaveSetting(req.DeviceID)
	if err != nil {
		return &settings_pb.GetBatCaveSettingResponse{}, err
	}

	return &settings_pb.GetBatCaveSettingResponse{
		DeviceID:       setting.DeviceID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}
