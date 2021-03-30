package main

import (
	"context"
	"github.com/T-Graduation-Project/user-server/protobuf"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "127.0.0.1:8003"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewUserClient(conn)
	//// 获取书籍列表
	//req := &protobuf.SignUpReq{}
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

	req3 := &protobuf.AuthReq{
		Token: r2.Token,
	}
	r3, err := c.Auth(context.Background(), req3)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r3)
}
