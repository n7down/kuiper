package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIGateway struct{}

func NewAPIGateway() *APIGateway {
	return &APIGateway{}
}

func (g *APIGateway) InitV1Routes(r *gin.Engine) error {
	v1 := r.Group("/api/v1")
	deviceGroup := v1.Group("/device")
	{
		deviceGroup.PUT("/:id", nil)
		deviceGroup.GET("/:id", nil)
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
