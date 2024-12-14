package presignedurl

import (
	ss "acli/internal/storage"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// I need options so that I can store and use it later

func NewS3PreSignCmd(s3Cfg ss.S3Config) *cobra.Command {
	preSignCmd := &cobra.Command{
		Use:   "presign",
		Short: "pre-sign s3 url",
		Long:  `pre-sign an s3 object by passing a amazon arn of the object`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			objectId := args[0]
			return runPreSignedUrl(s3Cfg, objectId)
		},
	}
	return preSignCmd
}

func runPreSignedUrl(s3Cfg ss.S3Config, objectId string) error {
	if s3Cfg.S3Client == nil {
		log.Fatalf("s3 client is nil")
	}
	preSignedClient := s3.NewPresignClient(s3Cfg.S3Client())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pr, err := preSignedClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &objectId},
		func(opts *s3.PresignOptions) {
			opts.Expires = 10 * time.Minute
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get object, bucket: %s, objId: %s, error: %w", bucket, objId, err)
	}
	return pr, nil
}
