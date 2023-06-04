package response

import "github.com/gin-gonic/gin"

type Resp interface {
	Succeed() gin.H
	Fail() gin.H
	Response(c *gin.Context)
}

func Response(c *gin.Context, r Resp) {
	r.Response(c)
}
