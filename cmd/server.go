package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tsusowake/go.server/internal/server"
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
	defer func() {
		fmt.Println("End: ...server")
	}()

	ctx := context.Background()

	srv, err := server.NewServer(ctx)
	if err != nil {
		return err
	}
	// TODO impl interrup
	// defer func() error {
	// 	return srv.RedisClient.Close()
	// }()
	return srv.Start("1323")
}
