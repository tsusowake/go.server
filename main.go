package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tsusowake/go.server/cmd"
)

func main() {
	c := &cobra.Command{Use: ""}
	cmd.AddCommands(c)
	switch err := c.Execute(); err {
	case nil:
		os.Exit(0)
	case context.Canceled:
		fmt.Println("canceld")
		os.Exit(0)
	default:
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
