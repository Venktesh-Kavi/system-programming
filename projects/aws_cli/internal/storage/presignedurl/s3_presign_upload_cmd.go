package presignedurl

import (
	"acli/pkg"
	"github.com/spf13/cobra"
)

func NewUploadPreSignCmd(cfg pkg.ConfigWrapper) *cobra.Command {
	uploadPreSignCmd := &cobra.Command{
		Use:       "presign put",
		Short:     "pre-sign s3 url for upload",
		Long:      `pre-sign an s3 object by passing in a path of a local file`,
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"filepath"},
	}
	uploadPreSignCmd.Flags().String("mode", "upload", "mode of operation")
	return uploadPreSignCmd
}
