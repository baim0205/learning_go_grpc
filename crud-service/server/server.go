package server

import (
	"context"
	"database/sql"
	"log"

	"github.com/rohim/crud-service/database"
	pbCrud "github.com/rohim/crud-service/proto/protoc"
	pbAuth "github.com/rohim/crud-service/proto/protoc-auth"
	"google.golang.org/grpc"
)

type CrudServer struct {
	pbCrud.UnimplementedCRUDServiceServer
}

func validateToken(token string) (bool, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return false, err
	}
	defer conn.Close()

	authClient := pbAuth.NewAuthServiceClient(conn)
	resp, err := authClient.ValidateToken(context.Background(), &pbAuth.ValidateTokenRequest{
		Token: token,
	})

	if err != nil || !resp.Valid {
		return false, err
	}
	return true, nil
}

// Membuat fungsi CRUD!!!!! Bang
func (s *CrudServer) Create(ctx context.Context, req *pbCrud.CreateRequest) (*pbCrud.CreateResponse, error) {
	token := req.GetToken()

	valid, err := validateToken(token)
	if err != nil || !valid {
		return nil, err
	}

	data := req.GetData()
	_, err = database.DB.Exec("INSERT INTO items (name) VALUES (?)", data)
	if err != nil {
		return nil, err
	}

	log.Printf("Created item: %s", data)
	return &pbCrud.CreateResponse{Message: "Item created successfully"}, nil
}

// Read function for CRUD Service
func (s *CrudServer) Read(ctx context.Context, req *pbCrud.ReadRequest) (*pbCrud.ReadResponse, error) {
	token := req.GetToken()

	valid, err := validateToken(token)
	if err != nil || !valid {
		return nil, err
	}

	id := req.GetId()
	var data string
	err = database.DB.QueryRow("SELECT name FROM items WHERE id = ?", id).Scan(&data)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	log.Printf("Read item: %s", data)
	return &pbCrud.ReadResponse{Data: data}, nil
}

// Update function for CRUD Service
func (s *CrudServer) Update(ctx context.Context, req *pbCrud.UpdateRequest) (*pbCrud.UpdateResponse, error) {
	token := req.GetToken()

	valid, err := validateToken(token)
	if err != nil || !valid {
		return nil, err
	}

	id := req.GetId()
	newData := req.GetNewData()

	result, err := database.DB.Exec("UPDATE items SET name = ? WHERE id = ?", newData, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, err
	}

	log.Printf("Updated item ID %s with new data: %s", id, newData)
	return &pbCrud.UpdateResponse{Message: "Item updated successfully"}, nil
}

// Delete function for CRUD Service
func (s *CrudServer) Delete(ctx context.Context, req *pbCrud.DeleteRequest) (*pbCrud.DeleteResponse, error) {
	token := req.GetToken()

	valid, err := validateToken(token)
	if err != nil || !valid {
		return nil, err
	}

	id := req.GetId()

	result, err := database.DB.Exec("DELETE FROM items WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, err
	}

	log.Printf("Deleted item ID %s", id)
	return &pbCrud.DeleteResponse{Message: "Item deleted successfully"}, nil
}
