package main

import (
	"fmt"
	"web3-practice/internal/config"
	"web3-practice/internal/server"

	"github.com/spf13/cobra"
)

const name = "gateway"

var cmd = &cobra.Command{Use: name}

func startCmd(srv *server.GrpcServer) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Starts the server.",
		Long:  `Starts a server that interacts with the network.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return fmt.Errorf("trailing args detected")
			}
			// Parsing of the command line is done so silence cmd usage
			cmd.SilenceUsage = true
			return srv.Start(args)
		},
	}
}

func initCmd(cfg *config.Config) error {
	srv, err := server.NewGrpcServer(cfg)
	if err != nil {
		return err
	}
	cmd.AddCommand(startCmd(srv))
	return cmd.Execute()
}

func main() {
	cfg, err := config.InitConfig(name)
	if err != nil {
		panic(err)
	}
	if err = initCmd(cfg); err != nil {
		panic(err)
	}
}
