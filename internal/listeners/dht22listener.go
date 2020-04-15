package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/n7down/iota/internal/sensors"
	"github.com/sirupsen/logrus"
)

func (e Env) NewDHT22Listener(listenerName string, dht22MqttURL string) (*Listener, error) {
	i := &Listener{}

	mqttUrl, err := url.Parse(dht22MqttURL)
	if err != nil {
		return i, err
	}

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
		sensor := &sensors.DHT22Sensor{}
		err := json.Unmarshal([]byte(msg.Payload()), sensor)
		if err != nil {
			logrus.Error(err.Error())
		}

		if err == nil {
			err = e.influxDB.LogDHT22(listenerName, sensor)
			logrus.Infof("Logged sensor: %v", sensor)
			if err != nil {
				logrus.Error(err.Error())
			}
		}
	}

	opts.SetDefaultPublishHandler(f)

	err = nil
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			err = token.Error()
		}
	}

	if err != nil {
		return i, err
	}

	i.mqttOptions = opts
	i.listenerName = listenerName
	i.listenerMQTTUrl = mqttUrl

	return i, nil
}
