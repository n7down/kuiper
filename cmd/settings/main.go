package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/n7down/kuiper/internal/logger"
	"github.com/n7down/kuiper/internal/logger/logruslogger"
	"github.com/n7down/kuiper/internal/settings/listeners"
	"github.com/n7down/kuiper/internal/settings/persistence/mysql"
	"google.golang.org/grpc"

	commonServers "github.com/n7down/kuiper/internal/common/servers"
	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
	settings "github.com/n7down/kuiper/internal/settings/servers"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

var (
	Version         string
	Build           string
	showVersion     *bool
	listenersServer *commonServers.ListenersServer
	port            string
	server          *settings.SettingServer
	log             logger.Logger
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		port = os.Getenv("PORT")
		dbConn := os.Getenv("DB_CONN")
		batCaveMQTTURL := os.Getenv("BAT_CAVE_MQTT_URL")

		log = logruslogger.NewLogrusLogger(true)
		persistence, err := mysql.NewMysqlPersistence(dbConn)
		if err != nil {
			log.Fatal(err)
		}

		server = settings.NewSettingServer(persistence)
		listenersServer = commonServers.NewListenersServer()
		settingsListenersEnv := listeners.NewSettingsListenersEnv(persistence, log)

		batCaveListener, err := settingsListenersEnv.NewBatCaveSettingsListener("bat_cave_listener", batCaveMQTTURL)
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(batCaveListener)
	}
}

func main() {
	if *showVersion {
		fmt.Printf("settings server: version %s build %s", Version, Build)
	} else {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}

		listenersServer.Connect()

		log.Infof("Listening on port: %s\n", port)
		grpcServer := grpc.NewServer()
		settings_pb.RegisterSettingsServiceServer(grpcServer, server)
		grpcServer.Serve(lis)
	}
}
