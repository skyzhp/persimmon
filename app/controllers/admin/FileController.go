package admin

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/utils"
	"io/ioutil"
	"image"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
	"bytes"
	"time"
	"fmt"
	"os"
)

type File struct {
	BaseController
}

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

const (
	qiniu = "qiniu"
	//...
)

func (c File) Uploads(file []byte) revel.Result {
	//Validation
	fileBuffer := bytes.NewReader(file)
	_, format, err := image.DecodeConfig(fileBuffer)
	if err != nil {
		return c.ResponseError(501, err.Error())
	}

	c.Validation.Required(file)
	c.Validation.MinSize(file, 2*KB).Message("Minimum a file size of 2KB expected")
	c.Validation.MaxSize(file, 9*MB).Message("File cannot be larger than 9MB")
	c.Validation.Required(err == nil).Key("file").Message("Incorrect file format")
	//c.Validation.Required(format == "jpeg" || format == "png").Message( "JPEG or PNG file format is expected")
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Ok: false, Code: 501, Msg: c.Validation.Errors})
	}

	nowDate := time.Now().Format("2006-01")
	uploadsDir := fmt.Sprintf("public/uploads/%s/%d", nowDate, time.Now().Day())
	mkErr := os.MkdirAll(uploadsDir, 0755)
	if mkErr != nil {
		return c.ResponseError(501, mkErr.Error())
	}
	if format == "jpeg" {
		format = "jpg"
	}

	filePath := fmt.Sprintf("%s/%s.%s", uploadsDir, utils.NewGuid(), format)
	saveErr := ioutil.WriteFile(filePath, file, 0777)

	if saveErr != nil {
		return c.ResponseError(500, err.Error())
	}

	fileSystem := revel.Config.StringDefault("file.system", "")
	var resInfo info.Res
	switch fileSystem {
	case qiniu:
		resInfo = uploadsService.QiniuUploads(format, filePath)
		//...还可以有其他的上传方式
	default:
		resInfo = info.Res{Ok: true, Code: 200, Item: fileSystem}
	}
	return c.RenderJSON(resInfo)
}