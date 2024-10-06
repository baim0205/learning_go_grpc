package main

import (
	"net"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/rohim/auth-service/database"
	pbAuth "github.com/rohim/auth-service/proto/protoc"
	"github.com/rohim/auth-service/server"
	"google.golang.org/grpc"
)

// TestMain digunakan untuk setup awal (tanpa *testing.M)
func TestMain(m *testing.M) {
	_ = godotenv.Load() // Muat environment variables dari file .env jika diperlukan
}

// TestAuthServiceServer menguji apakah server gRPC berjalan dengan baik
func TestAuthServiceServer(t *testing.T) {
	// Setup untuk gomock controller menggunakan *testing.T
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Log bahwa tes sudah dimulai
	t.Log("Starting TestAuthServiceServer")

	// Cek koneksi ke database
	t.Log("Checking database connection...")
	if err := database.Connect(); err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	} else {
		t.Log("Database connected successfully")
	}

	// Setup untuk listener
	t.Log("Setting up TCP listener on port 50051...")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	// Setup untuk gRPC server
	t.Log("Setting up gRPC server...")
	grpcServer := grpc.NewServer()
	authServer := &server.AuthServer{}
	pbAuth.RegisterAuthServiceServer(grpcServer, authServer)

	// Channel untuk menangkap error dari goroutine
	errChan := make(chan error, 1)

	// Jalankan gRPC server dalam goroutine
	t.Log("Starting gRPC server in a separate goroutine...")
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			errChan <- err // Kirim error ke channel
		}
		close(errChan)
	}()

	// Beri waktu agar server gRPC bisa mulai
	time.Sleep(100 * time.Millisecond)

	// Cek apakah server gRPC dapat dijangkau
	t.Log("Testing connection to gRPC server...")
	conn, err := net.Dial("tcp", ":50051")
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()
	t.Log("gRPC server is reachable, test passed")

	// Tunggu sampai goroutine selesai atau error terjadi
	select {
	case err := <-errChan:
		if err != nil {
			t.Fatalf("gRPC server failed: %v", err)
		}
	default:
		// Server berjalan dengan baik
		t.Log("gRPC server running without errors")
	}
}
