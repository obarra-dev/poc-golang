package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"service/myService"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcListener, err := net.Listen("tcp", ":9090")
	if err != nil {
		os.Exit(-1)
	}

	grpcSrv := &someGRPCServer{}
	baseServer := grpc.NewServer()
	myService.RegisterYourServiceServer(baseServer, grpcSrv)
	myService.RegisterImageServiceServer(baseServer, NewImageServer() )
	go baseServer.Serve(grpcListener)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()


	err = mux.HandlePath("POST", "/files", handleBinaryFileUpload)
	if err != nil {
		log.Fatalln(err)
		return err
	}


	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = myService.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}


	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}


func handleBinaryFileUpload(w http.ResponseWriter, r *http.Request,  params map[string]string) {
	file, err := os.Create("./file.pdf")
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
