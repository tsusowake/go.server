package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tsusowake/go.server/apps/app/cmd"
)

var rootCmd = &cobra.Command{
	Use: "",
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
func run() error {
	registerCommands()

	if err := rootCmd.Execute(); err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Println("canceled")
			return nil
		}
		return err
	}
	return nil
}

func registerCommands() {
	rootCmd.AddCommand(cmd.ServerCmd)
	rootCmd.AddCommand(cmd.JobCmd)
}
