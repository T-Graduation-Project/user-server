package service

import (
	"github.com/T-Graduation-Project/user-server/protobuf"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"time"
)

var (
	db  = g.DB("default")
	log = g.Log()
)

//type Authable interface {
//	Decode(tokenStr string) (*CustomClaims, error)
//	Encode(user *pb.User) (string, error)
//}

var privateKey = []byte("`xs#a_1-!")

type CustomClaims struct {
	Username string
	Password string
	// 使用标准的 payload
	jwt.StandardClaims
}

func Auth(req *protobuf.AuthReq) (*protobuf.AuthRsp, error) {
	rsp := &protobuf.AuthRsp{
		Code: 1,
		Msg:  "Auth filed",
	}
	token, err := jwt.ParseWithClaims(req.Token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	log.Info("token:", token.Raw)
	if !token.Valid || err != nil {
		return rsp, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		rsp.Msg = "claims change error"
		return rsp, err
	}
	rsp.Username = claims.Username
	rsp.Password = claims.Password
	rsp.Code = 0
	rsp.Msg = "success"
	return rsp, err
}

func SignUp(req *protobuf.SignUpReq) (*protobuf.SignUpRsp, error) {
	rsp := &protobuf.SignUpRsp{
		Code: 1,
		Msg:  "sign up error",
	}
	_, err := db.Table("user_info").Save(req)
	return rsp, err
}

func SignIn(req *protobuf.SignInReq) (*protobuf.SignInRsp, error) {
	rsp := &protobuf.SignInRsp{
		Code: 1,
		Msg:  "sign in error-",
	}
	user, err := db.Table(
		"user_info").Where("username=?", req.Username).One()
	if user.Map()["password"] != req.Password {
		return rsp, err
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
	rsp.Msg = "success"
	return rsp, err
}
