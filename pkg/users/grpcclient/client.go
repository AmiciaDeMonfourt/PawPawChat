package grpcclient

import (
	"log"
	"os"
	"pawpawchat/generated/proto/s3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	s3 s3.S3ServiceClient
}

func New() *GRPCClient {
	s3ADDR := os.Getenv("")
	if s3ADDR == "" {
		log.Fatal("missing s3 service addr")
	}

	s3Connection, err := grpc.NewClient(s3ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create s3 client: %s", err.Error())
	}

	s3Client := s3.NewS3ServiceClient(s3Connection)

	return &GRPCClient{
		s3: s3Client,
	}
}
