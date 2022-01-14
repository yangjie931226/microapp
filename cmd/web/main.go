package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	httpserver "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"log"
	"net/http"
	"web/controller"
)

const SERVER_NAME = "go.micro.webservice"


func main() {
	err := controller.ImportTemplates()
	if err!= nil {
		log.Fatal(err)
	}
	srv := httpserver.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8889"),
	)

	http.Handle("/", http.RedirectHandler("/students", http.StatusPermanentRedirect))

	h := new(controller.StudentsHandler)
	http.Handle("/students", h)
	http.Handle("/students/", h)

	hd := srv.NewHandler(h)
	if err := srv.Handle(hd); err != nil {
		log.Fatal()
	}
	reg := consul.NewRegistry()
	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(reg),
	)
	service.Init()
	service.Run()
}



