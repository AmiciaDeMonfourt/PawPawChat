package service

import (
	"bytes"
	"context"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"pawpawchat/generated/proto/s3"
	pb "pawpawchat/generated/proto/s3"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(wd, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func Start() {
	s3ADDR := os.Getenv("S3_ADDR")
	if s3ADDR == "" {
		log.Fatal("missing s3 addr")
	}

	s3Service := newS3Service()
	s3GRPCServer := grpc.NewServer()

	s3.RegisterS3ServiceServer(s3GRPCServer, s3Service)
	reflection.Register(s3GRPCServer)

	listener, err := net.Listen("tcp", s3ADDR)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("s3 service server start at %s", os.Getenv("S3_ADDR"))

	if err := s3GRPCServer.Serve(listener); err != nil {
		log.Fatalf("s3 service error: %s", err.Error())
	}
}

// Upload media in one block
func (s *s3Service) SinglePartUpload(ctx context.Context, r *pb.SinglePartUploadRequest) (*pb.SinglePartUploadResponse, error) {
	if r.GetMedia() == nil || r.GetKey() == "" || r.GetBucket() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing key or bucket. Request: %v", r)
	}

	err := s.bucket.Object().Upload(ctx, r.GetKey(), r.GetBucket(), r.GetMedia())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to laod media: "+err.Error())
	}

	url, err := s.bucket.Object().GetDownloadURL(r.GetKey(), r.GetBucket())
	if err != nil {
		return nil, err
	}

	return &pb.SinglePartUploadResponse{URL: url}, nil
}

func (s *s3Service) StreamUpload(stream pb.S3Service_StreamUploadServer) error {
	var buffer bytes.Buffer
	var bucket, key string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		bucket = req.GetBucket()
		key = req.GetKey()

		_, err = buffer.Write(req.Chunk)
		if err != nil {
			return err
		}
	}

	if err := s.bucket.Object().Upload(context.TODO(), key, bucket, buffer.Bytes()); err != nil {
		return status.Error(codes.Internal, "failed to load media: "+err.Error())
	}

	return nil
}

// Creating a presigned URL for 5 minutes to upload media
func (s *s3Service) GetUploadURL(ctx context.Context, r *pb.GetUploadURLRequest) (*pb.GetUploadURLResponse, error) {
	if r.GetKey() == "" || r.GetBucket() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing key or bucket. Request: %v", r)
	}

	url, err := s.bucket.Object().GetUploadURL(ctx, r.GetKey(), r.GetBucket())
	if err != nil {
		return &pb.GetUploadURLResponse{Error: "failed to create a presigned url: " + err.Error()}, nil
	}

	return &pb.GetUploadURLResponse{URL: url}, nil
}

// Creating a presigned URL for download media by key
func (s *s3Service) GetDownloadURL(ctx context.Context, r *pb.GetDownloadRequest) (*pb.GetDownloadResponse, error) {
	if r.GetKey() == "" || r.GetBucket() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing key or bucket. Request: %v", r)
	}

	url, err := s.bucket.Object().GetDownloadURL(r.GetKey(), r.GetBucket())
	if err != nil {
		return &pb.GetDownloadResponse{Error: "failed to create a presigned url: " + err.Error()}, nil
	}

	return &pb.GetDownloadResponse{URL: url}, nil
}
