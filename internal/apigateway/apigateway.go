package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n7down/iota/internal/client/settings"
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
	batCaveGroup := v1.Group("/bc")
	{
		// FIXME: get :id in business logic
		// FIXME: for example - deviceGroup.POST("/:id", g.settingsClient.SetSettings)
		// FIXME: and - deviceGroup.GET("/:id", g.settingsClient.SetSettings)
		batCaveGroup.POST("/set", g.settingsClient.SetBatCaveSettings)
		batCaveGroup.POST("/get", g.settingsClient.GetBatCaveSettings)
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
