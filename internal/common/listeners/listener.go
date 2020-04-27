package listeners

import (
	"errors"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Listener struct {
	MqttOptions     *mqtt.ClientOptions
	ListenerName    string
	ListenerMQTTUrl *url.URL
}

func (l Listener) Connect() error {
	mqttClient := mqtt.NewClient(l.MqttOptions)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return errors.New(fmt.Sprintf("Error with %s: %s", l.ListenerName, token.Error()))
	}
	fmt.Println(fmt.Sprintf("%s on %s is connected", l.ListenerName, l.ListenerMQTTUrl.String()))
	return nil
}
