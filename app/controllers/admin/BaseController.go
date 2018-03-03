package admin

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"strconv"
	"strings"
)

type BaseController struct {
	*revel.Controller
}

func (c BaseController) ResponseSuccess(msg string) revel.Result {
	return c.RenderJSON(info.Res{Ok: true, Code: 200, Msg: msg})
}

func (c BaseController) ResponseError(code int, msg string) revel.Result {
	return c.RenderJSON(info.Res{Ok: false, Code: code, Msg: msg})
}

func (c BaseController) GetSession(key string) string {
	v, ok := c.Session[key]
	if !ok {
		v = ""
	}
	return v
}

func (c BaseController) SetSession(userInfo *info.Users) {
	c.Session["UserId"] = strconv.Itoa(userInfo.Id)
	c.Session["Email"] = string(userInfo.Email)
	c.Session["Username"] = string(userInfo.Name)
}

func (c BaseController) ClearSession() {
	delete(c.Session, "UserId")
	delete(c.Session, "Email")
	delete(c.Session, "Username")
}

func (c BaseController) Input(key string) string {
	values := c.Params.Values[key]
	return strings.Join(values, "")
}

func (c BaseController) RequestData() map[string]string {
	values := c.Params.Values
	request := make(map[string]string, len(values))
	for key, value := range values {
		request[key] = strings.Join(value, "")
	}
	return request
}

func (c BaseController) GetDataByJSON() map[string]interface{} {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	return jsonData
}
