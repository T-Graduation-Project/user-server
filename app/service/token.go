package service

import (
	"context"
	"github.com/T-Graduation-Project/user-server/protobuf"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
)

var (
	db         = g.DB("default")
	log        = g.Log()
	privateKey = []byte("`xs#a_1-!")
)

type CustomClaims struct {
	Username string
	Password string
	jwt.StandardClaims
}

func (u UserApi) Auth(
	ctx context.Context, req *protobuf.AuthReq, rsp *protobuf.AuthRsp) error {
	rsp.Code = 1
	rsp.Msg = "Auth filed"
	token, err := jwt.ParseWithClaims(req.Token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if err != nil || !token.Valid {
		return err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		rsp.Msg = "claims change error"
		return err
	}
	rsp.Username = claims.Username
	rsp.Password = claims.Password
	rsp.Code = 0
	rsp.Msg = "success"
	return err
}
