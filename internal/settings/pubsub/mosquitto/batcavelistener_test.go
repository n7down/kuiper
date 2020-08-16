// +build unit,!integration

package mosquitto

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/n7down/kuiper/internal/logger/blanklogger"
	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/n7down/kuiper/internal/settings/persistence/mock"
	"github.com/stretchr/testify/assert"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	mockobject "github.com/n7down/kuiper/internal/mock"
)

func Test_BatCaveSettingsListenerMessageHandler_Should_Return_When_Message_And_Persistence_Settings_Are_The_Same(t *testing.T) {
	var (
		publishedCalled bool   = false
		mac             string = "111111111111"
		deepSleepDelay  uint32 = 30
	)

	mockCtrl := gomock.NewController(t)
	mockPersistence := mock.NewMockPersistence(mockCtrl)
	mockPersistence.EXPECT().GetBatCaveSetting(mac).Return(
		false,
		persistence.BatCaveSetting{
			DeviceID:       mac,
			DeepSleepDelay: deepSleepDelay,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		})

	mockToken := &mockobject.MockToken{
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

	mockClient := &mockobject.MockClient{
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

	mockMessage := &mockobject.MockMessage{
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
			return []byte(`{"m":"111111111111","s":30}`)
		},
		MockAck: func() {
		},
	}

	log := blanklogger.NewBlankLogger()
	pubSub := NewMosquittoPubSub(mockPersistence, log)
	pubSub.BatCaveSettingsListenerMessageHandler(mockClient, mockMessage)

	publishedCalledExpected := false
	publishedCalledActual := publishedCalled

	assert.Equal(t, publishedCalledExpected, publishedCalledActual, fmt.Sprintf("publishedCalled should return true instead returned %t", publishedCalled))
}

func Test_BatCaveSettingsListenerMessageHandler_Should_Publish_Changes_When_Message_And_Persistence_Are_Not_The_Same(t *testing.T) {
	var (
		publishedCalled bool = false
		publishedData   interface{}
		mac             string = "111111111111"
	)

	mockCtrl := gomock.NewController(t)
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

	mockToken := &mockobject.MockToken{
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

	mockClient := &mockobject.MockClient{
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

	mockMessage := &mockobject.MockMessage{
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
	pubSub := NewMosquittoPubSub(mockPersistence, log)
	pubSub.BatCaveSettingsListenerMessageHandler(mockClient, mockMessage)

	publishedDataActual := publishedData
	publishedDataExpected := []byte(`{"s":30}`)

	publishedCalledExpected := true
	publishedCalledActual := publishedCalled

	assert.Equal(t, publishedCalledExpected, publishedCalledActual, fmt.Sprintf("publishedCalled should return true instead returned %t", publishedCalled))
	assert.Equal(t, publishedDataExpected, publishedDataActual, fmt.Sprintf("publishedData should return %s instead returned %s", publishedDataExpected, publishedDataActual))
}

func Test_BatCaveSettingsListenerMessageHandler_Should_Send_Default_Settings_When_Setting_Does_Not_Exist_In_Persistence(t *testing.T) {
	var (
		publishedCalled bool = false
		publishedData   interface{}
		mac             string = "111111111111"
	)

	mockCtrl := gomock.NewController(t)
	mockPersistence := mock.NewMockPersistence(mockCtrl)
	mockPersistence.EXPECT().GetBatCaveSetting(mac).Return(
		true,
		persistence.BatCaveSetting{})

	mockPersistence.EXPECT().CreateBatCaveSetting(

		persistence.BatCaveSetting{
			DeviceID:       mac,
			DeepSleepDelay: 15,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		},
	)

	mockToken := &mockobject.MockToken{
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

	mockClient := &mockobject.MockClient{
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

	mockMessage := &mockobject.MockMessage{
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
	pubSub := NewMosquittoPubSub(mockPersistence, log)
	pubSub.BatCaveSettingsListenerMessageHandler(mockClient, mockMessage)

	publishedDataActual := publishedData
	publishedDataExpected := []byte(`{"s":15}`)

	publishedCalledExpected := true
	publishedCalledActual := publishedCalled

	assert.Equal(t, publishedCalledExpected, publishedCalledActual, fmt.Sprintf("publishedCalled should return true instead returned %t", publishedCalled))
	assert.Equal(t, publishedDataExpected, publishedDataActual, fmt.Sprintf("publishedData should return %s instead returned %s", publishedDataExpected, publishedDataActual))
}
