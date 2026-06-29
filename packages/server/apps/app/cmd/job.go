package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	JobCmd = &cobra.Command{
		Use:   "job",
		Short: "job",
		RunE:  runJob,
	}
)

func runJob(_ *cobra.Command, _ []string) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	if err := runJobServer(ctx); err != nil {
		return err
	}
	return nil
}

func runJobServer(ctx context.Context) error {
	return nil
}
