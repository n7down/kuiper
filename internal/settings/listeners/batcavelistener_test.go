//+build unit

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
	mockobject "github.com/n7down/kuiper/internal/settings/mock"
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

func Test_BatCaveSettingsListenerMessageHandler_Should_Publish_Changes_When_Message_And_Persistence_Are_Not_The_Same(t *testing.T) {

	// FIXME: pass in multiple test cases
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
