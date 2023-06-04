package audio2face

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/VMETA3/sui-hackathon/utils/request"
)

func (A2F *A2F) Connection() *A2F {
	resp, err := request.NewRequest(A2F.Host + string(Status)).Send()
	if err != nil {
		log.Fatalln("[Connection]", err)
	}

	if string(resp) != REPORT_SUCCESS {
		fmt.Println("[Connection]", "rep is ", resp)
		log.Fatalln("[Connection]", "The A2F server is not started")
	}

	log.Println("[Connection]", "SUCCESS")
	return A2F
}

func (A2F *A2F) VerifyUSD() *A2F {
AGAIN:
	resp, err := request.NewRequest(A2F.Host + string(GetInstances)).Send()
	if err != nil {
		log.Fatalln("[VerifyUSD]", err)
	}
	var Instances Instances
	json.Unmarshal(resp, &Instances)

	var Instance []string
	if len(Instances.Result.FullfaceInstances) > 0 {
		Instance = Instances.Result.FullfaceInstances
	} else if len(Instances.Result.RegularInstances) > 0 {
		Instance = Instances.Result.RegularInstances
	}

	if Instances.Status != "OK" || len(Instance) < 1 {
		log.Println("[VerifyUSD]", "FullfaceInstances is nil")
		A2F.SetUSD()
		goto AGAIN
	}

	match := false
	for _, v := range Instance {
		if v == A2F.A2FInstance {
			match = true
			break
		}
	}

	if !match {
		log.Println("[VerifyUSD]", "Unmatched Instance")
		log.Println("[VerifyUSD]", "Expectation", A2F.A2FInstance)
		log.Println("[VerifyUSD]", "Actual", Instance)
		log.Fatalln("[VerifyUSD]", "OVER")
	}

	log.Println("[VerifyUSD]", "SUCCESS")
	return A2F
}

func (A2F *A2F) VerifyPlayer() *A2F {
	resp, err := request.NewRequest(A2F.Host + string(PlayGetInstances)).Send()
	if err != nil {
		log.Fatalln("[VerifyPlayer]", err)
	}

	var PlayerInstances PlayerInstances
	json.Unmarshal(resp, &PlayerInstances)

	match := false
	for _, v := range PlayerInstances.Result {
		if v == A2F.A2FPlayer {
			match = true
			break
		}
	}
	if !match {
		log.Println("[VerifyPlayer]", "Unmatched Payer Instance")
		log.Println("[VerifyPlayer]", "Expectation", A2F.A2FPlayer)
		log.Println("[VerifyPlayer]", "Actual", PlayerInstances.Result)
		log.Fatalln("[VerifyPlayer]", "OVER")
	}

AGAIN:
	resp, err = request.NewRequest(A2F.Host+string(GetRootPath), []request.RequestOptions{
		request.SetMethod(request.POST),
		request.SetParams(request.ReqParams{
			"a2f_player": A2F.A2FPlayer,
		}),
	}...).Send()
	if err != nil {
		log.Fatalln("[VerifyPlayer]", err)
	}

	var RootPath *GETRootPath
	json.Unmarshal(resp, &RootPath)
	if RootPath.Details != nil {
		log.Fatalln("[VerifyPlayer]", RootPath.Details)
	}
	if RootPath.Result != A2F.PlayerPath {
		log.Println("[VerifyPlayer]", "Unmatched RootPath")
		log.Println("[VerifyPlayer]", "Expectation", A2F.PlayerPath)
		log.Println("[VerifyPlayer]", "Actual", RootPath.Result)
		A2F.SetRootPath()
		goto AGAIN
	}

	log.Println("[VerifyPlayer]", "SUCCESS")
	return A2F
}
