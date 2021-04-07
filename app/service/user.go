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
	if req.Password != req.RePassword {
		rsp.Msg = "repeat password inconsistent"
		return nil
	}
	flag, err := isExisted(req.Username)
	if err != nil {
		return err
	}
	if flag == true {
		rsp.Msg = "username exist"
		return err
	}
	user := &protobuf.UserInfo{
		Username: req.Username,
		Password: req.Password,
		Sex:      req.Sex,
		Phone:    req.Phone,
		Role:     "读者",
	}
	_, err = db.Table("user_info").Save(user)
	if err != nil {
		return err
	}
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
	userInfo := protobuf.UserInfo{
		Username: user.Map()["username"].(string),
		Password: user.Map()["password"].(string),
		Sex:      user.Map()["sex"].(string),
		Phone:    user.Map()["phone"].(string),
		Role:     user.Map()["role"].(string),
	}
	rsp.Token = generateToken(userInfo)
	rsp.Code = 0
	rsp.Msg = "success"
	return err
}

func (u UserApi) GetPersonalInfo(
	ctx context.Context, req *protobuf.GetPersonalInfoReq, rsp *protobuf.GetPersonalInfoRsp) error {
	//rsp.Code = 1
	//rsp.Msg = "sign up error"
	//flag, err := isExisted(req.Username)
	//if err != nil {
	//	return err
	//}
	//if flag == false {
	//	rsp.Msg = "user not exist"
	//	return err
	//}
	//rsp.Code = 0
	//rsp.Msg = "success"
	return nil
}

// 检查某一用户是否存在，判断方式：
func isExisted(username string) (bool, error) {
	results, err := db.Table("user_info").All("username=?", username)
	if err != nil {
		log.Error(err)
		return true, err
	}
	if results != nil {
		log.Println("Exist user:", results)
		return true, nil
	}
	return false, nil
}

func generateToken(user protobuf.UserInfo) string {
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    "token_service", // 签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := jwtToken.SignedString(privateKey)
	if err != nil {
		return ""
	}
	return stringToken
}
