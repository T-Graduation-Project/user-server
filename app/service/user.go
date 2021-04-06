package service

import (
	"context"
	"github.com/T-Graduation-Project/user-server/protobuf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserApi struct{}

func (u UserApi) SignUp(
	ctx context.Context, req *protobuf.SignUpReq, rsp *protobuf.SignUpRsp) error {
	rsp.Code = 1
	rsp.Msg = "sign up error"
	_, err := db.Table("user_info").Save(req)
	rsp.Code = 0
	rsp.Msg = "success"
	return err
}

func (u UserApi) SignIn(
	ctx context.Context, req *protobuf.SignInReq, rsp *protobuf.SignInRsp) error {
	rsp.Code = 1
	rsp.Msg = "sign in error"

	user, err := db.Table(
		"user_info").Where("username=?", req.Username).One()
	if user.Map()["password"] != req.Password {
		return err
	}
	// TODO 检查账号
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims{
		req.Username,
		req.Password,
		jwt.StandardClaims{
			Issuer:    "token_service", // 签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rsp.Token, err = jwtToken.SignedString(privateKey)
	rsp.Code = 0
	rsp.Msg = "success"
	return err
}
