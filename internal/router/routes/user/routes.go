package user

import "pawpawchat/internal/grpcclient"

type userRoutes struct {
	gRPCClient *grpcclient.Client
}

func NewUserRoutes(gRPCClient *grpcclient.Client) *userRoutes {
	return &userRoutes{
		gRPCClient: gRPCClient,
	}
}
