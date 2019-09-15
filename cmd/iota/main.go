package main

import (
	"container/list"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/n7down/iota/internal/listeners"
	"github.com/n7down/iota/internal/persistence/influx"
	"github.com/sirupsen/logrus"
)

var (
	Version         string
	Build           string
	showVersion     *bool
	bmp280Listener  *listeners.Listener
	dht22Listener   *listeners.Listener
	voltageListener *listeners.Listener
	timeListener    *listeners.Listener
)

type Iota struct {
	listenerList *list.List
}

func InitIota() *Iota {
	return &Iota{
		listenerList: list.New(),
	}
}

func (i Iota) AddListener(listener *listeners.Listener) {
	i.listenerList.PushBack(listener)
}

func (i Iota) Connection() {
	for l := i.listenerList.Front(); l != nil; l = l.Next() {
		l.Value.(*listeners.Listener).Connect()
	}
}

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

		env := listeners.NewEnv(influxDB)

		dht22MqttURL := os.Getenv("DHT22_MQTT_URL")
		dht22MqttUrl, err := url.Parse(dht22MqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		dht22Listener, err = env.NewDHT22Listener("dht22_listener", dht22MqttUrl)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		bmp280MqttURL := os.Getenv("BMP280_MQTT_URL")
		bmp280MqttUrl, err := url.Parse(bmp280MqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		bmp280Listener, err = env.NewBMP280Listener("bmp280_listener", bmp280MqttUrl)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		voltageMqttURL := os.Getenv("VOLTAGE_MQTT_URL")
		voltageMqttUrl, err := url.Parse(voltageMqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		voltageListener, err = env.NewVoltageListener("voltage_listener", voltageMqttUrl)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		timeMqttURL := os.Getenv("TIME_MQTT_URL")
		timeMqttUrl, err := url.Parse(timeMqttURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}

		timeListener, err = env.NewTimeListener("time", timeMqttUrl)
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

		err := dht22Listener.Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}

		err = bmp280Listener.Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}

		err = voltageListener.Connect()
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
