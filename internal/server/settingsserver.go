package server

import (
	"context"

	_ "github.com/go-sql-driver/mysql"

	settings_pb "github.com/n7down/iota/internal/pb/settings"
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

func (s *SettingsServer) SetSettings(ctx context.Context, req *settings_pb.SetSettingsRequest) (*settings_pb.SetSettingsResponse, error) {
	return &settings_pb.SetSettingsResponse{}, nil
}

func (s *SettingsServer) GetSettings(ctx context.Context, req *settings_pb.GetSettingsRequest) (*settings_pb.GetSettingsResponse, error) {
	return &settings_pb.GetSettingsResponse{}, nil
}
