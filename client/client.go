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

	//// 获取书籍列表
	//req := &protobuf.SignUpReq{}
	//r, err := c.SignUp(context.Background(), req)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r)
	//
	req2 := &protobuf.SignInReq{
		Username: "test",
		Password: "test",
	}
	r2, err := c.SignIn(context.Background(), req2)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r2)
	// 认证 token
	req3 := &protobuf.AuthReq{
		Token: r2.Token,
	}
	r3, err := c.Auth(context.Background(), req3)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r3)
}
