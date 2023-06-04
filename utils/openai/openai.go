package openai

import "github.com/VMETA3/sui-hackathon/utils"

type OPENAI struct {
	Key string
}

const HOST = "https://api.openai.com"

func Init() *OPENAI {
	return &OPENAI{
		Key: utils.GetENV("OPENAI_API_KEY"),
	}
}
