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

	existingUser, err := s.db.User().GetByEmail(user.Email)

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

func (s *UsersService) GetByEmail(ctx context.Context, req *users.GetByEmailRequest) (*users.GetByEmailResponse, error) {
	user, err := s.db.User().GetByEmail(req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		return &users.GetByEmailResponse{
			Error: "failed Get user with this email",
		}, nil
	}

	return &users.GetByEmailResponse{
		User: &users.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s *UsersService) GetById(ctx context.Context, req *users.GetByIdRequest) (*users.GetByIdResponse, error) {
	user, err := s.db.User().GetById(req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	if user == nil {
		return &users.GetByIdResponse{
			Error: "not found",
		}, nil
	}

	return &users.GetByIdResponse{
		User: &users.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s *UsersService) GetByUsername(ctx context.Context, req *users.GetByUsernameRequest) (*users.GetByUsernameResponse, error) {
	user, err := s.db.User().GetByUsername(req.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	if user == nil {
		return &users.GetByUsernameResponse{
			Error: "not found",
		}, nil
	}

	return &users.GetByUsernameResponse{
		User: &users.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}
