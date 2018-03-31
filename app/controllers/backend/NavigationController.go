package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type Navigation struct {
	BaseController
}

func (c Navigation) Index() revel.Result {
	navigation, err := navigationService.GetNavigation(false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, List: navigation})
}

func (c Navigation) Update(nav string) revel.Result {
	//Validation
	c.Validation.Required(nav)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
	}
	_, err := optionService.UpdateByName(navigationService.GetNavKey(), nav)
	if err != nil {
		return c.ResponseError(500, "update failed.")
	}

	return c.ResponseSuccess("update success.")
}
