package gpt

import (
	"github.com/VMETA3/sui-hackathon/service/gpt/controllers"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	Base := g.Group("/gpt")

	C := controllers.NewController()

	Base.POST("/single", C.SingleConversation)
	Base.POST("/continuous", C.ContinuousConversation)
}
