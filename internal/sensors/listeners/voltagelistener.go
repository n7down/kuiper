package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	listeners "github.com/n7down/kuiper/internal/common/listeners"
	sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"
)

func (e SensorsListenersEnv) NewVoltageListener(listenerName string, statsMqttURL string) (*listeners.Listener, error) {
	i := &listeners.Listener{}

	mqttUrl, err := url.Parse(statsMqttURL)
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

		// unmashal payload
		sensors := &sensors.VoltageSensor{}
		err := json.Unmarshal([]byte(msg.Payload()), sensors)
		if err != nil {
			e.logger.Error(err.Error())
		}

		if err == nil {
			err = e.persistence.CreateVoltage(sensors)
			e.logger.Infof("Logged sensor: %v", sensors)
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
