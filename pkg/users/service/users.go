package service

import (
	"context"
	"pawpawchat/generated/proto/users"
	db "pawpawchat/pkg/users/database"
	"pawpawchat/pkg/users/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersService struct {
	db *db.DataBase
	users.UnimplementedUsersServiceServer
}

func New(db *db.DataBase) *UsersService {
	return &UsersService{db: db}
}

func (s *UsersService) Create(ctx context.Context, req *users.CreateRequest) (*users.CreateResponse, error) {
	user, err := model.NewUser(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "incorrect user data: %v", err)
	}

	existingUser, err := s.db.User().FindByEmail(user.Email)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}

	if existingUser != nil {
		return &users.CreateResponse{
			Error: "user with this email already exists",
		}, nil
	}

	err = s.db.User().Create(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &users.CreateResponse{
		User: &users.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s *UsersService) FindByEmail(ctx context.Context, req *users.FindByEmailRequest) (*users.FindByEmailResponse, error) {
	user, err := s.db.User().FindByEmail(req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		return &users.FindByEmailResponse{
			Error: "failed find user with this email",
		}, nil
	}

	return &users.FindByEmailResponse{
		User: &users.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			HashPass: user.HashPass,
		},
	}, nil
}
