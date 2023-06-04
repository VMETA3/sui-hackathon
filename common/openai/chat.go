package openai

import (
	"context"
	"encoding/json"
	"github.com/VMETA3/sui-hackathon/utils/openai"
	"time"
)

func (OA *OPENAI) Single(Prompt, definition string) (answer string, err error) {
	var chat []*openai.TurboMessage
	if definition == "" {
		definition = "You are a humorous and gentle assistant"
	}
	chat = append(chat, &openai.TurboMessage{
		Role:    openai.SYSTEM,
		Content: definition,
	})
	chat = append(chat, &openai.TurboMessage{
		Role:    openai.USER,
		Content: Prompt,
	})
	resp, err := OA.Client.Turbo().Chat(chat)
	if err != nil {
		return
	}
	answer = resp.Content
	return
}

func (OA *OPENAI) Continuous(SESSION_ID, P, definition string) (answer string, err error) {
	var chat []*openai.TurboMessage
	ctx := context.Background()
	Key := "Chat:" + SESSION_ID
	Prompts, err := OA.Redis.Get(ctx, Key).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			return
		}
	}
	if Prompts != "" {
		json.Unmarshal([]byte(Prompts), &chat)
	}
	l := len(chat)
	if l == 0 {
		if definition == "" {
			definition = "You are a humorous and gentle assistant"
		}
		chat = append(chat, &openai.TurboMessage{
			Role:    openai.SYSTEM,
			Content: definition,
		})
	}
	chat = append(chat, &openai.TurboMessage{
		Role:    openai.USER,
		Content: P,
	})

	resp, err := OA.Client.Turbo().Chat(chat)
	if err != nil {
		return
	}
	answer = resp.Content
	chat = append(chat, &openai.TurboMessage{
		Role:    resp.Role,
		Content: resp.Content,
	})

	np, err := json.Marshal(chat)
	if err != nil {
		return
	}
	err = OA.Redis.Set(ctx, Key, np, time.Minute*30).Err()
	return
}
