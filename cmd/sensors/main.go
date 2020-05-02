package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/n7down/kuiper/internal/common/listeners"
	"github.com/n7down/kuiper/internal/sensors/persistence/influxdb"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	commonServers "github.com/n7down/kuiper/internal/common/servers"
	sensors "github.com/n7down/kuiper/internal/sensors/devicesensors"
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

		listenersServer = commonServers.NewListenersServer()

		dht22Callback := func(client mqtt.Client, msg mqtt.Message) {
			log.Infof("Received message: %s\n", msg.Payload())
			sensor := &sensors.DHT22Sensor{}
			err := json.Unmarshal([]byte(msg.Payload()), sensor)
			if err != nil {
				log.Error(err.Error())
			}

			if err == nil {
				err = influxDB.LogDHT22(sensor)
				log.Infof("Logged sensor: %v", sensor)
				if err != nil {
					log.Error(err.Error())
				}
			}
		}

		dht22Listener, err := listeners.NewListener("dht22_listener", os.Getenv("DHT22_MQTT_URL"), dht22Callback)
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(dht22Listener)

		voltageCallback := func(client mqtt.Client, msg mqtt.Message) {
			log.Infof("Received message: %s\n", msg.Payload())
			sensors := &sensors.VoltageSensor{}
			err := json.Unmarshal([]byte(msg.Payload()), sensors)
			if err != nil {
				log.Error(err.Error())
			}

			if err == nil {
				err = influxDB.LogVoltage(sensors)
				log.Infof("Logged sensor: %v", sensors)
				if err != nil {
					log.Error(err.Error())
				}
			}
		}

		voltageListener, err := listeners.NewListener("voltage_listener", os.Getenv("VOLTAGE_MQTT_URL"), voltageCallback)
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(voltageListener)

		statsCallback := func(client mqtt.Client, msg mqtt.Message) {
			log.Infof("Received message: %s\n", msg.Payload())
			sensors := &sensors.StatsSensor{}
			err := json.Unmarshal([]byte(msg.Payload()), sensors)
			if err != nil {
				log.Error(err.Error())
			}

			if err == nil {
				err = influxDB.LogStats(sensors)
				log.Infof("Logged sensor: %v", sensors)
				if err != nil {
					log.Error(err.Error())
				}
			}
		}

		statsListener, err := listeners.NewListener("stats_listener", os.Getenv("STATS_MQTT_URL"), statsCallback)
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(statsListener)

		hdc1080Callback := func(client mqtt.Client, msg mqtt.Message) {
			log.Infof("Received message: %s\n", msg.Payload())
			sensor := &sensors.HDC1080Sensor{}
			err := json.Unmarshal([]byte(msg.Payload()), sensor)
			if err != nil {
				log.Error(err.Error())
			}

			if err == nil {
				err = influxDB.LogHDC1080(sensor)
				log.Infof("Logged sensor: %v", sensor)
				if err != nil {
					log.Error(err.Error())
				}
			}
		}

		hdc1080Listener, err := listeners.NewListener("hdc1080_listener", os.Getenv("HDC1080_MQTT_URL"), hdc1080Callback)
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
