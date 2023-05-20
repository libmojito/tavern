package common

import (
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func OpenAIClient() *openai.Client {
	return openai.NewClient(viper.GetString("tavern.openai.apiKey"))
}
