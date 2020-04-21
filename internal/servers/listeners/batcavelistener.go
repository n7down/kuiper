package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/n7down/iota/internal/persistence"
	"github.com/n7down/iota/internal/persistence/mysql"
	"github.com/sirupsen/logrus"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

type BatCaveSettings struct {
	DeviceID       string `json:"deviceID"`
	DeepSleepDelay int32  `json:"deepSleepDelay"`
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
		settings := &BatCaveSettings{}
		err := json.Unmarshal([]byte(msg.Payload()), settings)
		if err != nil {
			logrus.Error(err)
		}

		if err == nil {

			// get the settings
			currentSettings, err := e.db.GetBatCaveSettings(settings.DeviceID)
			if err != nil {
				logrus.Error(err)
			} else {

				// check for the differences in the settings
				if settings.DeepSleepDelay != currentSettings.DeepSleepDelay {

					// update the database with the new settings
					newSettings := persistence.UpdateBatCaveSettings{
						DeepSleepDelay: settings.DeepSleepDelay,
					}

					// marshal new settings and publish it
					jsonData, err := json.Marshal(newSettings)
					if err != nil {
						logrus.Error(err)
					} else {

						// send back to the device the new settings
						deviceTopic := fmt.Sprintf("devices/%s", settings.DeviceID)
						token := client.Publish(deviceTopic, 0, false, jsonData)
						token.WaitTimeout(ONE_MINUTE)

						// update the database with the new settings
						err := e.db.UpdateBatCaveSettings(settings.DeviceID, newSettings)
						if err != nil {
							logrus.Error(err)
						}
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
