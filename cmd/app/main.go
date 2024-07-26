package main

import (
	"fmt"
	"web3-practice/internal/server"

	"github.com/spf13/cobra"
)

var (
	name = "app"
	cmd  = &cobra.Command{Use: name}
)

func startCmd(srv *server.Server) *cobra.Command {
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

func initCmd(cfg *server.Config) error {
	s, err := server.NewServer(cfg)
	if err != nil {
		return err
	}
	cmd.AddCommand(startCmd(s))
	return cmd.Execute()
}

func main() {
	c, err := server.InitConfig(name)
	if err != nil {
		panic(err)
	}
	if err = initCmd(c); err != nil {
		panic(err)
	}
}
