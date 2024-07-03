package client

import (
	"os"
	"pawpawchat/generated/proto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	users users.UsersServiceClient
}

func New() (*Client, error) {
	conn, err := grpc.NewClient(os.Getenv("USERS_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		users: users.NewUsersServiceClient(conn),
	}, nil
}

func (c *Client) Users() users.UsersServiceClient {
	return c.users
}
