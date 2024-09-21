package server

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"time"

	pbAuth "github.com/rohim/auth-service/proto/protoc"

	"github.com/dgrijalva/jwt-go"
	"github.com/rohim/auth-service/database"
)

type AuthServer struct {
	pbAuth.UnimplementedAuthServiceServer
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Helper function to hash password
func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func (s *AuthServer) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error) {
	username, password := req.GetUsername(), req.GetPassword()

	var storedHash string
	err := database.DB.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&storedHash)
	if err == sql.ErrNoRows {
		return &pbAuth.LoginResponse{Message: "User not found"}, nil
	} else if err != nil {
		return nil, err
	}

	// Validate password
	if storedHash != hashPassword(password) {
		return &pbAuth.LoginResponse{Message: "Invalid password"}, nil
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &pbAuth.LoginResponse{
		Token:   tokenString,
		Message: "Login successful",
	}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *pbAuth.ValidateTokenRequest) (*pbAuth.ValidateTokenResponse, error) {
	tokenStr := req.GetToken()
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return &pbAuth.ValidateTokenResponse{
			Valid:   false,
			Message: "Invalid token",
		}, nil
	}

	return &pbAuth.ValidateTokenResponse{
		Valid:   true,
		Message: "Token is valid",
	}, nil
}
