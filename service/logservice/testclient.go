package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/transport/grpc/v3"
	"github.com/asim/go-micro/v3"
	pb "logservice/proto"
)

func main() {
	reg := consul.NewRegistry()
	transport := grpc.NewTransport()

	service := micro.NewService(
		micro.Registry(reg),
		micro.Transport(transport),
	)

	// parse command line flags
	service.Init()
	// Create new greeter client
	logService := pb.NewLogService("go.micro.logservice", service.Client())

	// Call the greeter
	_, err := logService.WriteLog(context.TODO(), &pb.WriteLogRequest{Message: "test log"})
	if err != nil {
		fmt.Println(err)
		return
	}

}
