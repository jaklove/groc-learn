package main

import (
	"context"
	"fmt"
	"go-service/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main()  {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil{
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	//初始化客户端
	client := hello.NewHelloClient(conn)

	// 调用方法
	req := &hello.HelloRequest{Name: "zhourenjie"}
	sayHello, err := client.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(sayHello.Message)
}
