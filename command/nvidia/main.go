package main

import (
	"os"

	"github.com/VMETA3/sui-hackathon/service/nvidia"
	signal "github.com/VMETA3/sui-hackathon/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// init router
	r := RootRouter()

	Port := os.Getenv("NVIDIA_PORT")
	if Port == "" {
		Port = os.Getenv("SERVER_PORT")
		if Port == "" {
			Port = "8080"
		}
	}

	// listen and signal
	if err := signal.Listen(":"+Port, r); err != nil {
		panic("signal Listen failed:" + err.Error())
	}
}

func RootRouter() *gin.Engine {
	r := gin.Default()

	root := r.Group("/")

	nvidia.Router(root)
	return r
}
