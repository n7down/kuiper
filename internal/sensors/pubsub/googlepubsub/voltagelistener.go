package googlepubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"
)

func (p GooglePubSub) NewVoltageListener(ctx context.Context, listenerName string, subscription string) error {
	sub := p.client.Subscription(subscription)
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		m.Ack() // Acknowledge that we've consumed the message.

		p.logger.Infof("Received message: %s\n", m.Data)
		sensors := &sensors.VoltageSensor{}
		err := json.Unmarshal(m.Data, sensors)
		if err != nil {
			p.logger.Error(err.Error())
		}

		if err == nil {
			err = p.persistence.CreateVoltage(sensors)
			p.logger.Infof("Logged sensor: %v", sensors)
			if err != nil {
				p.logger.Error(err.Error())
			}
		}
	})
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%s on %s is connected", listenerName, subscription))
	return nil
}
