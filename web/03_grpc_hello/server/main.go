package main

import (
	"context"
	"fmt"
	"net"

	"github.com/burov/courses/web/03_grpc_hello/service"
	"google.golang.org/grpc"
)

//GreeterServer is Implementation of the rpc server
type GreeterServer struct{}

//Hello is a simple rpc method
func (GreeterServer) Hello(ctx context.Context, in *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{Message: fmt.Sprintf("Hello %s!", in.Name)}, nil
}

func main() {
	lsn, err := net.Listen("tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}
	defer lsn.Close()

	greeter := GreeterServer{}
	s := grpc.NewServer()

	service.RegisterGreeterServer(s, greeter)

	fmt.Println("Start server on 0.0.0.0:8080")
	if err := s.Serve(lsn); err != nil {
		panic(err)
	}

	fmt.Println("Server stopped")
}
