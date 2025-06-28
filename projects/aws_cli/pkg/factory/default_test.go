package factory

import (
	"context"
	"os"
	"testing"
)

// don't need to test whether the aws sdk loads the configuration.
func TestNewConfigWrapperInstance(t *testing.T) {
	ctx := context.Background()
	profile := "default"
	config := NewConfigUtil(ctx, profile)
	awsCfg := config.ConfigFn()
	crd, _ := awsCfg.Credentials.Retrieve(ctx)
	if crd.AccessKeyID != os.Getenv("AWS_ACCESS_KEY") {
		t.Errorf("Expected %s, but got %s", os.Getenv("AWS_ACCESS_KEY"), crd.AccessKeyID)
	}
	if crd.SecretAccessKey != os.Getenv("AWS_SECRET") {
		t.Errorf("Expected %s, but got %s", os.Getenv("AWS_SECRET"), crd.SecretAccessKey)
	}

	if awsCfg.Region != "ap-south-1" {
		t.Errorf("Expected %s, but got %s", "ap-south-1", awsCfg.Region)
	}
}
