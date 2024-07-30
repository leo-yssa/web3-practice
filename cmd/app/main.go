package main

import (
	"fmt"
	_ "web3-practice/docs"
	"web3-practice/internal/config"
	"web3-practice/internal/server"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{Use: server.Name}
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

func initCmd(cfg *config.Config) error {
	srv, err := server.NewServer(cfg)
	if err != nil {
		return err
	}
	cmd.AddCommand(startCmd(srv))
	return cmd.Execute()
}

// @title Advertise Platform
// @version 1.0.0
// @description Advertise Platform
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @consumes application/json multipart/form-data
// @produces application/json
// @termsOfService http://swagger.io/terms/

// @contact.name Leo-yssa

// @BasePath /
func main() {
	cfg, err := config.InitConfig(server.Name)
	if err != nil {
		panic(err)
	}
	if err = initCmd(cfg); err != nil {
		panic(err)
	}
}
