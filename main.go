package main

import (
	"github.com/T-Graduation-Project/user-server/app/service"
	"github.com/T-Graduation-Project/user-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/micro/go-micro/v2"
)

var (
	log = g.Log()
)

func main() {
	server := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)
	server.Init()
	protobuf.RegisterUserHandler(server.Server(), new(service.UserApi))
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
