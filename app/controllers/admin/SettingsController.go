package admin

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type Settings struct {
	BaseController
}

func (c Settings) Index() revel.Result {
	options := optionService.GetAllOption()
	return c.RenderJSON(info.Res{Ok: true, Code: 200, List: options})
}

func (c Settings) Update() revel.Result {
	data := c.RequestData()
	for key, value := range data {
		optionService.UpdateByName(key, value)
	}
	return c.ResponseSuccess("update success.")
}
