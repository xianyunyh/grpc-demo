package main

import (
	"context"
	pb "gRPC-study/models"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serve struct {
	pb.UnimplementedUserServer
}

func (s *serve) GetName(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{Id: req.Id, Name: time.Now().Format("15:04:05")}, nil
}

func main() {

	server := grpc.NewServer()
	pb.RegisterUserServer(server, &serve{})
	grpAddr := ":8080"
	l, err := net.Listen("tcp", grpAddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	go func() {
		err = server.Serve(l)
	}()
	createHttpServer(grpAddr)

}

func createHttpServer(grpcPoint string) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterUserHandlerFromEndpoint(ctx, mux, grpcPoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}
