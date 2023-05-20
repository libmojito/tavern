package openai

import (
	"io"

	"github.com/spf13/cobra"
)

func NewCmd(opts ...func(*cobra.Command)) *cobra.Command {
	cmd := NewCmdOpts(
		&cobra.Command{
			Use:   "openai",
			Short: "The commands that it serves",
			Long:  `The commands that it serves`,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
			},
		},
		opts...,
	)

	cmd.AddCommand(NewModelsCmd(opts...))
	cmd.AddCommand(NewChatCmd(opts...))
	cmd.AddCommand(NewCompletionCmd(opts...))

	return cmd

}

func NewCmdOpts(cmd *cobra.Command, opts ...func(*cobra.Command)) *cobra.Command {
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
