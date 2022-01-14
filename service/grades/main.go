package main

import (
	"github.com/asim/go-micro/v3"
	"grades/handler"
	pb "grades/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/transport/grpc/v3"
	"github.com/asim/go-micro/v3/logger"
)

func main() {
	reg := consul.NewRegistry()
	transport := grpc.NewTransport()
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.gradeservice"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Transport(transport),
		micro.Address(":30002"),
	)

	// Register handler
	if err := pb.RegisterGradesHandler(srv.Server(), new(handler.Grades));err != nil {
		logger.Fatal(err)
	}


	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
