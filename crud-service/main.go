package main

import (
	"crud-service/server"
	"log"
	"net"

	pbCrud "github.com/rohim/crud-service/proto/protoc"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	crudServer := &server.CrudServer{}

	pbCrud.RegisterCRUDServiceServer(grpcServer, crudServer)

	log.Println("CRUD service is running on port :50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
