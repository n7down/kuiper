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
	Version                   string
	Build                     string
	showVersion               *bool
	indoorHumidityListener    *listeners.Listener
	indoorTemperatureListener *listeners.Listener
	indoorPressureListener    *listeners.Listener
	indoorVoltageListener     *listeners.Listener
	timeListener              *listeners.Listener
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

		store, err := stores.NewInfluxStore(influxUrl)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		indoorHumidityMqttURL := os.Getenv("HUMIDITY_MQTT_URL")
		indoorHumidityMqttUrl, err := url.Parse(indoorHumidityMqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		indoorHumidityListener, err = listeners.NewHumidityListener("humidity", indoorHumidityMqttUrl, store)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		indoorTemperatureMqttURL := os.Getenv("TEMPERATURE_MQTT_URL")
		indoorTemperatureMqttUrl, err := url.Parse(indoorTemperatureMqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		indoorTemperatureListener, err = listeners.NewTemperatureListener("temp", indoorTemperatureMqttUrl, store)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		indoorPressureMqttURL := os.Getenv("PRESSURE_MQTT_URL")
		indoorPressureMqttUrl, err := url.Parse(indoorPressureMqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		indoorPressureListener, err = listeners.NewPressureListener("pressure", indoorPressureMqttUrl, store)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		indoorVoltageMqttURL := os.Getenv("VOLTAGE_MQTT_URL")
		indoorVoltageMqttUrl, err := url.Parse(indoorVoltageMqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		indoorVoltageListener, err = listeners.NewVoltageListener("voltage", indoorVoltageMqttUrl, store)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		timeMqttURL := os.Getenv("TIME_MQTT_URL")
		timeMqttUrl, err := url.Parse(timeMqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		timeListener, err = listeners.NewTimeListener("time", timeMqttUrl, store)
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

		err := indoorHumidityListener.Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}

		err = indoorTemperatureListener.Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}

		err = indoorPressureListener.Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}

		err = indoorVoltageListener.Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}

		err = timeListener.Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}

		<-c
	}
}
