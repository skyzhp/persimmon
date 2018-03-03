package home

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"strings"
)

type Base struct {
	*revel.Controller
}

func (c Base) ResponseSuccess(msg string) revel.Result {
	return c.RenderJSON(info.Res{Ok: true, Code: 200, Msg: msg})
}

func (c Base) ResponseError(code int, msg string) revel.Result {
	return c.RenderJSON(info.Res{Ok: false, Code: code, Msg: msg})
}

func (c Base) Input(key string) string {
	values := c.Params.Values[key]
	return strings.Join(values, "")
}