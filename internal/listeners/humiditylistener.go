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

type HumidityListener struct {
	mqttOptions *mqtt.ClientOptions
}

func NewHumidityListener(listenerName string, mqttUrl *url.URL, store *stores.InfluxStore) (*HumidityListener, error) {
	i := &HumidityListener{}

	topic := mqttUrl.Path[1:len(mqttUrl.Path)]
	if topic == "" {
		topic = "test"
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", mqttUrl.Host))
	opts.SetUsername(mqttUrl.User.Username())
	password, _ := mqttUrl.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(listenerName)

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		logrus.Infof("Received message: %s\n", msg.Payload())

		// unmashal payload
		sensors := &sensors.HumiditySensors{}
		err := json.Unmarshal([]byte(msg.Payload()), sensors)
		if err != nil {
			logrus.Error(err.Error())
		}

		if err == nil {
			err = sensors.LogSensors(store, listenerName)
			logrus.Infof("Logged sensor: %v", sensors)
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

func (l HumidityListener) Connect() error {
	mqttClient := mqtt.NewClient(l.mqttOptions)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
