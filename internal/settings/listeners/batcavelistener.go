package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/n7down/kuiper/internal/settings/persistence/mysql"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

type BatCaveSettingRequest struct {
	DeviceID       string `json:"m"`
	DeepSleepDelay int32  `json:"s"`
}

func (s *BatCaveSettingRequest) IsEqual(settings persistence.BatCaveSetting) bool {
	if s.DeepSleepDelay == settings.DeepSleepDelay {
		return true
	}
	return false
}

type BatCaveSettingResponse struct {
	DeepSleepDelay int32 `json:"s"`
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
		log.Infof("Received message: %s\n", msg.Payload())

		// unmashal payload
		req := &BatCaveSettingRequest{}
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
				isEqual := req.IsEqual(settingInPersistence)
				if !isEqual {
					settingsToSendToDevice := BatCaveSettingResponse{
						DeepSleepDelay: settingInPersistence.DeepSleepDelay,
					}

					// marshal data to send back to the device
					jsonData, err := json.Marshal(settingsToSendToDevice)
					if err != nil {
						log.Error(err)
					} else {

						// send back to the device the new settings
						deviceTopic := fmt.Sprintf("devices/%s", req.DeviceID)
						log.Infof("Sending message %s to %s", jsonData, deviceTopic)
						token := client.Publish(deviceTopic, 1, false, jsonData)
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
