package presignedurl

import (
	"acli/api"
	"acli/pkg"
	"context"
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
	"log"
	"time"
)

type PreSignedConfig struct {
	Bucket        string
	PreSignClient api.PreSignClientInterface
}

func NewS3PreSignCmd(cfg pkg.ConfigWrapper) *cobra.Command {
	preSignCmd := &cobra.Command{
		Use:       "presign get",
		Short:     "pre-sign s3 url",
		Long:      `pre-sign an s3 object by passing a amazon arn of the object`,
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"object id"},
		RunE: func(cmd *cobra.Command, args []string) error {
			bucketVal, err := cmd.Flags().GetString("bucket")
			if err != nil {
				return err
			}
			objectId := args[0]
			log.Println("received bucket: ", bucketVal)
			preSignClient := api.NewPreSignClient(cfg)
			pr, err := runPreSignedUrl(preSignClient, bucketVal, objectId)
			if err != nil {
				return err
			}
			fmt.Println(pr)
			return nil
		},
	}

	preSignCmd.Flags().String("bucket", "", "--bucket specify a bucket to operate on")

	return preSignCmd
}

func runPreSignedUrl(api api.PreSignClientInterface, bucket, objectId string) (*v4.PresignedHTTPRequest, error) {
	log.Println("running presign client api call")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pr, err := api.GetObjectInput(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &objectId},
		func(opts *s3.PresignOptions) {
			opts.Expires = 10 * time.Minute
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get object, bucket: %s, objId: %s, error: %w", bucket, objectId, err)
	}

	fmt.Println("Generated PreSigned URL: ", pr.URL)
	return pr, nil
}
