package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/n7down/kuiper/internal/settings/listeners"
	"github.com/n7down/kuiper/internal/settings/persistence/mysql"
	"github.com/n7down/kuiper/internal/settings/servers"
	"google.golang.org/grpc"

	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
	log "github.com/sirupsen/logrus"
)

var (
	Version     string
	Build       string
	showVersion *bool
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
}

func main() {
	if *showVersion {
		fmt.Printf("settings server: version %s build %s", Version, Build)
	} else {
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
			log.Fatal(err)
		}
		listenersServer.AddListener(batCaveListener)
		listenersServer.Connect()

		log.Infof("Listening on port: %s\n", port)
		grpcServer := grpc.NewServer()
		settings_pb.RegisterSettingsServiceServer(grpcServer, settingsServer)
		grpcServer.Serve(lis)
	}
}
