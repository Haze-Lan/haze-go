package main

import (
	"context"
	"github.com/Haze-Lan/haze-go/api"
	"github.com/Haze-Lan/haze-go/server"
	"google.golang.org/grpc"
	"log"
	"net"

)




func main() {
	haze:=  server.Init()
	if err:=  haze.Run();err!=nil{
		log.Fatalf("failed to listen: %v", err)
	}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type server struct {
	api.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &api.HelloReply{Message: "Hello " + in.GetName()}, nil
}


