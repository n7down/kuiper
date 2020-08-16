package servers

import (
	"context"
	"errors"

	"github.com/n7down/kuiper/internal/settings/persistence"

	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
)

type SettingServer struct {
	persistence persistence.Persistence
}

func NewSettingServer(persistence persistence.Persistence) *SettingServer {
	return &SettingServer{
		persistence: persistence,
	}
}

func (s *SettingServer) CreateBatCaveSetting(ctx context.Context, req *settings_pb.CreateBatCaveSettingRequest) (*settings_pb.CreateBatCaveSettingResponse, error) {
	settings := persistence.BatCaveSetting{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.persistence.CreateBatCaveSetting(settings)

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

	s.persistence.UpdateBatCaveSetting(setting)

	return &settings_pb.UpdateBatCaveSettingResponse{
		DeviceID:       setting.DeviceID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}

func (s *SettingServer) GetBatCaveSetting(ctx context.Context, req *settings_pb.GetBatCaveSettingRequest) (*settings_pb.GetBatCaveSettingResponse, error) {
	recordNotFound, setting := s.persistence.GetBatCaveSetting(req.DeviceID)
	if recordNotFound {
		return &settings_pb.GetBatCaveSettingResponse{}, errors.New("record not found")
	}

	return &settings_pb.GetBatCaveSettingResponse{
		DeviceID:       setting.DeviceID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}
