package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/burov/courses/web/03_grpc_hello/service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := service.NewGreeterClient(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		if text == "exit" {
			break
		}
		resp, err := client.Hello(context.Background(), &service.HelloRequest{Name: text})
		if err != nil {
			panic(err)
		}

		fmt.Println(resp.Message)
	}

	fmt.Println("Client connection closed!")
}
