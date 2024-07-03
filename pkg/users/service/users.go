package service

import (
	"context"
	"pawpawchat/generated/proto/users"
	"pawpawchat/internal/grpcclient"
	"pawpawchat/pkg/users/database"
	"pawpawchat/pkg/users/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type usersService struct {
	grpc *grpcclient.Client
	db   *database.DataBase
	users.UnimplementedUsersServiceServer
}

func newUsersService(db *database.DataBase) *usersService {
	return &usersService{db: db}
}

func (s *usersService) Create(ctx context.Context, req *users.CreateRequest) (*users.CreateResponse, error) {
	// parse user from request
	user, err := model.NewUser(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user data: %v", err)
	}

	// does exist user with this email
	existingUser, err := s.db.User().GetByEmail(user.Email)
	// db error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}
	// user exists
	if existingUser != nil {
		return &users.CreateResponse{
			Error: "user with this email already exists",
		}, nil
	}

	// add item to database
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

func (s *usersService) CheckCredentials(ctx context.Context, req *users.CheckCredentialsRequest) (*users.CheckCredentialsResponse, error) {
	// check request
	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing password")
	}

	// get user's hash password
	hashPass, err := s.getUserHashPass(req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// can't find user with this email
	if hashPass == "" {
		return &users.CheckCredentialsResponse{Error: "not found user with this email"}, nil
	}

	// create user model
	user := model.User{Email: req.Email, HashPass: hashPass}
	// check password
	if err := user.ValidatePassword(req.Password); err != nil {
		return &users.CheckCredentialsResponse{Error: "invalid password"}, nil
	}

	return &users.CheckCredentialsResponse{}, nil
}

func (s *usersService) GetByEmail(ctx context.Context, req *users.GetByEmailRequest) (*users.GetByEmailResponse, error) {
	// check request
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing email")
	}

	// searching
	user, err := s.db.User().GetByEmail(req.GetEmail())
	// internal database error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// searching completed without errors, but user not found
	if user == nil {
		return &users.GetByEmailResponse{Error: "not found user with this email"}, nil
	}

	return &users.GetByEmailResponse{
		User: &users.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s *usersService) GetById(ctx context.Context, req *users.GetByIdRequest) (*users.GetByIdResponse, error) {
	user, err := s.db.User().GetById(req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		return &users.GetByIdResponse{Error: "not found user with this id"}, nil
	}

	return &users.GetByIdResponse{
		User: &users.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s *usersService) GetByUsername(ctx context.Context, req *users.GetByUsernameRequest) (*users.GetByUsernameResponse, error) {
	user, err := s.db.User().GetByUsername(req.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	if user == nil {
		return &users.GetByUsernameResponse{Error: "not found user with this username"}, nil
	}

	return &users.GetByUsernameResponse{
		User: &users.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s *usersService) getUserHashPass(email string) (string, error) {
	return s.db.User().GetHashPass(email)
}
