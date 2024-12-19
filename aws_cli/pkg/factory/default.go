package factory

import (
	"acli/pkg"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

//Default factory apis provides all pre-made instances abstracting out creation logic from the clients.

// NewConfigUtil creates a new aws config wrapper instance
func NewConfigUtil(ctx context.Context, profile string) pkg.ConfigWrapper {
	return pkg.ConfigWrapper{
		ConfigFn: loadConfig(ctx, profile),
	}
}

func loadConfig(ctx context.Context, profile string) func() aws.Config {
	return func() aws.Config {
		cfg, _ := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
		return cfg
	}
}

func NewS3Client(cfg pkg.ConfigWrapper, opts ...func(options *s3.Options)) *s3.Client {
	log.Println("received cfg: ", cfg)
	return s3.NewFromConfig(cfg.ConfigFn(), opts...)
}
