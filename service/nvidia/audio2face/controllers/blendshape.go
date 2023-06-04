package controllers

import (
	"mime/multipart"
	"strings"

	"github.com/VMETA3/sui-hackathon/utils/response"

	"github.com/gin-gonic/gin"
)

type BlendshapeParams struct {
	FileData *multipart.FileHeader `form:"file"`
}

func (C *Controller) GetBlendshape(c *gin.Context) {
	var params BlendshapeParams
	if err := c.ShouldBind(&params); err != nil {
		response.Response(c, &response.StandardResponse{Err: err})
		return
	}

	FileName, err := C.A2F.Upload(c, params.FileData)
	if err != nil {
		response.Response(c, &response.StandardResponse{Err: err})
		return
	}

	Json, err := C.A2F.Audio2Blendshape(FileName)
	if err != nil {
		response.Response(c, &response.StandardResponse{Err: err})
		return
	}
	Json = strings.ReplaceAll(Json, "\n", "")
	Json = strings.ReplaceAll(Json, "\r", "")
	Json = strings.ReplaceAll(Json, " ", "")
	resp := &response.StandardResponse{
		Data: Json,
	}
	response.Response(c, resp)
}
