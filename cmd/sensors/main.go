package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/n7down/kuiper/internal/logger/logruslogger"
	"github.com/n7down/kuiper/internal/sensors/persistence/influxpersistence"
	"github.com/n7down/kuiper/internal/sensors/pubsub/mosquitto"
)

var (
	Version     string
	Build       string
	showVersion *bool
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		ctx := context.Background()
		log := logruslogger.NewLogrusLogger(true)

		influxURL := os.Getenv("INFLUX_URL")
		influxUrl, err := url.Parse(influxURL)
		if err != nil {
			log.Fatal(err.Error())
		}

		persistence, err := influxpersistence.NewInfluxPersistence(influxUrl)
		if err != nil {
			log.Fatal(err.Error())
		}

		pubSub := mosquitto.NewMosquittoPubSub(persistence, log)

		err = pubSub.NewDHT22Listener(ctx, "dht22_listener", os.Getenv("DHT22_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}

		err = pubSub.NewVoltageListener(ctx, "voltage_listener", os.Getenv("VOLTAGE_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}

		err = pubSub.NewStatsListener(ctx, "stats_listener", os.Getenv("STATS_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}

		err = pubSub.NewHDC1080Listener(ctx, "hdc1080_listener", os.Getenv("HDC1080_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	if *showVersion {
		fmt.Printf("sensors server: version %s build %s", Version, Build)
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	}
}
