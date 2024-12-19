package storage

import (
	"acli/internal/storage/presignedurl"
	"acli/internal/storage/upload"
	"acli/pkg"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdStorage(cfg pkg.ConfigWrapper) *cobra.Command {
	log.Println("started s3 command")
	var s3Cmd = &cobra.Command{
		Use:   "s3 <command>",
		Short: "Operations on the S3 resource",
		Long:  "Operations on the S3 resource",
		Args:  cobra.NoArgs,
	}
	s3Cmd.AddCommand(presignedurl.NewS3PreSignCmd(cfg))
	s3Cmd.AddCommand(presignedurl.NewUploadPreSignCmd(cfg))
	s3Cmd.AddCommand(upload.NewS3UploadCmd(cfg))
	return s3Cmd
}
