package cmd

import "github.com/spf13/cobra"

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command-server",
		Short: "The root command of this command server",
		Long:  `The root command of this command server`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewServeCmd())
	cmd.AddCommand(NewHelloCmd())

	return cmd
}
