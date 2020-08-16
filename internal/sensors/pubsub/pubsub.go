package pubsub

import (
	"context"
)

type PubSub interface {
	NewBMP280Listener(ctx context.Context, listenerName string, subscription string) error
	NewDHT22Listener(ctx context.Context, listenerName string, subscription string) error
	NewHDC1080Listener(ctx context.Context, listenerName string, subscription string) error
	NewStatsListener(ctx context.Context, listenerName string, subscription string) error
	NewVoltageListener(ctx context.Context, listenerName string, subscription string) error
}
