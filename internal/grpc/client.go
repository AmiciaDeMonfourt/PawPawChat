package grpc

import (
	"log"
	"os"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/generated/proto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	users users.UsersServiceClient
	auth  auth.AuthServiceClient
}

func NewClient() (*Client, error) {
	log.Println("New gRPCClient start:")

	usersAddr := os.Getenv("USERS_ADDR")
	authAddr := os.Getenv("AUTH_ADDR")

	usersConn, err := grpc.NewClient(usersAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	authConn, err := grpc.NewClient(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
