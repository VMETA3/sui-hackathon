package audio2face

import (
	"log"

	"github.com/VMETA3/sui-hackathon/utils"
)

type A2F struct {
	Corve           bool
	Host            string
	USDFile         string
	A2FInstance     string
	A2FPlayer       string
	SolverNode      string
	PlayerPath      string
	BlendshapesPath string
}

func InitA2F() *A2F {
	A2F := &A2F{
		Corve:           false,
		Host:            utils.GetENV("A2F_HOST"),
		USDFile:         utils.GetENV("A2F_USD_FILE"),
		A2FInstance:     utils.GetENV("A2F_INSTANCE"),
		A2FPlayer:       utils.GetENV("A2F_PLAYER_INSTANCE"),
		SolverNode:      utils.GetENV("A2F_SOLVER_NODE"),
		PlayerPath:      utils.GetENV("A2F_PLAYER_PATH"),
		BlendshapesPath: utils.GetENV("A2F_BLENDSHAPES_PATH"),
	}

	A2F.Connection().VerifyUSD().VerifyPlayer()

	log.Println("[InitA2F]", "SUCCESS")
	return A2F
}
