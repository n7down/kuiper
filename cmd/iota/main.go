package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/n7down/iota/internal/listeners"
	"github.com/n7down/iota/internal/stores"
	"github.com/sirupsen/logrus"
)

var (
	Version        string
	Build          string
	showVersion    *bool
	influxListener *listeners.InfluxListener
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		logrus.SetReportCaller(true)

		mqttURL := os.Getenv("MQTT_URL")
		mqttUrl, err := url.Parse(mqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		influxURL := os.Getenv("INFLUX_URL")
		influxUrl, err := url.Parse(influxURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		store, err := stores.NewInfluxStore(influxUrl)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		influxListener, err = listeners.NewInfluxListener(mqttUrl, store)
		if err != nil {
			logrus.Fatal(err.Error())
		}
	}
}

func main() {
	if *showVersion {
		fmt.Printf("iota version %s build %s", Version, Build)
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		err := influxListener.Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}
		logrus.Info("Connected to server\n")

		<-c
	}
}
