package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

type Utils struct {
	BaseController
}

const (
	baidu = "baidu"
	py    = "pinyin"
	//...
)

func (c Utils) Translate(words string) revel.Result {
	translateDriver := revel.Config.StringDefault("translate.driver", "pinyin")
	var res info.Res
	switch translateDriver {
	case baidu:
		if dst, err := baiduFanyiService.Fanyi(words); err == nil {
			res = info.Res{Status: 200, Item: dst}
		} else {
			res = info.Res{Status: 500, Info: err.Error()}
		}
	case py:
		newWords := pinyin.LazyPinyin(words, pinyin.NewArgs())
		wordsStrings := strings.Join(newWords, "")
		res = info.Res{Status: 200, Item: wordsStrings}
	default:
		res = info.Res{Status: 200, Item: ""}
	}

	return c.RenderJSON(res)
}
