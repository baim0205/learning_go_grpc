package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/rohim/auth-service/database" // Import package database Anda
	pbAuth "github.com/rohim/auth-service/proto/protoc"
	"github.com/rohim/auth-service/server"
	"google.golang.org/grpc"
)

func main() {
	// Cek koneksi ke database
	// if err := database.Connect(); err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// log.Println("Connected to database successfully")

	// Muat variabel dari file .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Cek koneksi ke database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database successfully")

	// Cek Service On!!!
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
