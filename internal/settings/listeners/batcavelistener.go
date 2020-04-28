package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/n7down/kuiper/internal/settings/listeners/request"
	"github.com/n7down/kuiper/internal/settings/persistence"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	listeners "github.com/n7down/kuiper/internal/common/listeners"
	log "github.com/sirupsen/logrus"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

func (e SettingsListenerEnv) NewBatCaveSettingsListener(listenerName string, mqttURL string) (*listeners.Listener, error) {
	i := &listeners.Listener{}

	u, err := url.Parse(mqttURL)
	if err != nil {
		return i, err
	}

	topic := u.Path[1:len(u.Path)]
	if topic == "" {
		topic = "test"
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", u.Host))
	opts.SetUsername(u.User.Username())
	password, _ := u.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(listenerName)

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		log.Infof("Received message: %s\n", msg.Payload())

		// unmashal payload
		req := &request.BatCaveSettingRequest{}
		err := json.Unmarshal([]byte(msg.Payload()), req)
		if err != nil {
			log.Error(err)
		} else {

			// get the settings
			recordNotFound, settingInPersistence := e.db.GetBatCaveSetting(req.DeviceID)
			if recordNotFound {
				newSetting := persistence.BatCaveSetting{
					DeviceID:       req.DeviceID,
					DeepSleepDelay: req.DeepSleepDelay,
				}

				// create the new setting
				e.db.CreateBatCaveSetting(newSetting)
			} else {

				// check for the differences in the settings
				isEqual, commands := req.IsEqual(settingInPersistence)
				if !isEqual {
					for command := range commands {

						// send back to the device the new settings
						deviceTopic := fmt.Sprintf("devices/%s", req.DeviceID)
						log.Infof("Sending message %s to %s", command, deviceTopic)
						token := client.Publish(deviceTopic, 0, false, command)
						token.WaitTimeout(ONE_MINUTE)
					}
				}
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
	i.ListenerMQTTUrl = u

	return i, nil
}
