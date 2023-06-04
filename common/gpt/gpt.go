package gpt

import (
	"github.com/VMETA3/sui-hackathon/utils"
	chatgpt "github.com/VMETA3/sui-hackathon/utils/chatGPT"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var G *GPT

type GPT struct {
	Redis *redis.Client
	TCG   *chatgpt.ChatGPT
}

func InitGPT() *GPT {
	if G == nil {
		DB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			log.Fatalln("REDIS_DB is incorrect,", err)
		}
		G = &GPT{
			Redis: redis.NewClient(&redis.Options{
				Addr:     utils.GetENV("REDIS_HOST"),
				Password: os.Getenv("REDIS_PWD"), // no password set
				DB:       DB,                     // use default DB
			}),
			TCG: chatgpt.InitGPT(),
		}
	}
	return G
}
