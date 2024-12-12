package storage

import (
	"acli/internal/storage/presignedurl"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"
)

func NewCmdStorage(cfg aws.Config) *cobra.Command {
	var s3Cmd = &cobra.Command{
		Use:   "s3 <command>",
		Short: "Operations on the S3 resource",
		Long:  "Operations on the S3 resource",
		Args:  cobra.NoArgs,
	}

	s3Cmd.AddCommand(presignedurl.NewS3PreSignCmd(cfg))
	s3Cmd.AddCommand(presignedurl.NewUploadPreSignCmd(cfg))
	s3Cmd.Flags().String("bucket", "", "A valid S3 bucket to tap into")

	return s3Cmd
}
