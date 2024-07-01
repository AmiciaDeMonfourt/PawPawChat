package grpcclient

import (
	"pawpawchat/config"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/generated/proto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	users users.UsersServiceClient
	auth  auth.AuthServiceClient
}

func New() (*Client, error) {
	usersConn, err := grpc.NewClient(config.App().UsersAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	authConn, err := grpc.NewClient(config.App().AuthAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		users: users.NewUsersServiceClient(usersConn),
		auth:  auth.NewAuthServiceClient(authConn),
	}, nil
}

func (c *Client) Users() users.UsersServiceClient {
	return c.users
}

func (c *Client) Auth() auth.AuthServiceClient {
	return c.auth
}
