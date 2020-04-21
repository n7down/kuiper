package main

import (
	"fmt"
	"net"
	"os"

	"github.com/n7down/iota/internal/persistence/mysql"
	"github.com/n7down/iota/internal/servers"
	"google.golang.org/grpc"

	settings_pb "github.com/n7down/iota/internal/pb/settings"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	port := os.Getenv("PORT")
	dbConn := os.Getenv("DB_CONN")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: updating settings on device
	// 1. device wakes up and send message to settings service with all set settings
	// - sends message to /device/settings
	// - sends deviceID and all settings
	// 2. settings service checks for differences in database
	// 3. if there is a difference in the settings for the device - it sends the difference to the device

	settingsDB, err := mysql.NewSettingsMySqlDB(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	settingsServer := servers.NewSettingsServer(settingsDB)

	log.Infof("Listening on port: %s\n", port)
	grpcServer := grpc.NewServer()
	settings_pb.RegisterSettingsServiceServer(grpcServer, settingsServer)
	grpcServer.Serve(lis)
}
