package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/tsusowake/go.server/internal/server"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "server",
		RunE:  runServer,
	}
)

func runServer(_ *cobra.Command, _ []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := server.Run(ctx); err != nil {
		return err
	}
	return nil
}
