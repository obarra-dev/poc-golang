package main

import (
	"context"
	"fmt"
	"github.com/myuser/myrepo/myService"
	"log"
	"net"

	"google.golang.org/grpc"

)

type server struct{
	myService.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *myService.HelloRequest) (*myService.HelloReply, error) {
	fmt.Println(in.Name)
	fmt.Println(in.Filters)
	return &myService.HelloReply{Message: in.Name + " world!!"}, nil
}


func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	myService.RegisterGreeterServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}