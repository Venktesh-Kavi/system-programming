package upload

import "github.com/spf13/cobra"

func NewS3UploadCmd() *cobra.Command {
	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "upload to s3",
		Long:  `upload a file to s3`,
	}

	uploadCmd.Flags().String("file")
	return uploadCmd
}
