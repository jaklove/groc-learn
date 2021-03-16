package main

import (
	"flag"
	"go-service/proto"
	"go-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

var port string

func init()  {
	flag.StringVar(&port, "port", "8003", "启动端口号")
	flag.Parse()
}

func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunHttpServer(port string)*http.Server  {
	server := http.NewServeMux()
	server.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_,_ = writer.Write([]byte(`pong`))
	})
	return &http.Server{
		Addr: ":"+port,
		Handler: server,
	}
}

func RunGrpcServer()*grpc.Server  {
	newServer := grpc.NewServer()
	proto.RegisterTagServiceServer(newServer,server.NewTagServer())
	reflection.Register(newServer)
	return newServer
}

func main()  {

}
