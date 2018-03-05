package backend

import (
	"github.com/revel/revel"
	"encoding/json"
	"github.com/cong5/persimmon/app/info"
)

type Navigation struct {
	BaseController
}

type NavigationMenu struct {
	name string
	url  string
}

func (c Navigation) Index() revel.Result {
	navigation := optionService.GetValueByName("navigations")
	NavigationMenu := NavigationMenu{}
	err := json.Unmarshal([]byte(navigation), NavigationMenu)
	if err != nil {
		return c.ResponseError(501, err.Error())
	}
	return c.RenderJSON(info.Res{Status: 200, Item: navigation})
}

func (c Navigation) Update(content *[]info.Navigation) revel.Result {
	optionValue, err := json.Marshal(content)
	if err != nil {
		return c.ResponseError(501, "Params error.")
	}

	//Validation
	c.Validation.Required(optionValue)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
	}

	ret := optionService.UpdateByName("navigations", string(optionValue))
	if ret {
		return c.ResponseError(500, "update failed.")
	}
	return c.ResponseSuccess("update success.")
}
