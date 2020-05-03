package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/n7down/kuiper/internal/common/listeners"
	"github.com/n7down/kuiper/internal/settings/listeners/request"
	"github.com/n7down/kuiper/internal/settings/listeners/response"
	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/n7down/kuiper/internal/settings/persistence/mysql"
	"google.golang.org/grpc"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	commonServers "github.com/n7down/kuiper/internal/common/servers"
	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
	settings "github.com/n7down/kuiper/internal/settings/servers"
	log "github.com/sirupsen/logrus"
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
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		port = os.Getenv("PORT")
		dbConn := os.Getenv("DB_CONN")
		batCaveMQTTURL := os.Getenv("BAT_CAVE_MQTT_URL")

		settingsDB, err := mysql.NewSettingsMySqlDB(dbConn)
		if err != nil {
			log.Fatal(err)
		}
		server = settings.NewSettingServer(settingsDB)

		listenersServer = commonServers.NewListenersServer()

		batCaveCallback := func(client mqtt.Client, msg mqtt.Message) {
			log.Infof("Received message: %s\n", msg.Payload())

			// unmashal payload
			var (
				req request.BatCaveSettingRequest
				res response.BatCaveSettingResponse
			)

			err := json.Unmarshal([]byte(msg.Payload()), &req)
			if err != nil {
				log.Error(err)
				return
			}

			// get the settings
			recordNotFound, settingInPersistence := settingsDB.GetBatCaveSetting(req.DeviceID)
			if recordNotFound {
				newSetting := persistence.BatCaveSetting{
					DeviceID:       req.DeviceID,
					DeepSleepDelay: req.DeepSleepDelay,
				}

				// create the new setting
				settingsDB.CreateBatCaveSetting(newSetting)

				// send back default values
				res = response.GetBatCaveSettingDefault()

			} else {

				// check for the differences in the settings
				var isEqual bool
				isEqual, res = req.IsEqual(settingInPersistence)
				log.Infof("Settings are equal: %t - %v %v", isEqual, settingInPersistence, res)
				if isEqual {

					// settings are the same on the device and in persistence - return
					return
				}
			}

			json, err := json.Marshal(res)
			if err != nil {
				log.Error(err)
				return
			}

			// send back to the device the new settings
			deviceTopic := fmt.Sprintf("devices/%s", req.DeviceID)
			log.Infof("Sending message %s to %s", json, deviceTopic)
			token := client.Publish(deviceTopic, 0, false, json)
			token.WaitTimeout(ONE_MINUTE)

		}

		batCaveListener, err := listeners.NewListener("bat_cave_listener", batCaveMQTTURL, batCaveCallback)
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
		log.SetReportCaller(true)

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
