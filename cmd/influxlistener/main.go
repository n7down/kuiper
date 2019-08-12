package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/sirupsen/logrus"
)

type Sensor struct {
	client          client.Client
	sensorDataPoint SensorDataPoint
}

type SensorDataPoint struct {
	Location string `json:"location"`

	// label ais the label on the device
	Label string `json:"label"`

	// version is the firmware version on the device
	Version string `json:"version"`

	// device is determined by the hardware that makes up the device
	Device                   string `json:"device"`
	Voltage                  string `json:"voltage"`
	DHT22Temperature         string `json:"dht22temp"`
	DHT22Humidity            string `json:"dht22humidity"`
	BMP280Temperature        string `json:"bmp280temperture"`
	BMP280BarometricPressure string `json:"bmp280pressure"`
}

func (s SensorDataPoint) LogSensor() error {
	influxClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "dbuser",
		Password: "password",
	})

	if err != nil {
		return err
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "sensors",
		Precision: "s",
	})

	// indexed
	tags := map[string]string{
		"location": s.Location,
		"version":  s.Version,
		"device":   s.Device,
	}

	// not indexed
	fields := map[string]interface{}{
		"dht22_temp":           s.DHT22Temperature,
		"dht22_humidity":       s.DHT22Humidity,
		"bmp280_temp":          s.BMP280Temperature,
		"bmp280_bero_pressure": s.BMP280BarometricPressure,
	}

	point, err := client.NewPoint(
		"sensor",
		tags,
		fields,
		time.Now(),
	)

	bp.AddPoint(point)

	err = influxClient.Write(bp)
	if err != nil {
		return err
	}

	return nil
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logrus.Infof("Received message: %s\n", msg.Payload())

	// unmashal payload
	sensorDataPoint := &SensorDataPoint{}
	err := json.Unmarshal([]byte(msg.Payload()), sensorDataPoint)
	if err != nil {
		logrus.Error(err.Error())
	}

	err = sensorDataPoint.LogSensor()
	if err != nil {
		logrus.Error(err.Error())
	}
}

func init() {

}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	mqttUri, err := url.Parse(os.Getenv("MQTT_URL"))
	if err != nil {
		logrus.Fatal(err.Error())
	}

	topic := mqttUri.Path[1:len(mqttUri.Path)]
	if topic == "" {
		topic = "test"
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", mqttUri.Host))
	opts.SetUsername(mqttUri.User.Username())
	password, _ := mqttUri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID("influx-listener")
	opts.SetDefaultPublishHandler(f)

	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			logrus.Fatal(token.Error())
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logrus.Fatal(token.Error())
	} else {
		logrus.Info("Connected to server\n")
	}
	<-c
}
