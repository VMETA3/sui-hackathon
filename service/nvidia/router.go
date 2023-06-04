package nvidia

import (
	"github.com/VMETA3/sui-hackathon/service/nvidia/audio2face"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	Base := g.Group("/nvidia")

	audio2face.Router(Base)
}
