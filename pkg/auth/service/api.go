package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/generated/proto/users"
	"pawpawchat/pkg/auth/model"
	"pawpawchat/pkg/auth/validation"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(wd, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func Start() {
	authADDR := os.Getenv("AUTH_ADDR")

	authService, err := new()
	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(gRPCServer, authService)

	l, err := net.Listen("tcp", authADDR)
	if err != nil {
		log.Fatal(err)
	}

	if err := authService.KafkaConsumer.Subscribe([]string{"users"}); err != nil {
		log.Fatal(err)
	}

	authService.KafkaConsumer.Consume()
	go authService.proccessMessage()

	log.Printf("auth service server start at %s", authADDR)
	if err = gRPCServer.Serve(l); err != nil {
		log.Fatal(err)
	}

}

// depricated
func (s *AuthService) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	// Validate request
	if err := validation.SignUpRequest(req); err != nil {
		return &auth.SignUpResponse{Error: err.Error()}, nil
	}

	// Parse to model
	authInfo := model.NewAuthInfo(req)
	if authInfo == nil {
		return &auth.SignUpResponse{Error: fmt.Sprintf("bad request: %v", req)}, nil
	}

	// Insert into database
	if err := s.db.AuthInfo().Create(context.TODO(), authInfo); err != nil {
		return &auth.SignUpResponse{Error: fmt.Sprintf("internal error when inserting a record: %v", err.Error())}, nil
	}

	usersRepsonse, err := s.client.Users().Create(ctx, &users.CreateRequest{
		FirstName:  req.GetFirstName(),
		SecondName: req.GetSecondName(),
		UserID:     authInfo.UserID,
	})

	if err != nil {
		return &auth.SignUpResponse{Error: err.Error()}, nil
	}

	if usersRepsonse == nil || usersRepsonse.GetUser() == nil {
		return &auth.SignUpResponse{Error: fmt.Sprintf("user response or user == nil: %v", usersRepsonse)}, nil
	}

	if usersRepsonse.GetError() != "" {
		return &auth.SignUpResponse{Error: usersRepsonse.GetError()}, nil
	}

	return &auth.SignUpResponse{}, nil
}
