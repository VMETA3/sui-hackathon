package audio2face

import (
	"errors"
	"mime/multipart"
	"os"
	"path"

	"github.com/VMETA3/sui-hackathon/utils"

	"github.com/gin-gonic/gin"
)

type FileType string
type FileSuffix string

const (
	AudioMove   FileType = "audio/wave"
	OctetStream FileType = "application/octet-stream"

	WAV FileSuffix = "wav"
)

func (A2F *A2F) FileTypeFiltering(Filedata *multipart.FileHeader) (Pass bool) {
	for _, v := range Filedata.Header.Values("Content-Type") {
		if v == string(AudioMove) || v == string(OctetStream) {
			Pass = true
		}
	}

	if !Pass || string((path.Ext(Filedata.Filename))[1:]) != string(WAV) {
		return false
	}

	return true
}

func (A2F *A2F) Upload(c *gin.Context, Filedata *multipart.FileHeader) (FileName string, err error) {
	if !A2F.FileTypeFiltering(Filedata) {
		err = errors.New("unsupported file type")
		return
	}

	FileName = Filedata.Filename
	if !A2F.Corve {
		if _, err = os.Stat(A2F.PlayerPath + "\\" + FileName); err == nil {
			FileName = utils.GetHash32String() + "_" + Filedata.Filename
		}
	}
	err = c.SaveUploadedFile(Filedata, A2F.PlayerPath+"\\"+FileName)
	return
}
