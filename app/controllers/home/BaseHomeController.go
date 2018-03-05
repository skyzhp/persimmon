package home

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"strings"
)

type BaseHomeController struct {
	*revel.Controller
}

func (c BaseHomeController) ResponseSuccess(msg string) revel.Result {
	return c.RenderJSON(info.Res{Status: 200, Info: msg})
}

func (c BaseHomeController) ResponseError(code int, msg string) revel.Result {
	return c.RenderJSON(info.Res{Status: code, Info: msg})
}

func (c BaseHomeController) Input(key string) string {
	values := c.Params.Values
	if values == nil {
		return ""
	}

	valStr := values[key]
	if len(valStr) == 0 {
		return ""
	}
	return strings.Join(valStr, "")
}
