package settings

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/iota/internal/client/settings/request"
	"github.com/n7down/iota/internal/client/settings/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	settings_pb "github.com/n7down/iota/internal/pb/settings"
)

const (
	FIVE_MINUTES = 5 * time.Minute
)

type SettingsClient struct {
	settingsClient settings_pb.SettingsServiceClient
}

func NewSettingsClient(serverEnv string) (*SettingsClient, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	settingsConn, err := grpc.Dial(serverEnv, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		return nil, err
	}

	client := &SettingsClient{
		settingsClient: settings_pb.NewSettingsServiceClient(settingsConn),
	}
	return client, nil
}

func (client *SettingsClient) GetSettings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.GetSettingsRequest
		res response.GetSettingsResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.settingsClient.GetSettings(ctx, &settings_pb.GetSettingsRequest{DeviceID: req.DeviceID})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if r == nil {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.GetSettingsResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

func (client *SettingsClient) SetSettings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.SetSettingsRequest
		res response.SetSettingsResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.settingsClient.SetSettings(ctx, &settings_pb.SetSettingsRequest{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: int32(req.DeepSleepDelay),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if r == nil {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.SetSettingsResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}
