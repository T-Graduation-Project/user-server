# demo-server

## RUN
~~~bash
# 本地运行
protoc --gofast_out=plugins=grpc:. protocol/borrow.proto
go run main.go

# 构建镜像并运行
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
docker build -t demo-server .
docker run -p 8199:8199 -d --name "demo-server" demo-server
docker exec -it bash demo-server
# 上传镜像
docker push 
~~~