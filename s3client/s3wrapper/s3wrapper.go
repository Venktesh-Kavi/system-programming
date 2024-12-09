package s3wrapper

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type S3Config struct {
	cfg    aws.Config
	client *s3.Client
}

func MakeS3Config(cfg aws.Config) S3Config {
	return S3Config{
		cfg:    cfg,
		client: s3.NewFromConfig(cfg),
	}
}

// GetPreSignedUrl returns a preSigned url for an s3 object
func (s3Client S3Config) GetPreSignedUrl(ctx context.Context, bucket string, objId string) (*v4.PresignedHTTPRequest, error) {
	if s3Client.client == nil {
		log.Fatalf("s3 client is nil")
	}
	preSignedClient := s3.NewPresignClient(s3Client.client)
	pr, err := preSignedClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &objId},
		func(opts *s3.PresignOptions) {
			opts.Expires = 10 * time.Minute
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get object, bucket: %s, objId: %s, error: %w", bucket, objId, err)
	}
	return pr, nil
}

func (s3Client S3Config) PutPreSignUrl(ctx context.Context, bucket string, sfp string) (*v4.PresignedHTTPRequest, error) {
	if s3Client.client == nil {
		log.Fatalf("s3 client is nil")
	}
	filename := filepath.Base(sfp)
	fileExt := filepath.Ext(sfp)
	ct := http.DetectContentType([]byte(fileExt))
	fmt.Printf("generating a uploadable presigned url, bucket: %s, file: %s, ext: %s, content type: %s\n", bucket, sfp, fileExt, ct)
	// Note: /dev/null is black hole file were any output written is discarded. `os.DevNull` is a constant here. Use ioutil.Discard for writing to a black hole.
	discardReader, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	preSignedOp, err := s3.NewPresignClient(s3Client.client, func(opts *s3.PresignOptions) {
		opts.Expires = 10 * time.Minute
	}).PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &filename,
		Body:        discardReader,
		ContentType: &ct,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to put object, %w", err)
	}
	return preSignedOp, nil
}

func (s3Client S3Config) UploadToBucket(putUrl string, sfp string, metaHeaders ...map[string]string) error {
	log.Printf("opening file: %s\n", sfp)
	file, err := os.Open(sfp)
	if err != nil {
		return fmt.Errorf("unable to open file, %v", err)
	}

	// defer functions are typically used for cleaning up resources, they are executed after the surrounding function.
	// here the closure accesses the file variable from the outside scope and closes on it.
	// this is a good practice to close the file after the function is done with it.
	// here we cannot throw the error using fmt.Errorf as defer fn's are typically used for clearing up resources and not for returning values.
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file, %v", err)
		}
	}(file)

	// create a buffer. Buffer is variable sized array of bytes with reader and writes.
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	log.Printf("uploading file: %s\n", sfp)
	req, err := http.NewRequest(http.MethodPut, putUrl, buf)
	if err != nil {
		return fmt.Errorf("error uploading file to put url, %v", err)
	}
	req.Header.Set("Content-Type", http.DetectContentType([]byte(filepath.Ext(sfp))))

	client := http.Client{}

	// there is already a shorthand declaration of err in #9. Unless and until there is a new variable assignment, err cannot be reused.
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("error uploading file to put url, %v", err)
	}

	if resp.StatusCode >= 400 {
		op, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error uploading file to put url, %v", string(op))
	}
	log.Printf("recieved response for file upload: %s, %v\n", sfp, resp.Status)
	return nil
}
