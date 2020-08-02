package settings

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/n7down/kuiper/internal/mock"
	"github.com/stretchr/testify/assert"

	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
)

func Test_CreateBatCaveSetting(t *testing.T) {
	var (
		deviceID       string = "001100110011"
		deepSleepDelay uint32 = 15
		expectedCode          = http.StatusOK
		// got          gin.H
		// want         = make(map[string]interface{})
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{
		gin.Param{
			Key: "DeviceID", Value: deviceID,
		},
		gin.Param{
			Key: "DeepSleepDelay", Value: fmt.Sprint(deepSleepDelay),
		},
	}

	mockCtrl := gomock.NewController(t)
	mockSettingsServiceClient := mock.NewMockSettingsServiceClient(mockCtrl)

	settingsClient := NewSettingsClientWithMock(mockSettingsServiceClient)

	mockSettingsServiceClient.EXPECT().CreateBatCaveSetting(
		c,
		&settings_pb.CreateBatCaveSettingRequest{
			DeviceID:       deviceID,
			DeepSleepDelay: deepSleepDelay,
		},
	).Return(
		&settings_pb.CreateBatCaveSettingResponse{
			DeviceID:       deviceID,
			DeepSleepDelay: deepSleepDelay,
		}, nil,
	)

	settingsClient.CreateBatCaveSetting(c)

	// router.POST("/bc", settingsClient.CreateBatCaveSetting)

	// reqParam := fmt.Sprintf(`{"deviceID":"%s","deepSleepDelay":%d}`, deviceID, deepSleepDelay)

	// req, err := http.NewRequest("POST", "/bc", strings.NewReader(string(reqParam)))
	// assert.NoError(t, err)

	// router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	log.Print(string(w.Body.Bytes()))

	// err = json.Unmarshal(w.Body.Bytes(), &got)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// assert.Equal(t, want, got)
	assert.Fail(t, "not implemented")
}

func Test_GetBatCaveSetting(t *testing.T) {
	assert.Fail(t, "not implemented")
}

func Test_UpdateBatCaveSetting(t *testing.T) {
	assert.Fail(t, "not implemented")
}
