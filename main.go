package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tsusowake/go.server/cmd"
)

func main() {
	c := &cobra.Command{Use: ""}
	cmd.AddCommands(c)
	switch err := c.Execute(); {
	case err == nil:
		os.Exit(0)
	case errors.Is(err, context.Canceled):
		fmt.Println("canceled")
		os.Exit(0)
	default:
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
