package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	listeners "github.com/n7down/kuiper/internal/common/listeners"
	sensors "github.com/n7down/kuiper/internal/sensors/devicesensors"
)

func (e SensorsListenersEnv) NewDHT22Listener(listenerName string, dht22MqttURL string) (*listeners.Listener, error) {
	i := &listeners.Listener{}

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
			err = e.influxDB.LogDHT22(sensor)
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

	i.MqttOptions = opts
	i.ListenerName = listenerName
	i.ListenerMQTTUrl = mqttUrl

	return i, nil
}
