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

	// 认证 token
	req3 := &protobuf.AuthReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QiLCJzZXgiOiLkv53lr4YiLCJwaG9uZSI6IjEyMyIsInJvbGUiOiLor7vogIUiLCJleHAiOjE2MTgwNzA1MTQsImlzcyI6InRva2VuX3NlcnZpY2UifQ.hU6cEbdYIbxefViZGBIf1ddV_FaWu9_6CkieXsXqlJU",
	}
	r3, err := c.Auth(context.Background(), req3)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r3)
}
