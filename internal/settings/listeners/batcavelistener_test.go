// +build unit

package listeners

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/n7down/kuiper/internal/logger/blanklogger"
	"github.com/n7down/kuiper/internal/settings/listeners"
	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/n7down/kuiper/internal/settings/persistence/mock"
	"github.com/stretchr/testify/assert"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Test_BatCaveSettingsListenerMessageHandler_Should_Return_When_Message_And_Persistence_Settings_Are_The_Same(t *testing.T) {
	// testCases := []struct {
	// 	name            string
	// 	req             BatCaveSettingRequest
	// 	persistence     persistence.BatCaveSetting
	// 	expectedValue   bool
	// 	expectedSetting response.BatCaveSettingResponse
	// }{
	// 	{
	// 		name: "Deep_Sleep_Delay_Are_Equal",
	// 		req: BatCaveSettingRequest{
	// 			DeepSleepDelay: 15,
	// 		},
	// 		persistence: persistence.BatCaveSetting{
	// 			DeepSleepDelay: 15,
	// 		},
	// 		expectedValue: true,
	// 		expectedSetting: response.BatCaveSettingResponse{
	// 			DeepSleepDelay: 15,
	// 		},
	// 	},
	// 	{
	// 		name: "Deep_Sleep_Delay_Has_Changes_In_Persistence",
	// 		req: BatCaveSettingRequest{
	// 			DeepSleepDelay: 15,
	// 		},
	// 		persistence: persistence.BatCaveSetting{
	// 			DeepSleepDelay: 20,
	// 		},
	// 		expectedValue: false,
	// 		expectedSetting: response.BatCaveSettingResponse{
	// 			DeepSleepDelay: 20,
	// 		},
	// 	},
	// }

	// for _, testCase := range testCases {
	// 	t.Run(testCase.name, func(t *testing.T) {
	// 		isEqual, res := testCase.req.IsEqual(testCase.persistence)
	// 		assert.Equal(t, testCase.expectedValue, isEqual, "should have the same boolean value")
	// 		assert.Equal(t, testCase.expectedSetting, res, "should have the same setting")
	// 	})
	// }
	assert.Fail(t, "Not implemented")
}

type MockToken struct {
	MockWait        func() bool
	MockWaitTimeout func(time.Duration) bool
	MockError       func() error
}

func (t *MockToken) Wait() bool {
	return t.MockWait()
}

func (t *MockToken) WaitTimeout(d time.Duration) bool {
	return t.MockWaitTimeout(d)
}

func (t *MockToken) Error() error {
	return t.Error()
}

var mockToken = &MockToken{}

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

func Test_BatCaveSettingsListenerMessageHandler_Should_Publish_Changes_When_Message_And_Persistence_Are_Not_The_Same(t *testing.T) {

	var (
		publishedCalled bool = false
		publishedData   interface{}
		mac             string = "111111111111"
	)

	mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()
	mockPersistence := mock.NewMockPersistence(mockCtrl)
	mockPersistence.EXPECT().GetBatCaveSetting(mac).Return(
		false,
		persistence.BatCaveSetting{
			DeviceID:       mac,
			DeepSleepDelay: 30,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		})

	mockToken := &MockToken{
		MockWait: func() bool {
			return false
		},
		MockWaitTimeout: func(time.Duration) bool {
			return false
		},
		MockError: func() error {
			return nil
		},
	}

	mockClient := &MockClient{
		MockIsConnected: func() bool {
			return true
		},
		MockIsConnectionOpen: func() bool {
			return true
		},
		MockConnect: func() mqtt.Token {
			return mockToken
		},
		MockDisconnect: func(quiesce uint) {
		},
		MockPublish: func(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
			publishedCalled = true
			publishedData = payload
			return mockToken
		},
		MockSubscribe: func(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
			return mockToken
		},
		MockSubscribeMultiple: func(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
			return mockToken
		},
		MockUnsubscribe: func(topics ...string) mqtt.Token {
			return mockToken
		},
		MockAddRoute: func(topic string, callback mqtt.MessageHandler) {
		},
		MockOptionsReader: func() mqtt.ClientOptionsReader {
			return mqtt.ClientOptionsReader{}
		},
	}

	mockMessage := &MockMessage{
		MockDuplicate: func() bool {
			return false
		},
		MockQos: func() byte {
			return byte(0)
		},
		MockRetained: func() bool {
			return false
		},
		MockTopic: func() string {
			return ""
		},
		MockMessageID: func() uint16 {
			return 0
		},
		MockPayload: func() []byte {
			return []byte(`{"m":"111111111111","s":25}`)
		},
		MockAck: func() {
		},
	}

	log := blanklogger.NewBlankLogger()
	settingsListenersEnv := listeners.NewSettingsListenersEnv(mockPersistence, log)
	settingsListenersEnv.BatCaveSettingsListenerMessageHandler(mockClient, mockMessage)

	publishedDataExpected := []byte(`{"s":30}`)
	publishedDataActual := publishedData

	publishedCalledExpected := true
	publishedCalledActual := publishedCalled

	assert.Equal(t, publishedCalledExpected, publishedCalledActual, fmt.Sprintf("publishedCalled should return true instead returned %t", publishedCalled))
	assert.Equal(t, publishedDataExpected, publishedDataActual, fmt.Sprintf("publishedData should return %s instead returned %s", publishedDataExpected, publishedDataActual))
}

func Test_BatCaveSettingsListenerMessageHandler_Should_Send_Default_Settings_When_Setting_Does_Not_Exist_In_Persistence(t *testing.T) {
	assert.Fail(t, "Not implemented")
}
