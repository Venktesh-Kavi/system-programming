package internal

import (
	"acli/internal/storage"
	"acli/pkg"
	"acli/pkg/factory"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

const profileFlag = "profile"

func Start() *cobra.Command {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("initialised context")
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
			log.Println("Config: ", cfg)
			return nil
		},
	}
	rootCmd.Flags().StringP(profileFlag, "p", "default", "aws profile to work with")
	rootCmd.AddCommand(storage.NewCmdStorage(cfg))
	return rootCmd
}
