package storage

import (
	"acli/internal/storage/presignedurl"
	"acli/pkg"
	"acli/pkg/factory"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

type S3Config struct {
	S3Client      func() *s3.Client
	PreSignClient func() *s3.PresignClient
}

func NewCmdStorage(cfg pkg.ConfigWrapper) *cobra.Command {
	var s3Cmd = &cobra.Command{
		Use:   "s3 <command>",
		Short: "Operations on the S3 resource",
		Long:  "Operations on the S3 resource",
		Args:  cobra.NoArgs,
	}

	s3Cfg := S3Config{
		S3Client:      s3ClientFunc(cfg),
		PreSignClient: preSignClientFUnc(cfg),
	}
	s3Cmd.AddCommand(presignedurl.NewS3PreSignCmd(s3Cfg))
	s3Cmd.AddCommand(presignedurl.NewUploadPreSignCmd(s3Cfg))
	s3Cmd.Flags().String("bucket", "", "A valid S3 bucket to tap into")
	return s3Cmd
}

func s3ClientFunc(cfg pkg.ConfigWrapper) func() *s3.Client {
	return func() *s3.Client {
		return factory.NewS3Client(cfg)
	}
}

func preSignClientFUnc(cfg pkg.ConfigWrapper) func() *s3.PresignClient {
	return func() *s3.PresignClient {
		return factory.NewPreSignClient(cfg)
	}
}
