package mock

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MockClient struct {
	MockIsConnected       func() bool
	MockIsConnectionOpen  func() bool
	MockConnect           func() mqtt.Token
	MockDisconnect        func(quiesce uint)
	MockPublish           func(topic string, qos byte, retained bool, payload interface{}) mqtt.Token
	MockSubscribe         func(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token
	MockSubscribeMultiple func(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token
	MockUnsubscribe       func(topics ...string) mqtt.Token
	MockAddRoute          func(topic string, callback mqtt.MessageHandler)
	MockOptionsReader     func() mqtt.ClientOptionsReader
}

func (c *MockClient) IsConnected() bool {
	return c.MockIsConnected()
}

func (c *MockClient) IsConnectionOpen() bool {
	return c.MockIsConnectionOpen()
}

func (c *MockClient) Connect() mqtt.Token {
	return c.Connect()
}

func (c *MockClient) Disconnect(quiesce uint) {
	c.Disconnect(quiesce)
}

func (c *MockClient) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return c.MockPublish(topic, qos, retained, payload)
}

func (c *MockClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	return c.Subscribe(topic, qos, callback)
}

func (c *MockClient) SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
	return c.SubscribeMultiple(filters, callback)
}

func (c *MockClient) Unsubscribe(topics ...string) mqtt.Token {
	return c.Unsubscribe(topics...)
}

func (c *MockClient) AddRoute(topic string, callback mqtt.MessageHandler) {
	c.AddRoute(topic, callback)
}

func (c *MockClient) OptionsReader() mqtt.ClientOptionsReader {
	return c.OptionsReader()
}
