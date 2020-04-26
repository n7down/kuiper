package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n7down/kuiper/internal/client/settings"
)

type APIGateway struct {
	settingsClient *settings.SettingsClient
}

func NewAPIGateway(s *settings.SettingsClient) *APIGateway {
	return &APIGateway{
		settingsClient: s,
	}
}

func (g *APIGateway) InitV1Routes(r *gin.Engine) error {
	v1 := r.Group("/api/v1")
	deviceGroup := v1.Group("/settings")
	{
		deviceGroup.POST("/bc", g.settingsClient.CreateBatCaveSetting)
		deviceGroup.GET("/bc/:device_id", g.settingsClient.GetBatCaveSetting)
		deviceGroup.PUT("/bc/:device_id", g.settingsClient.UpdateBatCaveSetting)
	}
	return nil
}

func (g *APIGateway) Run(router *gin.Engine, port string) error {
	err := http.ListenAndServe(port, router)
	if err != nil {
		return err
	}
	return nil
}
