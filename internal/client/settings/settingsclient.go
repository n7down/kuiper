package settings

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/kuiper/internal/client/settings/request"
	"github.com/n7down/kuiper/internal/client/settings/response"
	"google.golang.org/grpc"

	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
)

const (
	FIVE_MINUTES = 5 * time.Minute
)

type SettingsClient struct {
	settingsClient settings_pb.SettingsServiceClient
}

func NewSettingsClient(serverEnv string) (*SettingsClient, error) {
	settingsConn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &SettingsClient{
		settingsClient: settings_pb.NewSettingsServiceClient(settingsConn),
	}
	return client, nil
}

func NewSettingsClientWithMock(mockSettingsServiceClient settings_pb.SettingsServiceClient) *SettingsClient {
	client := &SettingsClient{
		settingsClient: mockSettingsServiceClient,
	}
	return client
}

func (client *SettingsClient) CreateBatCaveSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateBatCaveSettingRequest
		res response.CreateBatCaveSettingResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	req.DeviceID = strings.ToLower(req.DeviceID)

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.settingsClient.CreateBatCaveSetting(ctx, &settings_pb.CreateBatCaveSettingRequest{DeviceID: req.DeviceID, DeepSleepDelay: req.DeepSleepDelay})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.CreateBatCaveSettingResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

func (client *SettingsClient) GetBatCaveSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.GetBatCaveSettingRequest
		res response.GetBatCaveSettingResponse
	)

	deviceID := c.Params.ByName("device_id")

	req = request.GetBatCaveSettingRequest{
		DeviceID: deviceID,
	}

	req.DeviceID = strings.ToLower(req.DeviceID)

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.settingsClient.GetBatCaveSetting(ctx, &settings_pb.GetBatCaveSettingRequest{DeviceID: req.DeviceID})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if r.DeviceID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.GetBatCaveSettingResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

func (client *SettingsClient) UpdateBatCaveSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.UpdateBatCaveSettingRequest
		res response.UpdateBatCaveSettingResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	deviceID := c.Params.ByName("device_id")

	req = request.UpdateBatCaveSettingRequest{
		DeviceID:       deviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	req.DeviceID = strings.ToLower(req.DeviceID)

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.settingsClient.UpdateBatCaveSetting(ctx, &settings_pb.UpdateBatCaveSettingRequest{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if r.DeviceID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.UpdateBatCaveSettingResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}
