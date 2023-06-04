package main

import (
	"os"

	"github.com/VMETA3/sui-hackathon/service/ai"
	"github.com/VMETA3/sui-hackathon/service/gpt"
	signal "github.com/VMETA3/sui-hackathon/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// init router
	r := RootRouter()

	Port := os.Getenv("SERVER_PORT")
	if Port == "" {
		Port = "8080"
	}

	// listen and signal
	if err := signal.Listen(":"+Port, r); err != nil {
		panic("signal Listen failed:" + err.Error())
	}
}

func RootRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/ai/*")

	root := r.Group("/")
	gpt.Router(root)
	ai.Router(root)
	// nvidia.Router(root)

	return r
}
