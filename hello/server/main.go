package main

import (
	"context"
	"fmt"
	"go-service/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)


// 定义helloService并实现约定的接口
type helloService struct {}

var HelloService = helloService{}

func (h helloService)SayHello(ctx context.Context,in *hello.HelloRequest)(*hello.HelloResponse,error)  {
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.",in.Name)
	return resp,nil
}


func main()  {
	listen, err := net.Listen("tcp", Address)
	if err != nil{
		grpclog.Fatalf("Failed to listen: %v",err)
	}

	//实例化grpc server
	s := grpc.NewServer()

	hello.RegisterHelloServer(s,HelloService)
	fmt.Println("Listen on " + Address)
	s.Serve(listen)
}

