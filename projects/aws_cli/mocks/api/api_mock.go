package api

import (
	"context"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type MockPreSignClient func(ctx context.Context, input *s3.GetObjectInput, optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)

func (m MockPreSignClient) GetObjectInput(ctx context.Context, input *s3.GetObjectInput, optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return m(ctx, input, optFns...)
}
