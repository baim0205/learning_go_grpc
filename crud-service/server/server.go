package server

import (
	"context"
	"fmt"
	"log"

	pbCrud "github.com/rohim/crud-service/proto/protoc"
	pbAuth "github.com/rohim/learning_go_grpc/crud-service/proto/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type CrudServer struct {
	pbCrud.UnimplementedCRUDServiceServer
}

func validateToken(token string) (bool, error) {
	// Connect to AuthService for token validation
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to AuthService: %v", err)
	}
	defer conn.Close()

	client := pbAuth.NewAuthServiceClient(conn)
	resp, err := client.ValidateToken(context.Background(), &pbAuth.ValidateTokenRequest{
		Token: token,
	})

	if err != nil {
		return false, status.Errorf(status.Code(err), "AuthService validation error: %v", err)
	}

	return resp.GetValid(), nil
}

func (s *CrudServer) Create(ctx context.Context, req *pbCrud.CreateRequest) (*pbCrud.CreateResponse, error) {
	token := req.GetToken()

	valid, err := validateToken(token)
	if err != nil || !valid {
		return nil, fmt.Errorf("invalid token")
	}

	data := req.GetData()
	log.Printf("Creating data: %s", data)

	return &pbCrud.CreateResponse{Message: "Data created successfully"}, nil
}

// Similarly, implement Read, Update, Delete with token validation
