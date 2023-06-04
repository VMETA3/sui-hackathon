package controllers

import (
	"errors"

	"github.com/VMETA3/sui-hackathon/utils/response"

	"github.com/gin-gonic/gin"
)

type GPTParams struct {
	Prompt     string `json:"Prompt"`
	Definition string `json:"Definition"`
}

func (C *Controller) GetPrompt(c *gin.Context) (Prompt string, err error) {
	var params GPTParams
	if err = c.ShouldBindJSON(&params); err != nil {
		return
	}
	Prompt = params.Prompt
	return
}

func (C *Controller) SingleConversation(c *gin.Context) {
	var params GPTParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Response(c, &response.StandardResponse{Err: err})
		return
	}

	// answer, err := C.G.Single(params.Prompt)
	answer, err := C.Openai.Single(params.Prompt, params.Definition)

	resp := &response.StandardResponse{
		Data: answer,
		Err:  err,
	}

	response.Response(c, resp)
}

func (C *Controller) ContinuousConversation(c *gin.Context) {
	SID := c.GetHeader("chat-id")
	if SID == "" {
		response.Response(c, &response.StandardResponse{Err: errors.New("the necessary parameter chat-id is missing")})
		return
	}
	var params GPTParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Response(c, &response.StandardResponse{Err: err})
		return
	}

	// answer, err := C.G.Continuous(SID, params.Prompt)
	answer, err := C.Openai.Continuous(SID, params.Prompt, params.Definition)

	response.Response(c, &response.StandardResponse{
		Data: answer,
		Err:  err,
	})
}
