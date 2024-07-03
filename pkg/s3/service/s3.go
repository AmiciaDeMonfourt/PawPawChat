package service

import (
	"log"
	pb "pawpawchat/generated/proto/s3"
	"pawpawchat/pkg/s3/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	region = "ru-central1"
)

type s3Service struct {
	bucket *repository.Bucket
	pb.UnimplementedS3ServiceServer
}

func newS3Service() *s3Service {
	return &s3Service{
		bucket: repository.NewBucket(newClient()),
	}
}

func newClient() *s3.S3 {
	return s3.New(newSession())
}

func newSession() *session.Session {
	config := &aws.Config{
		Region:           aws.String(region),
		EndpointResolver: endpoints.ResolverFunc(endpointResolver),
	}

	session, err := session.NewSession(config)
	if err != nil {
		log.Fatal("failed to create a session:", err)
	}

	return session
}

func endpointResolver(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
	return endpoints.ResolvedEndpoint{
		URL: "https://storage.yandexcloud.net",
	}, nil
}
