package service

import (
	"context"
	"encoding/json"
	"log"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/pkg/auth/client"
	"pawpawchat/pkg/auth/consumer"
	"pawpawchat/pkg/auth/database"
	"pawpawchat/pkg/auth/model"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type AuthService struct {
	client        *client.Client
	db            *database.DataBase
	KafkaConsumer *consumer.Consumer
	auth.UnsafeAuthServiceServer
}

// REFACTIOR
func new() (*AuthService, error) {
	// Creating new grpc client with connection to other microservices
	client, err := client.New()
	if err != nil {
		return nil, err
	}

	return &AuthService{
		client:        client,
		db:            database.New(),
		KafkaConsumer: consumer.New("kafka:9092", "user_signup"),
	}, nil
}

func (s *AuthService) proccessMessage() {
	log.Printf("message processing is started")
	for msg := range s.KafkaConsumer.MsgChannel {
		go func(msg *kafka.Message) {
			log.Printf("message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			var signUpReq model.AuthInfo
			if err := json.Unmarshal(msg.Value, &signUpReq); err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				return
			}

			if err := s.UserSignUp(context.TODO(), &signUpReq); err != nil {
				log.Printf("failed to process sign up request: %v", err)
				return
			}

		}(msg)
	}
}

func (s *AuthService) UserSignUp(ctx context.Context, req *model.AuthInfo) error {
	return s.db.AuthInfo().Create(ctx, req)
}

// createResp, err := s.client.Users().Create(ctx, &users.CreateRequest{
// 	Email:    req.GetEmail(),
// 	Password: req.GetPassword(),
// })

// // Check internal errors
// if err != nil {
// 	return nil, status.Error(codes.Internal, err.Error())
// }

// // Check is email unique
// if createResp.User == nil {
// 	return &auth.SignUpResponse{
// 		Error: createResp.GetError(),
// 	}, nil
// }

// // Generate token for new user
// tokenStr, err := jwt.GenerateToken(createResp.User.GetId())
// if err != nil {
// 	return nil, status.Error(codes.Internal, err.Error())
// }

// // Return token string and users credentials
// return &auth.SignUpResponse{
// 	TokenStr: tokenStr,
// 	User: &auth.User{
// 		Id:    createResp.User.GetId(),
// 		Email: createResp.User.GetEmail(),
// 	},
// }, nil

func (s *AuthService) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error) {
	return nil, nil
}

// 	// Parse request and finding the user with the same credentials
// 	resp, err := s.client.Users().GetByEmail(ctx, &users.GetByEmailRequest{Email: req.GetEmail()})
// 	if err != nil {
// 		return &auth.SignInResponse{Error: err.Error()}, nil
// 	}

// 	// Check is user exists
// 	if resp.GetUser() == nil {
// 		return &auth.SignInResponse{Error: "not find user with this email"}, nil
// 	}

// 	// Check user credentials
// 	checkCredentialResp, err := s.client.Users().CheckCredentials(context.TODO(), &users.CheckCredentialsRequest{Email: req.GetEmail(), Password: req.GetPassword()})
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	if err := checkCredentialResp.GetError(); err != "" {
// 		return &auth.SignInResponse{Error: err}, nil
// 	}

// 	// Generate token
// 	tokenStr, err := jwt.GenerateToken(resp.User.GetId())
// 	if err != nil {
// 		return nil, nil
// 	}

// 	// Return token and user
// 	return &auth.SignInResponse{
// 		TokenStr: tokenStr,
// 		User: &auth.User{
// 			Id:    resp.GetUser().GetId(),
// 			Email: resp.GetUser().GetEmail(),
// 		},
// 	}, nil
// }

func (s *AuthService) CheckAuth(ctx context.Context, req *auth.CheckAuthRequest) (*auth.CheckAuthResponse, error) {
	return nil, nil
}

// 	if req == nil || req.GetTokenStr() == "" {
// 		return nil, status.Error(codes.InvalidArgument, "missing or empty token")
// 	}

// 	id, err := jwt.ExtractUserId(req.GetTokenStr())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &auth.CheckAuthResponse{
// 		Userid: id,
// 	}, nil
// }
