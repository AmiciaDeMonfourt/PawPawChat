package repository

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type Bucket struct {
	object *Object
}

func NewBucket(client *s3.S3) *Bucket {
	return &Bucket{
		object: NewObject(client),
	}
}

func (b *Bucket) Object() *Object {
	return b.object
}
