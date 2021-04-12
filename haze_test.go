package main

import (
	"context"
	"github.com/Haze-Lan/haze-go/provide/endpoint"
	"github.com/Haze-Lan/haze-go/provide/model"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:80"
)

func Test(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := endpoint.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Authentication(ctx, &model.LoginRequest{Name: "33", Pass: "222"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
