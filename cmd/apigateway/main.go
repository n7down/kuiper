package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/n7down/iota/internal/apigateway"
	"github.com/n7down/iota/internal/client/settings"

	log "github.com/sirupsen/logrus"
)

func init() {
}

func main() {
	log.SetReportCaller(true)

	port := os.Getenv("PORT")
	settingsHost := os.Getenv("SETTINGS_HOST")

	settingsClient, err := settings.NewSettingsClient(settingsHost)
	if err != nil {
		log.Fatal(err)
	}

	apiGateway := apigateway.NewAPIGateway(settingsClient)
	router := gin.Default()

	err = apiGateway.InitV1Routes(router)
	if err != nil {
		log.Fatal(err)
	}

	routerPort := fmt.Sprintf(":%s", port)
	log.Infof("Listening on port: %s\n", port)
	err = apiGateway.Run(router, routerPort)
	if err != nil {
		log.Fatal(err)
	}
}
