package gpt

import (
	"context"
	chatgpt "github.com/VMETA3/sui-hackathon/utils/chatGPT"
	"time"
)

func (C *GPT) Single(Prompt string) (answer string, err error) {
	Stop := []string{"\n\nAI Assistant:"}
	resp, err := C.TCG.Completion(chatgpt.TEXT_DAVINCI, Prompt, Stop)
	if err != nil {
		return
	}
	answer = resp.Choices[0].Text
	return
}

func (C *GPT) Continuous(SESSION_ID, P string) (answer string, err error) {
	ctx := context.Background()

	Key := "Chat:" + SESSION_ID

	Prompt, err := C.Redis.Get(ctx, Key).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			return
		}
	}

	Prompt += "\n" + SESSION_ID + ": " + P + "\n"
	Stop := []string{"\n" + SESSION_ID + ":", "\n\nAI Assistant:"}

	resp, err := C.TCG.Completion(chatgpt.TEXT_DAVINCI, Prompt, Stop)

	if err != nil {
		return
	}
	answer = resp.Choices[0].Text
	Prompt += resp.Choices[0].Text

	err = C.Redis.Set(ctx, Key, Prompt, time.Minute*30).Err()
	return
}
