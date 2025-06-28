package upload

import (
	"acli/pkg"
	"github.com/spf13/cobra"
)

func NewS3UploadCmd(cfg pkg.ConfigWrapper) *cobra.Command {
	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "upload to s3",
		Long:  `upload a file to s3`,
	}

	return uploadCmd
}
