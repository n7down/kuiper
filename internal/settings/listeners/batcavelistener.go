package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/n7down/iota/internal/settings/persistence"
	"github.com/n7down/iota/internal/settings/persistence/mysql"
	"github.com/sirupsen/logrus"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

type BatCaveSettingsRequest struct {
	DeviceID       string `json:"deviceID"`
	DeepSleepDelay int32  `json:"deepSleepDelay"`
}

func (s *BatCaveSettingsRequest) IsEqual(settings persistence.GetBatCaveSettings) bool {
	if s.DeepSleepDelay == settings.DeepSleepDelay {
		return true
	}
	return false
}

type BatCaveSettingsResponse struct {
	DeepSleepDelay int32 `json:"deepSleepDelay"`
}

type Env struct {
	db *mysql.SettingsMySqlDB
}

func NewEnv(db *mysql.SettingsMySqlDB) *Env {
	return &Env{
		db: db,
	}
}

func (e Env) NewBatCaveSettingsListener(listenerName string, mqttURL string) (*Listener, error) {
	i := &Listener{}

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
		logrus.Infof("Received message: %s\n", msg.Payload())

		// unmashal payload
		req := &BatCaveSettingsRequest{}
		err := json.Unmarshal([]byte(msg.Payload()), req)
		if err != nil {
			logrus.Error(err)
		} else {

			// get the settings
			settingsInPersistence, err := e.db.GetBatCaveSettings(req.DeviceID)
			if err != nil {
				logrus.Error(err)
			} else {

				// check for the differences in the settings
				if !req.IsEqual(settingsInPersistence) {

					settingsToSendToDevice := BatCaveSettingsResponse{
						DeepSleepDelay: settingsInPersistence.DeepSleepDelay,
					}

					// marshal data to send back to the device
					jsonData, err := json.Marshal(settingsToSendToDevice)
					if err != nil {
						logrus.Error(err)
					} else {

						// send back to the device the new settings
						deviceTopic := fmt.Sprintf("devices/%s", req.DeviceID)
						token := client.Publish(deviceTopic, 0, false, jsonData)
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

	i.mqttOptions = opts
	i.listenerName = listenerName
	i.listenerMQTTUrl = u

	return i, nil
}
