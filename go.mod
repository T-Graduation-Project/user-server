module github.com/T-Graduation-Project/user-server

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gogf/gf v1.15.1
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.34.0 // indirect
)

replace google.golang.org/grpc v1.34.0 => google.golang.org/grpc v1.26.0
