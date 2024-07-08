package user

import (
	"pawpawchat/internal/grpc"
	"pawpawchat/internal/producer"
)

type userRoutes struct {
	gRPCClient *grpc.Client
	producer   *producer.Producer
}

func NewUserRoutes(gRPCClient *grpc.Client, producer *producer.Producer) *userRoutes {
	return &userRoutes{
		gRPCClient: gRPCClient,
		producer:   producer,
	}
}
