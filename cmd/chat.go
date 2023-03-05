package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "chat",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunChatServer()
	},
}

func RunChatServer() error {
	fmt.Println("Start: chat-server")
	defer fmt.Println("End: chat-server")
	fmt.Println("wait for 3 sec")
	time.Sleep(3 * time.Second)
	return nil
}
