package cmd

import (
	"io"
	"strings"

	"github.com/spf13/cobra"
)

func NewHelloCmd(opts ...func(*cobra.Command)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "An example hello command that just echos back given args",
		Long:  `An example hello command that just echos back given args`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("hello %s\n", strings.Join(args, " "))
		},
		DisableFlagParsing: true,
	}

	for _, o := range opts {
		o(cmd)
	}

	return cmd
}

func WithOut(w io.Writer) func(*cobra.Command) {
	return func(cmd *cobra.Command) {
		cmd.SetOut(w)
	}
}

func WithErr(w io.Writer) func(*cobra.Command) {
	return func(cmd *cobra.Command) {
		cmd.SetErr(w)
	}
}
