package openai

import (
	"github.com/VMETA3/sui-hackathon/utils"
	"github.com/VMETA3/sui-hackathon/utils/openai"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var oa *OPENAI

type OPENAI struct {
	Redis  *redis.Client
	Client *openai.OPENAI
}

func Init() *OPENAI {
	if oa == nil {
		DB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			log.Fatalln("REDIS_DB is incorrect,", err)
		}
		oa = &OPENAI{
			Redis: redis.NewClient(&redis.Options{
				Addr:     utils.GetENV("REDIS_HOST"),
				Password: os.Getenv("REDIS_PWD"), // no password set
				DB:       DB,                     // use default DB
			}),
			Client: openai.Init(),
		}
	}
	return oa
}
