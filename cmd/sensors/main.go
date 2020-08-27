package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/n7down/kuiper/internal/logger"
	"github.com/n7down/kuiper/internal/logger/logruslogger"
	"github.com/n7down/kuiper/internal/sensors/persistence/influxpersistence"
	"github.com/n7down/kuiper/internal/sensors/pubsub/mosquitto"
	"google.golang.org/grpc"

	sensors_pb "github.com/n7down/kuiper/internal/pb/sensors"
	sensors "github.com/n7down/kuiper/internal/sensors/servers"
)

var (
	Version     string
	Build       string
	showVersion *bool
	port        string
	log         logger.Logger
	server      *sensors.SensorsServer
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		port = os.Getenv("PORT")
		ctx := context.Background()
		log = logruslogger.NewLogrusLogger(true)

		influxURL := os.Getenv("INFLUX_URL")
		influxUrl, err := url.Parse(influxURL)
		if err != nil {
			log.Fatal(err.Error())
		}

		persistence, err := influxpersistence.NewInfluxPersistence(influxUrl, log)
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

		server = sensors.NewSensorsServer(persistence)
	}
}

func main() {
	if *showVersion {
		fmt.Printf("sensors server: version %s build %s", Version, Build)
	} else {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Listening on port: %s\n", port)
		grpcServer := grpc.NewServer()
		sensors_pb.RegisterSensorsServiceServer(grpcServer, server)
		grpcServer.Serve(lis)
	}
}
