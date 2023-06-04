package chatgpt

import (
	"context"

	"github.com/VMETA3/sui-hackathon/utils"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var CG *ChatGPT

type GPTModule string

type ChatGPT struct {
	Client *gogpt.Client
}

const (
	TEXT_DAVINCI GPTModule = "text-davinci-003"
	TURBO        GPTModule = "gpt-3.5-turbo"
	WHISPER_ONE  GPTModule = "whisper-1"
)

func InitGPT() *ChatGPT {
	if CG == nil {
		CG = &ChatGPT{
			Client: gogpt.NewClient(utils.GetENV("OPENAI_API_KEY")),
		}
	}
	return CG
}

func (CG *ChatGPT) Completion(Model GPTModule, Prompt string, Stop []string) (gogpt.CompletionResponse, error) {
	return CG.Client.CreateCompletion(context.Background(), gogpt.CompletionRequest{
		Model:       string(Model),
		MaxTokens:   500,
		Temperature: 0.4,
		Prompt:      Prompt,
		Stop:        Stop,
	})
}
