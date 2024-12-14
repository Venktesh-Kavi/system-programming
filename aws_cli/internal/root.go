package internal

import (
	ss "acli/internal/storage"
	"acli/pkg"
	"acli/pkg/factory"
	"context"
	"github.com/spf13/cobra"
)

const profileFlag = "profile"

func Main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const version = "v0.1"
	var cfg pkg.ConfigWrapper // passing config to downstream commands makes subcommands un-mockable. Wrapping it in a factory.
	rootCmd := &cobra.Command{
		Use:     "gaws",
		Short:   "aws interactions written with go",
		Long:    "aws interactions written with go",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			fv := cmd.Flags().Lookup(profileFlag).Value.String()
			cfg = factory.NewConfigUtil(ctx, fv)
			return nil
		},
	}
	rootCmd.Flags().StringP(profileFlag, "p", "default", "aws profile to work with")
	rootCmd.AddCommand(ss.NewCmdStorage(cfg))
}
