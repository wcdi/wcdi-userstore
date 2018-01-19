package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wcdi/wcdi-userstore/server"
)

func init() {
	RootCmd.AddCommand(newServerCommand())
}

func newServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newStartServerCmd(),
	)

	return cmd
}

func newStartServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "server start !",
		RunE:  runStartServerCmd,
	}

	flags := cmd.Flags()
	flags.StringP("host", "", "0.0.0.0", "hostname")
	flags.IntP("port", "", 8080, "port number")

	return cmd
}

func runStartServerCmd(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	host, _ := flags.GetString("host")
	port, _ := flags.GetInt("port")

	return server.StartServer(host, port)
}
