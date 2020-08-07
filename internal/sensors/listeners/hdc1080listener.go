package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	listeners "github.com/n7down/kuiper/internal/common/listeners"
	sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"
)

func (e SensorsListenersEnv) NewHDC1080Listener(listenerName string, urlString string) (*listeners.Listener, error) {
	i := &listeners.Listener{}

	mqttUrl, err := url.Parse(urlString)
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
		e.logger.Infof("Received message: %s\n", msg.Payload())
		sensor := &sensors.HDC1080Sensor{}
		err := json.Unmarshal([]byte(msg.Payload()), sensor)
		if err != nil {
			e.logger.Error(err.Error())
		}

		if err == nil {
			err = e.persistence.LogHDC1080(sensor)
			e.logger.Infof("Logged sensor: %v", sensor)
			if err != nil {
				e.logger.Error(err.Error())
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
