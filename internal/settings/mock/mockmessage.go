package mock

type MockMessage struct {
	MockDuplicate func() bool
	MockQos       func() byte
	MockRetained  func() bool
	MockTopic     func() string
	MockMessageID func() uint16
	MockPayload   func() []byte
	MockAck       func()
}

func (m *MockMessage) Duplicate() bool {
	return m.MockDuplicate()
}

func (m *MockMessage) Qos() byte {
	return m.MockQos()
}

func (m *MockMessage) Retained() bool {
	return m.Retained()
}

func (m *MockMessage) Topic() string {
	return m.Topic()
}

func (m *MockMessage) MessageID() uint16 {
	return m.MessageID()
}

func (m *MockMessage) Payload() []byte {
	return m.MockPayload()
}

func (m *MockMessage) Ack() {
	m.MockAck()
}
