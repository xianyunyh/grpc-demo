package main

import (
	"context"
	"gRPC-study/handler"
	pb "gRPC-study/models"
	"log"
	"net"
	"net/http"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpAddr  = ":8081"
	httpAddr = ":8080"
)

func main() {
	go func() {
		creategRPCServer(grpAddr)
	}()
	log.Println("start grpc server ", grpAddr)
	err := createHttpServer(grpAddr, httpAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("start gateway http server", httpAddr)

}

func creategRPCServer(addr string) error {
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpc_recovery.UnaryServerInterceptor(),
	))

	pb.RegisterPostServer(server, handler.NewPostHandler())
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return server.Serve(l)
}

func createHttpServer(grpcAddr string, httpAddr string) error {

	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterPostHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(httpAddr, mux)
}
