package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/n7down/iota/internal/apigateway"
	"github.com/n7down/iota/internal/utils"

	log "github.com/sirupsen/logrus"
)

func init() {
}

func main() {
	log.SetReportCaller(true)

	port := utils.GetEnv("PORT", "8080")

	apiGateway := apigateway.NewAPIGateway()
	router := gin.Default()

	err := apiGateway.InitV1Routes(router)
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
