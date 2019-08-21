package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/n7down/iota/internal/sensors"
	"github.com/n7down/iota/internal/stores"
	"github.com/sirupsen/logrus"
)

const (
	clientID = "influx-listener"
)

type InfluxListener struct {
	mqttOptions *mqtt.ClientOptions
}

func NewInfluxListener(mqttUrl *url.URL, store *stores.InfluxStore) (*InfluxListener, error) {
	i := &InfluxListener{}

	topic := mqttUrl.Path[1:len(mqttUrl.Path)]
	if topic == "" {
		topic = "test"
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", mqttUrl.Host))
	opts.SetUsername(mqttUrl.User.Username())
	password, _ := mqttUrl.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		logrus.Infof("Received message: %s\n", msg.Payload())

		// unmashal payload
		sensor := &sensors.AmbientEnvironmentalSensor{}
		err := json.Unmarshal([]byte(msg.Payload()), sensor)
		if err != nil {
			logrus.Error(err.Error())
		}

		if err == nil {
			err = sensor.LogSensor(store)
			logrus.Infof("Logged sensor: %v", sensor)
			if err != nil {
				logrus.Error(err.Error())
			}
		}
	}

	opts.SetDefaultPublishHandler(f)

	var err error = nil
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			//logrus.Fatal(token.Error())
			err = token.Error()
		}
	}

	if err != nil {
		return i, err
	}

	i.mqttOptions = opts

	return i, nil
}

func (i InfluxListener) Connect() error {
	mqttClient := mqtt.NewClient(i.mqttOptions)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
