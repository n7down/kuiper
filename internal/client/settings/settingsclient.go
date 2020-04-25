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

func (client *SettingsClient) CreateBatCaveSettings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateBatCaveSettingsRequest
		res response.CreateBatCaveSettingsResponse
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

	r, err := client.settingsClient.CreateBatCaveSettings(ctx, &settings_pb.CreateBatCaveSettingsRequest{DeviceID: req.DeviceID, DeepSleepDelay: req.DeepSleepDelay})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.CreateBatCaveSettingsResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

func (client *SettingsClient) GetBatCaveSettings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.GetBatCaveSettingsRequest
		res response.GetBatCaveSettingsResponse
	)

	deviceID := c.Params.ByName("device_id")

	req = request.GetBatCaveSettingsRequest{
		DeviceID: deviceID,
	}

	req.DeviceID = strings.ToLower(req.DeviceID)

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.settingsClient.GetBatCaveSettings(ctx, &settings_pb.GetBatCaveSettingsRequest{DeviceID: deviceID})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if r.DeviceID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.GetBatCaveSettingsResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

func (client *SettingsClient) UpdateBatCaveSettings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.UpdateBatCaveSettingsRequest
		res response.UpdateBatCaveSettingsResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	deviceID := c.Params.ByName("device_id")

	req = request.UpdateBatCaveSettingsRequest{
		DeviceID: deviceID,
	}

	req.DeviceID = strings.ToLower(req.DeviceID)

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.settingsClient.UpdateBatCaveSettings(ctx, &settings_pb.UpdateBatCaveSettingsRequest{
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

	res = response.UpdateBatCaveSettingsResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}
