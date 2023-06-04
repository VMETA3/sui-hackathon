package controllers

import (
	"github.com/VMETA3/sui-hackathon/common/gpt"
	"github.com/VMETA3/sui-hackathon/common/openai"
)

var Con *Controller

type Controller struct {
	G      *gpt.GPT
	Openai *openai.OPENAI
}

func NewController() *Controller {
	if Con == nil {
		Con = &Controller{
			G:      gpt.InitGPT(),
			Openai: openai.Init(),
		}
	}
	return Con
}
