package presignedurl

import (
	"acli/mocks/api"
	"acli/pkg/factory"
	"context"
	"errors"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"reflect"
	"testing"
)

func TestPreSignCmd(t *testing.T) {
	cfg := factory.NewConfigUtil(context.Background(), "default")
	preSignCmd := NewS3PreSignCmd(cfg)
	if preSignCmd == nil {
		t.Fatalf("Expected NewS3PreSignCmd to return a valid cobra command, got nil")
	}

	if preSignCmd.ValidArgs[0] != "object id" {
		t.Fatalf("Expected ValidArgs to be [\"object id\"], got %v", preSignCmd.ValidArgs)
	}

	if preSignCmd.Use != "presign get" {
		t.Fatalf("Expected Use to be \"presign\", got %v", preSignCmd.Use)
	}
}

func TestRunPreSignUrlFunc(t *testing.T) {
	cases := []struct {
		name     string
		client   func(t *testing.T) api.MockPreSignClient
		bucket   string
		objectId string
		expected *v4.PresignedHTTPRequest
		wantsErr bool
	}{
		{
			name: "PreSignedUrl Generation Success",
			client: func(t *testing.T) api.MockPreSignClient {
				t.Helper()
				return api.MockPreSignClient(func(ctx context.Context, input *s3.GetObjectInput, optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
					// provide a dummy generated presigned url
					return &v4.PresignedHTTPRequest{URL: "https://test.com"}, nil
				})
			},
			bucket:   "default",
			objectId: "1234",
			expected: &v4.PresignedHTTPRequest{URL: "https://test.com"},
		},
		{
			name: "PreSignedUrl Generation Failure",
			client: func(t *testing.T) api.MockPreSignClient {
				t.Helper()
				return func(ctx context.Context, input *s3.GetObjectInput, optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
					return nil, errors.New("failed to generate presigned url")
				}
			},
			bucket:   "default",
			objectId: "1234",
			expected: nil,
			wantsErr: true,
		},
		{
			name: "Empty Bucket As Args",
			client: func(t *testing.T) api.MockPreSignClient {
				return func(ctx context.Context, input *s3.GetObjectInput, optFns ...func(optinos *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
					return &v4.PresignedHTTPRequest{URL: "https://example.com"}, nil
				}
			},
			objectId: "1234",
			expected: &v4.PresignedHTTPRequest{URL: "https://example.com"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runPreSignedUrl(tt.client(t), tt.bucket, tt.objectId)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("got %v, expected %v\n", got, tt.expected)
			}

			if tt.wantsErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
			}
		})
	}
}
