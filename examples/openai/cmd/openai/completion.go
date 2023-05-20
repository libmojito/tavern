package openai

import (
	"context"
	"encoding/json"
	"log"

	"github.com/libmojito/tavern/examples/openai/common"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

const (
	CompletionOptModel  = "model"
	CompletionOptPrompt = "prompt"
)

func NewCompletionCmd(opts ...func(*cobra.Command)) *cobra.Command {

	cmd := NewCmdOpts(
		&cobra.Command{
			Use:   "completion",
			Short: "Call the completion API",
			Long:  `Call the completion API`,
			Run: func(cmd *cobra.Command, args []string) {
				model, _ := cmd.Flags().GetString(CompletionOptModel)
				prompt, _ := cmd.Flags().GetString(CompletionOptPrompt)

				client := common.OpenAIClient()
				rsp, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
					Model:  model,
					Prompt: prompt,
				})
				if err != nil {
					log.Fatal("error from OpenAI: ", err)
				}

				bs, err := json.MarshalIndent(rsp, "", "  ")
				if err != nil {
					log.Fatal("error marshaling response: ", err)
				}
				cmd.Println(string(bs))
			},
		},
		opts...,
	)

	cmd.Flags().String(CompletionOptModel, openai.GPT3Ada, "model to use")
	cmd.Flags().String(CompletionOptPrompt, "", "the prompt (required)")
	cmd.MarkFlagRequired("prompt")

	return cmd
}
