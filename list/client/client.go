package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-service/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"sync"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	var wg sync.WaitGroup


	//初始化客户端
	client := hello.NewListServiceClient(conn)
	req := &hello.ListRequest{Name: "zhourenjie"}

	for i := 0;i< 10000;i++ {
		wg.Add(1)
		go func() {
			list, err := client.GetList(context.Background(), req)
			if err != nil {
				grpclog.Fatalln(err)
			}

			jsonstr,_ := json.Marshal(list)
			fmt.Println(string(jsonstr))
			wg.Done()
		}()
	}

	wg.Wait()




}
