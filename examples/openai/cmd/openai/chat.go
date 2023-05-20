package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/libmojito/tavern/examples/openai/common"
	"github.com/libmojito/tavern/examples/openai/slice"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

const (
	ChatOptMessage = "message"
	ChatOptModel   = "model"
)

type roleMessage struct {
	role    string
	content string
}

func decodeMessage(msg string) (roleMessage, error) {
	var rm roleMessage
	xs := strings.SplitN(msg, ":", 2)
	if len(xs) != 2 {
		return rm, fmt.Errorf("message is not in the format of [role]:[msg], given: %s", msg)
	}

	return roleMessage{
		role:    xs[0],
		content: xs[1],
	}, nil
}

func NewChatCmd(opts ...func(*cobra.Command)) *cobra.Command {
	cmd := NewCmdOpts(
		&cobra.Command{
			Use:   "chat",
			Short: "Call the chat completion API",
			Long:  `Call the chat completion API`,
			Run: func(cmd *cobra.Command, args []string) {
				client := common.OpenAIClient()

				model, _ := cmd.Flags().GetString(ChatOptModel)
				messages, _ := cmd.Flags().GetStringSlice(ChatOptMessage)
				rmsgs, err := slice.MapE(messages, decodeMessage)

				if err != nil {
					log.Fatal("error in message option: ", err)
				}
				rsp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
					Model: model,
					Messages: slice.Map(rmsgs, func(rmsg roleMessage) openai.ChatCompletionMessage {
						return openai.ChatCompletionMessage{
							Role:    rmsg.role,
							Content: rmsg.content,
						}
					}),
				})
				if err != nil {
					log.Fatal("error when calling chat: ", err)
				}
				output, err := json.MarshalIndent(rsp, "", "  ")
				if err != nil {
					log.Fatal("error marshaling response: ", err)
				}
				cmd.Println(string(output))
			},
		},
		opts...,
	)

	cmd.Flags().String(ChatOptModel, openai.GPT3Dot5Turbo, "model to use")
	cmd.Flags().StringSlice(
		"message",
		[]string{},
		"the messages in the format of [system|user|assistant]:[message content]",
	)
	cmd.MarkFlagRequired("message")

	return cmd
}
