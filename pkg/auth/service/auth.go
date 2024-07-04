package service

import (
	"context"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/generated/proto/users"
	"pawpawchat/pkg/auth/client"
	"pawpawchat/pkg/auth/jwt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	client *client.Client
	auth.UnsafeAuthServiceServer
}

func new() (*authService, error) {
	// Creating new grpc client with connection to other microservices
	client, err := client.New()
	if err != nil {
		return nil, err
	}

	return &authService{
		client: client,
	}, nil
}

func (s *authService) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	// Parse request, passing data to 'users' microservice to create a new user
	createResp, err := s.client.Users().Create(ctx, &users.CreateRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})

	// Check internal errors
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Check is email unique
	if createResp.User == nil {
		return &auth.SignUpResponse{
			Error: createResp.GetError(),
		}, nil
	}

	// Generate token for new user
	tokenStr, err := jwt.GenerateToken(createResp.User.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return token string and users credentials
	return &auth.SignUpResponse{
		TokenStr: tokenStr,
		User: &auth.User{
			Id:    createResp.User.GetId(),
			Email: createResp.User.GetEmail(),
		},
	}, nil
}

func (s *authService) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error) {
	// Parse request and finding the user with the same credentials
	resp, err := s.client.Users().GetByEmail(ctx, &users.GetByEmailRequest{Email: req.GetEmail()})
	if err != nil {
		return &auth.SignInResponse{Error: err.Error()}, nil
	}

	// Check is user exists
	if resp.GetUser() == nil {
		return &auth.SignInResponse{Error: "not find user with this email"}, nil
	}

	// Check user credentials
	checkCredentialResp, err := s.client.Users().CheckCredentials(context.TODO(), &users.CheckCredentialsRequest{Email: req.GetEmail(), Password: req.GetPassword()})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := checkCredentialResp.GetError(); err != "" {
		return &auth.SignInResponse{Error: err}, nil
	}

	// Generate token
	tokenStr, err := jwt.GenerateToken(resp.User.GetId())
	if err != nil {
		return nil, nil
	}

	// Return token and user
	return &auth.SignInResponse{
		TokenStr: tokenStr,
		User: &auth.User{
			Id:    resp.GetUser().GetId(),
			Email: resp.GetUser().GetEmail(),
		},
	}, nil
}

func (s *authService) CheckAuth(ctx context.Context, req *auth.CheckAuthRequest) (*auth.CheckAuthResponse, error) {
	if req == nil || req.GetTokenStr() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing or empty token")
	}

	id, err := jwt.ExtractUserId(req.GetTokenStr())
	if err != nil {
		return nil, err
	}

	return &auth.CheckAuthResponse{
		Userid: id,
	}, nil
}
