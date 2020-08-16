package googlepubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/n7down/kuiper/internal/settings/listeners/request"
	"github.com/n7down/kuiper/internal/settings/listeners/response"
	"github.com/n7down/kuiper/internal/settings/persistence"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

func (p GooglePubSub) BatCaveSettingsListenerMessageHandler(ctx context.Context, msg []byte) {
	p.logger.Infof("Received message: %s\n", msg)

	// unmashal payload
	var (
		req request.BatCaveSettingRequest
		res response.BatCaveSettingResponse
	)

	err := json.Unmarshal(msg, &req)
	if err != nil {
		p.logger.Error(err)
		return
	}

	// get the settings
	recordNotFound, settingInPersistence := p.persistence.GetBatCaveSetting(req.DeviceID)
	if recordNotFound {

		// send back default values
		res = response.GetBatCaveSettingDefault()

		newSetting := persistence.BatCaveSetting{
			DeviceID:       req.DeviceID,
			DeepSleepDelay: res.DeepSleepDelay,
		}

		// create the new setting
		p.persistence.CreateBatCaveSetting(newSetting)

	} else {

		// check for the differences in the settings
		var isEqual bool
		isEqual, res = req.IsEqual(settingInPersistence)
		p.logger.Infof("Settings are equal: %t - %v %v", isEqual, settingInPersistence, res)
		if isEqual {

			// settings are the same on the device and in persistence - return
			return
		}
	}

	json, err := json.Marshal(res)
	if err != nil {
		p.logger.Error(err)
		return
	}

	// send back to the device the new settings
	deviceTopic := fmt.Sprintf("devices/%s", req.DeviceID)
	p.logger.Infof("Sending message %s to %s", json, deviceTopic)
	topic := p.client.Topic(deviceTopic)
	topic.Publish(ctx, &pubsub.Message{
		Data: json,
	})
}

func (p GooglePubSub) NewBatCaveSettingsListener(listenerName string, subscription string) error {
	ctx := context.Background()
	sub := p.client.Subscription(subscription)
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		m.Ack() // Acknowledge that we've consumed the message.

		p.BatCaveSettingsListenerMessageHandler(ctx, m.Data)
	})
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%s on %s is connected", listenerName, subscription))
	return nil

}
