package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/sirupsen/logrus"
)

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

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		logrus.Fatal(err.Error())
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func listen(uri *url.URL, topic string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		logrus.Infof("* [%s] %s\n", msg.Topic(), string(msg.Payload()))

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
	})
}

func main() {
	uri, err := url.Parse(os.Getenv("MQTT_URL"))
	if err != nil {
		logrus.Fatal(err.Error())
	}
	topic := uri.Path[1:len(uri.Path)]

	fmt.Println(fmt.Sprintf("URL: %v", uri))

	go listen(uri, topic)

	client := connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)

	d := SensorDataPoint{
		Location:                 "master bathroom",
		Label:                    "1",
		Version:                  "v0.1",
		Device:                   "bc",
		DHT22Humidity:            "60.0",
		DHT22Temperature:         "23.4",
		BMP280Temperature:        "25.0",
		BMP280BarometricPressure: "26.0",
	}

	j, err := json.Marshal(d)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	for _ = range timer.C {
		client.Publish(topic, 0, false, string(j))
	}
}
