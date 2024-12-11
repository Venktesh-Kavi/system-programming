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

// Commands represents ACTIONS, Args are Things and Flags are modifiers to these actions.
// APPNAME VERB NOUN --ADJECTIVE / APPNAME COMMAND ARG --FLAGS
// hugo server --port=8080 / hugo server URL --bare

const s3ConfigCtxKey = "s3-config"

var version = "0.1v"
var rootCmd = &cobra.Command{
	Use:     "gaws",
	Version: version,
	Short:   "gaws is a CLI tool for interacting with AWS resources with a go client",
	Long:    `gaws is a CLI tool for interacting with AWS resources with a go client`,
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

var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Operations on the S3 resource",
	Long:  "Operations on the S3 resource",
	Args:  cobra.ExactArgs(1),
}

const getPreSignedUrl = "get"
const putPreSignedUrl = "put"

var s3UrlCmd = &cobra.Command{
	Use:       "url",
	Short:     "Get a preSignedUrl for an s3 object",
	Long:      "Get a preSignedUrl for an s3 object",
	ValidArgs: []string{"bucket", "objectKey/filePath"},
	Args:      cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cc, err := retrieveConfig(cmd.Context())
		s3Service := s3Wrapper.S3Service{
			Config: cc,
		}
		if err != nil {
			return err
		}
		bucket, key := args[0], args[1]
		urlOp := cmd.Flags().Lookup("preSignedUrl").Value.String()
		if urlOp == getPreSignedUrl {
			preSignedUrl, err := s3Service.GetPreSignedUrl(ctx, bucket, key)
			if err != nil {
				return err
			}
			fmt.Println(preSignedUrl.URL)
		} else if urlOp == putPreSignedUrl {
			putUrl, err := s3Service.PutPreSignUrl(ctx, bucket, key)
			handleErr(err)
			fmt.Printf("Upload PreSigned URL: %v\n", putUrl.URL)
		}
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

var s3UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file to an s3 bucket via a pre-signed url",
	Long:  "Upload a file to an s3 bucket via a pre-signed url",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		putUrl, filePath := args[0], args[1]
		cc, err := retrieveConfig(cmd.Context())
		s3Service := s3Wrapper.S3Service{
			Config: cc,
		}
		if err != nil {
			return err
		}
		ue := s3Service.UploadToBucket(putUrl, filePath)
		if ue != nil {
			log.Println("successfully uploaded file, path: {}", filePath)
		}
		return nil
	},
}

func Execute() {
	rootCmd.AddCommand(s3Cmd)
	s3UrlCmd.Flags().String("preSignedUrl", "get", "use get/put to generate get/put preSigned url")
	s3Cmd.AddCommand(s3UrlCmd)
	s3Cmd.AddCommand(s3UploadCmd)
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
