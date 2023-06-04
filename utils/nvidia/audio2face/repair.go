package audio2face

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/VMETA3/sui-hackathon/utils/request"
)

func (A2F *A2F) SetUSD() {
	resp, err := request.NewRequest(A2F.Host+string(USDLoad), []request.RequestOptions{
		request.SetMethod(request.POST),
		request.SetParams(request.ReqParams{
			"file_name": A2F.USDFile,
		}),
	}...).Send()
	if err != nil {
		log.Fatalln("[SetUSD]", err)
	}

	var Resp StandardResp
	json.Unmarshal(resp, &Resp)

	if Resp.Details != nil {
		log.Fatalln("[SetUSD]", Resp.Details)
	}
	log.Println("[SetUSD]", "SUCCESS")
}

func (A2F *A2F) SetRootPath() {
	resp, err := request.NewRequest(A2F.Host+string(SetRootPath), []request.RequestOptions{
		request.SetMethod(request.POST),
		request.SetParams(request.ReqParams{
			"a2f_player": A2F.A2FPlayer,
			"dir_path":   A2F.PlayerPath,
		}),
	}...).Send()
	if err != nil {
		log.Fatalln("[SetRootPath]", err)
	}

	var Resp StandardResp
	json.Unmarshal(resp, &Resp)

	if Resp.Details != nil {
		log.Fatalln("[SetRootPath]", Resp.Details)
	}
	log.Println("[SetRootPath]", "SUCCESS")
}

func (A2F *A2F) SetTrack(FileName string) (err error) {
	resp, err := request.NewRequest(A2F.Host+string(SetTrack), []request.RequestOptions{
		request.SetMethod(request.POST),
		request.SetParams(request.ReqParams{
			"a2f_player": A2F.A2FPlayer,
			"file_name":  FileName,
			"time_range": []int{0, -1},
		}),
	}...).Send()
	if err != nil {
		return
	}

	var Resp StandardResp
	json.Unmarshal(resp, &Resp)

	if Resp.Details != nil {
		log.Println("[SetTrack]", "["+FileName+"]", Resp.Details)
		err = errors.New("SetTrack failure")
	}
	return
}

func (A2F *A2F) GenerateKeys() (err error) {
	resp, err := request.NewRequest(A2F.Host+string(GenerateKeys), []request.RequestOptions{
		request.SetMethod(request.POST),
		request.SetParams(request.ReqParams{
			"a2f_instance": A2F.A2FInstance,
		}),
	}...).Send()
	if err != nil {
		return
	}

	var Resp StandardResp
	json.Unmarshal(resp, &Resp)

	if Resp.Details != nil {
		log.Println("[GenerateKeys]", Resp.Details)
		err = errors.New("GenerateKeys failure")
	}
	return
}
