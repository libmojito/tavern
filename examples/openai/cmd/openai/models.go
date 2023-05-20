package openai

import (
	"context"
	"encoding/json"
	"log"

	"github.com/libmojito/tavern/examples/openai/common"
	"github.com/libmojito/tavern/examples/openai/slice"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

const (
	ModelsOptID      = "id"
	ModelsOptVerbose = "verbose"
)

func NewModelsCmd(opts ...func(*cobra.Command)) *cobra.Command {
	cmd := NewCmdOpts(
		&cobra.Command{
			Use:   "models",
			Short: "Command to interact with open ai models",
			Long:  `Command to interact with open ai models`,
			Run: func(cmd *cobra.Command, args []string) {
				client := common.OpenAIClient()
				models, err := client.ListModels(context.Background())
				if err != nil {
					cmd.PrintErr("err:", err)
				}

				modelID, err := cmd.Flags().GetString(ModelsOptID)
				if err != nil {
					cmd.PrintErr("err:", err)
				}
				verbose, err := cmd.Flags().GetBool(ModelsOptVerbose)
				if err != nil {
					cmd.PrintErr("err:", err)
				}

				// listing mode:
				if modelID == "" {
					for _, m := range models.Models {
						if !verbose {
							cmd.Printf("- %s\n", m.ID)
						} else {
							output, err := json.MarshalIndent(m, "", "  ")
							if err != nil {
								log.Fatal("error when mashaling model")
							}
							cmd.Println(string(output))
						}
					}
					return
				}

				// detail mode
				m, found := slice.Find(models.Models, func(m openai.Model) bool {
					if m.ID == modelID {
						return true
					}
					return false
				})
				if found {
					output, err := json.MarshalIndent(m, "", "  ")
					if err != nil {
						log.Fatal("error when mashaling model")
					}
					cmd.Println(string(output))
				} else {
					cmd.Println("model does not exist")
				}
			},
		},
		opts...,
	)

	cmd.Flags().String(ModelsOptID, "", "Model ID")
	cmd.Flags().Bool(ModelsOptVerbose, false, "Show more details when listing")

	return cmd
}
