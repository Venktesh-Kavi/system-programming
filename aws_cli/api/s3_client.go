package api

import (
	"acli/pkg"
	"acli/pkg/factory"
	"context"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type PreSignClientInterface interface {
	GetObjectInput(ctx context.Context, input *s3.GetObjectInput, optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

type PreSignClient struct {
	Client *s3.PresignClient
}

func (p PreSignClient) GetObjectInput(ctx context.Context, input *s3.GetObjectInput, optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return p.GetObjectInput(ctx, input, optFns...)
}

func NewS3Client(cfg pkg.ConfigWrapper) *s3.Client {
	return factory.NewS3Client(cfg)
}

func NewPreSignClient(cfg pkg.ConfigWrapper) PreSignClient {
	return PreSignClient{
		Client: s3.NewPresignClient(factory.NewS3Client(cfg)),
	}
}
