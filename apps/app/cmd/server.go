package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/tsusowake/go.server/internal/server"
)

var (
	ServerCmd = &cobra.Command{
		Use:   "server",
		Short: "server",
		RunE:  runServer,
	}
)

func runServer(_ *cobra.Command, _ []string) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	if err := server.Run(ctx); err != nil {
		return err
	}
	return nil
}
