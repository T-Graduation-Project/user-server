package main

import (
	"context"
	"github.com/T-Graduation-Project/user-server/protobuf"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(micro.Name("borrow.client"))
	service.Init()
	c := protobuf.NewUserService("user", service.Client())

	//req := &protobuf.SignUpReq{
	//	Username:   "test2",
	//	Password:   "test",
	//	RePassword: "test",
	//	Sex:        "保密",
	//	Phone:      "123",
	//}
	//r, err := c.SignUp(context.Background(), req)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r)

	req2 := &protobuf.SignInReq{
		Username: "test",
		Password: "test",
	}
	r2, err := c.SignIn(context.Background(), req2)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r2)

	r3, err := c.GetPersonalInfo(context.Background(), &protobuf.GetPersonalInfoReq{
		Username: "test",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r3)
}
