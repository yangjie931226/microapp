package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/transport/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"logservice/handler"
	pb "logservice/proto"
)



func main() {
	handler.Run("./logservice.log")
	reg := consul.NewRegistry()
	transport := grpc.NewTransport()
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.logservice"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Transport(transport),
		micro.Address(":30001"),
	)



	// Register handler
	if err := pb.RegisterLogServiceHandler(srv.Server(), new(handler.Logservice)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
