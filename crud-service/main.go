package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/rohim/crud-service/database"
	pbCrud "github.com/rohim/crud-service/proto/protoc"
	"github.com/rohim/crud-service/server"
	"google.golang.org/grpc"
)

func main() {
	// // Connect to the database
	// database.Connect()

	// // Cek koneksi ke database
	// if err := database.Connect(); err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// log.Println("Connected to database successfully")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Cek koneksi ke database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database successfully")

	// Start listening on port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create new gRPC server
	grpcServer := grpc.NewServer()

	// Register CRUD service
	pbCrud.RegisterCRUDServiceServer(grpcServer, &server.CrudServer{})

	log.Println("CRUD Service is running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
