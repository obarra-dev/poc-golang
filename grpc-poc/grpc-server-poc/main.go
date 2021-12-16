package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/myuser/myrepo/myService"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type greaterServer struct{
	myService.UnimplementedGreeterServer
}

func NewGreaterServer() *greaterServer {
	return &greaterServer{}
}

func (s *greaterServer) SayHello(ctx context.Context, in *myService.HelloRequest) (*myService.HelloReply, error) {
	fmt.Println(in.Name)
	fmt.Println(in.Filters)
	return &myService.HelloReply{Message: in.Name + " world from the greaterServer!!"}, nil
}

type imageServer struct{
	myService.UnimplementedImageServiceServer
}

func NewImageServer() *imageServer {
	return &imageServer{}
}

func (s *imageServer) UploadImage(stream myService.ImageService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	fmt.Println(req.GetName())

	var imageBuffer bytes.Buffer
	var imageSize uint

	for {
		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			log.Println(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := uint(len(chunk))

		log.Printf("received a chunk with size: %d", size)

		//1 MB = 2^20 bytes = 1 << 20 bytes
		const maxImageSize = 1 << 20

		imageSize += size
		if imageSize > maxImageSize {
			log.Println(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize)
		}
		_, err = imageBuffer.Write(chunk)
		if err != nil {
			log.Println(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}

	imageID, err := SaveFile(req.GetName(), imageBuffer)

	res := &myService.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		log.Println(status.Errorf(codes.Unknown, "cannot send response: %v", err))
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	log.Printf("saved image with id: %s, size: %d", res.Id, res.Size)
	return nil
}


func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC greaterServer object
	s := grpc.NewServer()
	// Attach the Greeter service to the greaterServer
	myService.RegisterGreeterServer(s, NewGreaterServer())
	myService.RegisterImageServiceServer(s, NewImageServer())
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}