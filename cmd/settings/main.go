package main

import (
	"fmt"
	"net"
	"os"

	settings_pb "github.com/n7down/iota/internal/pb/settings"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Listening on port: %s\n", port)
	grpcServer := grpc.NewServer()
	settings_pb.RegisterSettingsServiceServer(grpcServer, nil)
	grpcServer.Serve(lis)
}
