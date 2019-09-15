package listeners

import (
	"errors"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Listener struct {
	mqttOptions     *mqtt.ClientOptions
	listenerName    string
	listenerMQTTUrl *url.URL
}

func (l Listener) Connect() error {
	mqttClient := mqtt.NewClient(l.mqttOptions)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return errors.New(fmt.Sprintf("Error with %s: %s", l.listenerName, token.Error()))
	}
	fmt.Println(fmt.Sprintf("%s on %s is connected", l.listenerName, l.listenerMQTTUrl.String()))
	return nil
}
