package pubsub

type PubSub interface {
	NewBatCaveSettingsListener(listenerName string, mqttURL string) error
}
