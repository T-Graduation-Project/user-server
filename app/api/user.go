package api

import (
	"context"
	"github.com/T-Graduation-Project/user-server/app/service"
	"github.com/T-Graduation-Project/user-server/protobuf"
)

type UserApi struct{}

var User = &UserApi{}

func (u UserApi) SignUp(
	c context.Context, req *protobuf.SignUpReq) (*protobuf.SignUpRsp, error) {
	rsp, err := service.SignUp(req)
	return rsp, err
}

func (u UserApi) SignIn(
	c context.Context, req *protobuf.SignInReq) (*protobuf.SignInRsp, error) {
	rsp, err := service.SignIn(req)
	return rsp, err
}

func (u UserApi) Auth(
	c context.Context, req *protobuf.AuthReq) (*protobuf.AuthRsp, error) {
	rsp, err := service.Auth(req)
	return rsp, err
}
