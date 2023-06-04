package ai

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	Base := g.Group("/ai")

	// Base.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })
	Base.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "VMeta3 TEST AI"})
	})
}
