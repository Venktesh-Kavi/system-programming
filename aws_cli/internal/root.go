package internal

import (
	ss "acli/internal/storage"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
)

const profileFlag = "profile"

func Main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const version = "v0.1"
	var cfg aws.Config
	rootCmd := &cobra.Command{
		Use:     "gaws",
		Short:   "aws interactions written with go",
		Long:    "aws interactions written with go",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			fv := cmd.Flags().Lookup(profileFlag).Value.String()
			cfg, _ = config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(fv))
			return nil
		},
	}
	rootCmd.Flags().StringP(profileFlag, "p", "default", "aws profile to work with")

	rootCmd.AddCommand(ss.NewCmdStorage(cfg))
}
