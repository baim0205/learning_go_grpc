package main

import (
	// server "auth-service/server"
	"log"
	"net"

	pbAuth "github.com/rohim/auth-service/proto/protoc"

	"github.com/rohim/auth-service/server"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authServer := &server.AuthServer{}

	pbAuth.RegisterAuthServiceServer(grpcServer, authServer)

	log.Println("Auth service is running on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
