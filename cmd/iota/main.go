package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/n7down/iota/internal/listeners"
	"github.com/n7down/iota/internal/persistence/influx"
	"github.com/n7down/iota/internal/server"
	"github.com/sirupsen/logrus"
)

var (
	Version     string
	Build       string
	showVersion *bool
	iotaServer  *server.IotaServer
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		logrus.SetReportCaller(true)

		influxURL := os.Getenv("INFLUX_URL")
		influxUrl, err := url.Parse(influxURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		influxDB, err := influx.NewInflux(influxUrl)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		iotaServer = server.NewIotaServer()
		env := listeners.NewEnv(influxDB)

		dht22Listener, err := env.NewDHT22Listener("dht22_listener", os.Getenv("DHT22_MQTT_URL"))
		if err != nil {
			logrus.Fatal(err.Error())
		}
		iotaServer.AddListener(dht22Listener)

		bmp280Listener, err := env.NewBMP280Listener("bmp280_listener", os.Getenv("BMP280_MQTT_URL"))
		if err != nil {
			logrus.Fatal(err.Error())
		}
		iotaServer.AddListener(bmp280Listener)

		voltageListener, err := env.NewVoltageListener("voltage_listener", os.Getenv("VOLTAGE_MQTT_URL"))
		if err != nil {
			logrus.Fatal(err.Error())
		}
		iotaServer.AddListener(voltageListener)

		timeListener, err := env.NewTimeListener("time_listener", os.Getenv("TIME_MQTT_URL"))
		if err != nil {
			logrus.Fatal(err.Error())
		}
		iotaServer.AddListener(timeListener)
	}
}

func main() {
	if *showVersion {
		fmt.Printf("iota version %s build %s", Version, Build)
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		iotaServer.Connect()

		<-c
	}
}
