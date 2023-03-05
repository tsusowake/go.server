package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServer()
		},
	}
)

func RunServer() error {
	fmt.Println("Start: server...")
	defer fmt.Println("End: ...server")
	fmt.Println("wait for 3 sec")
	time.Sleep(3 * time.Second)
	return nil
}
