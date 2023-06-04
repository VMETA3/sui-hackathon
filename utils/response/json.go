package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StandardResponse struct {
	Code uint64
	Msg  string
	Data interface{}
	Err  error
}

func (r *StandardResponse) Succeed() gin.H {
	if r.Msg == "" {
		r.Msg = "ok"
	}
	return gin.H{"code": r.Code, "message": r.Msg, "data": r.Data}
}

func (r *StandardResponse) Fail() gin.H {
	return gin.H{"code": -1, "message": r.Err.Error(), "data": nil}
}

func (r *StandardResponse) Response(c *gin.Context) {
	if r.Err != nil {
		c.JSON(http.StatusInternalServerError, r.Fail())
	} else {
		c.JSON(http.StatusOK, r.Succeed())
	}
}
