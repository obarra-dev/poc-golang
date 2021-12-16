package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-client-poc/myService"
	"log"
	"time"
)

func main() {
	var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := myService.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := myService.HelloRequest{Name: "Omar"}
	feature, err := client.SayHello(ctx, &request)
	if err != nil {
		log.Fatalf("%v.SayHello(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}
