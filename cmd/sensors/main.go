package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/n7down/kuiper/internal/sensors/listeners"
	"github.com/n7down/kuiper/internal/sensors/persistence/influxdb"

	commonServers "github.com/n7down/kuiper/internal/common/servers"
	log "github.com/sirupsen/logrus"
)

var (
	Version         string
	Build           string
	showVersion     *bool
	listenersServer *commonServers.ListenersServer
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		log.SetReportCaller(true)

		influxURL := os.Getenv("INFLUX_URL")
		influxUrl, err := url.Parse(influxURL)
		if err != nil {
			log.Fatal(err.Error())
		}

		influxDB, err := influxdb.NewInfluxDB(influxUrl)
		if err != nil {
			log.Fatal(err.Error())
		}

		sensorsListenersEnv := listeners.NewSensorsListenersEnv(influxDB)
		listenersServer = commonServers.NewListenersServer()

		dht22Listener, err := sensorsListenersEnv.NewDHT22Listener("dht22_listener", os.Getenv("DHT22_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(dht22Listener)

		voltageListener, err := sensorsListenersEnv.NewVoltageListener("voltage_listener", os.Getenv("VOLTAGE_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(voltageListener)

		statsListener, err := sensorsListenersEnv.NewStatsListener("stats_listener", os.Getenv("STATS_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(statsListener)

		hdc1080Listener, err := sensorsListenersEnv.NewHDC1080Listener("hdc1080_listener", os.Getenv("HDC1080_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(hdc1080Listener)
	}
}

func main() {
	if *showVersion {
		fmt.Printf("sensors server: version %s build %s", Version, Build)
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		listenersServer.Connect()

		<-c
	}
}
