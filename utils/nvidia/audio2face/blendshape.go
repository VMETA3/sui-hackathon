package audio2face

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/VMETA3/sui-hackathon/utils/request"
)

func (A2F *A2F) isExist(FileName string) (Name string, Exist bool) {
	Name = strings.TrimSuffix(FileName, path.Ext(FileName))
	File := A2F.BlendshapesPath + "\\" + Name + ".json"

	_, err := os.Stat(File)
	if err == nil {
		Exist = true
	}

	return
}

func (A2F *A2F) Audio2Blendshape(FileName string) (Json string, err error) {
	Name, Exist := A2F.isExist(FileName)
	if Exist {
		goto READ
	}

	err = A2F.SetTrack(FileName)
	if err != nil {
		return
	}

	err = A2F.GenerateKeys()
	if err != nil {
		return
	}

	err = A2F.ExportBlendshapes(Name)
	if err != nil {
		return
	}
READ:
	return A2F.ReadBlendshapes(A2F.BlendshapesPath + "\\" + Name + ".json")
}

func (A2F *A2F) ExportBlendshapes(Name string) (err error) {
	resp, err := request.NewRequest(A2F.Host+string(ExportBlendshapes), []request.RequestOptions{
		request.SetMethod(request.POST),
		request.SetParams(request.ReqParams{
			"solver_node":      A2F.SolverNode,
			"export_directory": A2F.BlendshapesPath,
			"file_name":        Name,
			"format":           "json",
			// "batch":            "false",
			// "fps":              "0",
		}),
	}...).Send()
	if err != nil {
		return
	}

	var Resp StandardResp
	json.Unmarshal(resp, &Resp)

	if Resp.Details != nil {
		log.Println("[ExportBlendshapes]", "["+A2F.BlendshapesPath+Name+".json]", Resp.Details)
		err = errors.New("ExportBlendshapes failure")
	}
	return
}

func (A2F *A2F) ReadBlendshapes(FilePath string) (Jaon string, err error) {
	file, err := os.Open(FilePath)
	if err != nil {
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	Jaon = string(content)
	return
}
