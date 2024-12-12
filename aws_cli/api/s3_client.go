package api

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func MakeS3Client(cfg aws.Config) *s3.Client {
	client := s3.NewFromConfig(cfg)
	return client
}
