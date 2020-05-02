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

func (l *Listener) Connect() error {
	mqttClient := mqtt.NewClient(l.MqttOptions)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return errors.New(fmt.Sprintf("Error with %s: %s", l.ListenerName, token.Error()))
	}
	fmt.Println(fmt.Sprintf("%s on %s is connected", l.ListenerName, l.ListenerMQTTUrl.String()))
	return nil
}

func NewListener(listenerName string, urlString string, callback func(mqtt.Client, mqtt.Message)) (*Listener, error) {
	l := &Listener{}

	mqttUrl, err := url.Parse(urlString)
	if err != nil {
		return l, err
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

	var f mqtt.MessageHandler = callback

	opts.SetDefaultPublishHandler(f)

	err = nil
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			err = token.Error()
		}
	}

	if err != nil {
		return l, err
	}

	l.MqttOptions = opts
	l.ListenerName = listenerName
	l.ListenerMQTTUrl = mqttUrl

	return l, nil
}
