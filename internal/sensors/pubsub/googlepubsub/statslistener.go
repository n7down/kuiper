package googlepubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"
)

func (p GooglePubSub) NewStatsListener(ctx context.Context, listenerName string, subscription string) error {
	// Use a callback to receive messages via subscription1.
	sub := p.client.Subscription(subscription)
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		m.Ack()

		p.logger.Infof("Received message: %s\n", m.Data)
		sensors := &sensors.StatsSensor{}
		err := json.Unmarshal(m.Data, sensors)
		if err != nil {
			p.logger.Error(err.Error())
		}

		if err == nil {
			err = p.persistence.CreateStats(sensors)
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
