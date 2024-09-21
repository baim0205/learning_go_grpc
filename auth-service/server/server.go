package server

import (
	"context"
	"time"

	pbAuth "github.com/rohim/auth-service/proto/protoc"

	"github.com/dgrijalva/jwt-go"
)

type AuthServer struct {
	pbAuth.UnimplementedAuthServiceServer
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *AuthServer) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error) {
	username, password := req.GetUsername(), req.GetPassword()

	// Simple check for username/password (you may want to check from DB)
	if username == "admin" && password == "password" {
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

	return &pbAuth.LoginResponse{
		Message: "Invalid username or password",
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
