package pkg

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

type ConfigWrapper struct {
	ConfigFn func() aws.Config
}
