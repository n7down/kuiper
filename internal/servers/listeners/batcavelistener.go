package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/n7down/iota/internal/persistence/mysql"
	"github.com/sirupsen/logrus"
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
		// sensor := &sensors.BMP280Sensor{}
		settings := &BatCaveSettings{}
		err := json.Unmarshal([]byte(msg.Payload()), settings)
		if err != nil {
			logrus.Error(err)
		}

		if err == nil {
			// err = e.influxDB.LogBMP280(listenerName, sensor)
			// logrus.Infof("Logged sensor: %v", sensor)
			// if err != nil {
			// 	logrus.Error(err.Error())
			// }

			// TODO: updating settings on device
			// 1. device wakes up and send message to settings service with all set settings
			// - sends message to /device/settings
			// - sends deviceID and all settings
			// 2. settings service checks for differences in database
			// 3. if there is a difference in the settings for the device - it sends the difference to the device

			// get the settings
			_, err := e.db.GetSettings(settings.DeviceID)
			if err != nil {
				logrus.Error(err)
			} else {
				// TODO: check for the differences in the settings
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
