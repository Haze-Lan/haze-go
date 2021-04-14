package main

import (
	"context"
	"fmt"
	"github.com/Haze-Lan/haze-go/haze-provider/endpoint"
	"github.com/Haze-Lan/haze-go/haze-provider/model"
	"github.com/Haze-Lan/haze-go/server"
	"google.golang.org/grpc"
	"log"
)

func main() {
	server.Run()
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", "etcd", "account-haze"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	clien:= endpoint.NewAccountClient(conn)
	fmt.Println(clien.Authentication(context.TODO(),&model.LoginRequest{Name: "22222",Pass: "22222222222222222"}))
}
