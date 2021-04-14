package main

import (
	"context"
	"fmt"
	endpoint2 "github.com/Haze-Lan/haze-go/examples/haze-provider/endpoint"
	model2 "github.com/Haze-Lan/haze-go/examples/haze-provider/model"
	"github.com/Haze-Lan/haze-go/server"
	"google.golang.org/grpc"
	"log"
)

func main() {
	go  server.Run()
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", "etcd", "account-haze"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	clien:= endpoint2.NewAccountClient(conn)
	fmt.Println(clien.Authentication(context.TODO(),&model2.LoginRequest{Name: "22222",Pass: "22222222222222222"}))
}
