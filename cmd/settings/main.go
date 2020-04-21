package main

import (
	"fmt"
	"net"
	"os"

	"github.com/n7down/iota/internal/persistence/mysql"
	"github.com/n7down/iota/internal/servers"
	"github.com/n7down/iota/internal/servers/listeners"
	"google.golang.org/grpc"

	settings_pb "github.com/n7down/iota/internal/pb/settings"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	port := os.Getenv("PORT")
	dbConn := os.Getenv("DB_CONN")
	batCaveMQTTURL := os.Getenv("BAT_CAVE_MQTT_URL")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	settingsDB, err := mysql.NewSettingsMySqlDB(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	settingsServer := servers.NewSettingsServer(settingsDB)

	env := listeners.NewEnv(settingsDB)
	listenersServer := servers.NewListenersServer()
	batCaveListener, err := env.NewBatCaveSettingsListener("bat_cave_listener", batCaveMQTTURL)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	listenersServer.AddListener(batCaveListener)
	listenersServer.Connect()

	log.Infof("Listening on port: %s\n", port)
	grpcServer := grpc.NewServer()
	settings_pb.RegisterSettingsServiceServer(grpcServer, settingsServer)
	grpcServer.Serve(lis)
}
