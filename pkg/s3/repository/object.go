package repository

import (
	"bytes"
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Object struct {
	client *s3.S3
}

func NewObject(client *s3.S3) *Object {
	return &Object{
		client: client,
	}
}

func (o *Object) GetDownloadURL(key, bucket string) (string, error) {
	// bucket and key info
	objectInfo := &s3.GetObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)}

	// send request
	req, _ := o.client.GetObjectRequest(objectInfo)

	// sign the urlon five minutes
	url, err := req.Presign(5 * time.Minute)

	if err != nil {
		return "", err
	}

	return url, nil
}

func (o *Object) GetUploadURL(ctx context.Context, key, bucket string) (string, error) {
	getObjectInfo := &s3.PutObjectInput{Key: aws.String(key), Bucket: aws.String(bucket)}
	req, _ := o.client.PutObjectRequest(getObjectInfo)

	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (o *Object) Upload(ctx context.Context, key, bucket string, media []byte) error {
	// object data
	putObjectInfo := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(media),
	}

	// send put request
	if _, err := o.client.PutObjectWithContext(ctx, putObjectInfo); err != nil {
		return err
	}

	// sended object info
	putObjectHead := s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	// wait until the object is created
	if err := o.client.WaitUntilObjectExistsWithContext(ctx, &putObjectHead); err != nil {
		return err
	}

	return nil
}
