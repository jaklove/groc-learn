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

type ListService struct {

}

var listService = ListService{}

//获取数据
func (l ListService) GetList(ctx context.Context,request *hello.ListRequest)(*hello.GetListReply,error) {
	resp := new(hello.GetListReply)
	resp.Pager = &hello.Pager{
		Page: 0,
		PageSize: 10,
		TotalRows: 10,
	}

	//定义
	var request_name string
	if request.Name == "zhourenjie"{
		request_name = "zhourenjie"
	}else {
		request_name = "周仁杰"
	}
	resp.List = append(resp.List,&hello.List{
		Name: request_name,
		State: 1,
		Id: 1,
	} )
	num++

	fmt.Println("被调用的服务:",num)
	return resp,nil
}
var num int

func main()  {
	listen, err := net.Listen("tcp", Address)
	if err != nil{
		grpclog.Fatalln(err)
	}

	//实例化service
	server := grpc.NewServer()
	hello.RegisterListServiceServer(server,listService)

	fmt.Println("Listen on " + Address)

	server.Serve(listen)
}
