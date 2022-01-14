package utils

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/transport/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
)

func GetMicroClient() client.Client {
	transport := grpc.NewTransport()
	consulReg := consul.NewRegistry()
	microService := micro.NewService(
		micro.Registry(consulReg),
		micro.Transport(transport),
	)
	return microService.Client()
}
