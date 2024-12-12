package presignedurl

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

// I need options so that I can store and use it later
type S3PreSignOption struct {
	ObjectId string
	S3Client func(cfg aws.Config) *s3.Client
}

func NewS3PreSignCmd(cfg aws.Config) *cobra.Command {
	preSignCmd := &cobra.Command{
		Use:   "presign",
		Short: "pre-sign s3 url",
		Long:  `pre-sign an s3 object by passing a amazon arn of the object`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			objectId := args[0]
			return runPreSignedUrl(objectId)
		},
	}

	return preSignCmd
}
func runPreSignedUrl(objectId string) error {

}
