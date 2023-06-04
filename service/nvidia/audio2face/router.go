package audio2face

import (
	"github.com/VMETA3/sui-hackathon/service/nvidia/audio2face/controllers"
	"github.com/VMETA3/sui-hackathon/utils/middleware"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	C := controllers.NewController()

	Base := g.Group("/a2f")
	Base.Use(middleware.VerifyUser())
	{
		Base.POST("/getBlendshape", C.GetBlendshape)
	}

}
