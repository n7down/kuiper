package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	listeners "github.com/n7down/kuiper/internal/common/listeners"
)

type Time struct {
	DeviceName string `json:"deviceName"`
}

func connectToMQTT(clientId string, uri *url.URL) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		logrus.Error(err.Error())
	}
	return client
}

func (e SensorsListenersEnv) NewTimeListener(listenerName string, timeMqttURL string) (*listeners.Listener, error) {
	i := &listeners.Listener{}

	mqttUrl, err := url.Parse(timeMqttURL)
	if err != nil {
		return i, err
	}

	subscribeTopic := mqttUrl.Path[1:len(mqttUrl.Path)]
	if subscribeTopic == "" {
		subscribeTopic = "test"
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", mqttUrl.Host))
	opts.SetUsername(mqttUrl.User.Username())
	password, _ := mqttUrl.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(listenerName)

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		logrus.Infof("Received message: %s\n", msg.Payload())

		// unmashal payload
		t := Time{}
		err := json.Unmarshal([]byte(msg.Payload()), &t)
		if err != nil {
			logrus.Error(err.Error())
		}

		if err == nil {
			currentTime := time.Now().UTC()
			clientID := fmt.Sprintf("%s-%s", listenerName, t.DeviceName)
			client := connectToMQTT(clientID, mqttUrl)
			publishTopicName := fmt.Sprintf("%s/%s", subscribeTopic, t.DeviceName)
			client.Publish(publishTopicName, 0, false, currentTime.String())
			logrus.Infof("Sending time: %v %v", publishTopicName, currentTime.String())
			client.Disconnect(1000)
		}
	}

	opts.SetDefaultPublishHandler(f)

	err = nil
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(subscribeTopic, 0, f); token.Wait() && token.Error() != nil {
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
