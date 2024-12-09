package cli

/**
A command line utility duplicating the functionality of aws s3 cli commands. Purely done to try out Go AWS SDK interactions and the cli is done as an alternative to trigger the services without a conventional webserver.
*/

import (
	s3Wrapper "aws-go-slate/s3wrapper"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

const s3ConfigCtxKey = "s3-config"

var rootCmd = &cobra.Command{
	Use:   "gaws",
	Short: "gaws is a CLI tool for interacting with AWS resources with a go client",
	Long:  `gaws is a CLI tool for interacting with AWS resources with a go client`,
	// runs before every child command
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("dev-oblatko-s3-usr"))
		if err != nil {
			return err
		}
		s3Config := s3Wrapper.MakeS3Config(cfg)
		cmd.SetContext(context.WithValue(ctx, s3ConfigCtxKey, s3Config))
		return nil
	},
}

var getPreSignedUrl = &cobra.Command{
	Use:   "get-presigned-url",
	Short: "Get a presigned url for an s3 object",
	Long:  "Get a presigned url for an s3 object",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cc, err := retrieveConfig(cmd.Context())
		if err != nil {
			return err
		}
		bucket, key := args[0], args[1]
		preSignedUrl, err := cc.GetPreSignedUrl(ctx, bucket, key)
		if err != nil {
			return err
		}
		fmt.Println(preSignedUrl.URL)
		return nil
	},
}

var putPreSignedUrl = &cobra.Command{
	Use:   "put-presigned-url",
	Short: "Generate a Uploadable presigned url",
	Long:  "Generate a Uploadable presigned url",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucket, sfp := args[0], args[1]
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cc, err := retrieveConfig(cmd.Context())
		if err != nil {
			return err
		}
		putUrl, err := cc.PutPreSignUrl(ctx, bucket, sfp)
		handleErr(err)
		fmt.Printf("Upload PreSigned URL: %v\n", putUrl.URL)
		return nil
	},
}

func retrieveConfig(ctx context.Context) (s3Wrapper.S3Config, error) {
	s3Config := ctx.Value(s3ConfigCtxKey)
	if s3Config == nil {
		log.Fatalf("unable to initialise s3 config")
	}
	if _, ok := s3Config.(s3Wrapper.S3Config); !ok {
		return s3Wrapper.S3Config{}, fmt.Errorf("error type casting s3 config")
	}
	return s3Config.(s3Wrapper.S3Config), nil
}

var uploadFile = &cobra.Command{
	Use:   "upload-file",
	Short: "Upload a file to an s3 bucket via a pre-signed url",
	Long:  "Upload a file to an s3 bucket via a pre-signed url",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		putUrl, filePath := args[0], args[1]
		cc, err := retrieveConfig(cmd.Context())
		if err != nil {
			return err
		}
		ue := cc.UploadToBucket(putUrl, filePath)
		if ue != nil {
			log.Println("successfully uploaded file, path: {}", filePath)
		}
		return nil
	},
}

func Execute() {
	rootCmd.AddCommand(getPreSignedUrl)
	rootCmd.AddCommand(putPreSignedUrl)
	rootCmd.AddCommand(uploadFile)
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error executing command: %v", err)
		handleErr(err)
	}
}

func handleErr(err error) {
	if err == nil {
		return
	}
	_, bufErr := fmt.Fprintf(os.Stderr, "error executing command: %v", err)
	if bufErr != nil {
		fmt.Printf("unable to write error to the buffer, here is the error: %v", bufErr)
	}
	os.Exit(1)
}
