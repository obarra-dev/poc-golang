package main

import (
	"bufio"
	"context"
	"google.golang.org/grpc"
	"grpc-client-poc/myService"
	"io"
	"log"
	"os"
	"path/filepath"
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
	greeterClient := myService.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := myService.HelloRequest{Name: "Omar"}
	feature, err := greeterClient.SayHello(ctx, &request)
	if err != nil {
		log.Fatalf("%v.SayHello(_) = _, %v: ", greeterClient, err)
	}
	log.Println(feature)

	imageClient := myService.NewImageServiceClient(conn)
	uploadImage(imageClient,  "file/give/sample.pdf")
}

func uploadImage(client  myService.ImageServiceClient, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.UploadImage(ctx, )
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := &myService.UploadImageRequest{
		Data: &myService.UploadImageRequest_Name{Name: filepath.Ext(imagePath),
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &myService.UploadImageRequest{
			Data: &myService.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())

}