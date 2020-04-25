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
	"github.com/n7down/kuiper/internal/sensors/servers"

	log "github.com/sirupsen/logrus"
)

var (
	Version       string
	Build         string
	showVersion   *bool
	sensorsServer *servers.SensorsServer
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

		sensorsServer = servers.NewSensorsServer()
		env := listeners.NewEnv(influxDB)

		dht22Listener, err := env.NewDHT22Listener("dht22_listener", os.Getenv("DHT22_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		sensorsServer.AddListener(dht22Listener)

		bmp280Listener, err := env.NewBMP280Listener("bmp280_listener", os.Getenv("BMP280_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		sensorsServer.AddListener(bmp280Listener)

		voltageListener, err := env.NewVoltageListener("voltage_listener", os.Getenv("VOLTAGE_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		sensorsServer.AddListener(voltageListener)

		timeListener, err := env.NewTimeListener("time_listener", os.Getenv("TIME_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		sensorsServer.AddListener(timeListener)

		statsListener, err := env.NewStatsListener("stats_listener", os.Getenv("STATS_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
		sensorsServer.AddListener(statsListener)
	}
}

func main() {
	if *showVersion {
		fmt.Printf("sensors server: version %s build %s", Version, Build)
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		sensorsServer.Connect()

		<-c
	}
}
