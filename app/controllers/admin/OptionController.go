package admin

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
)

type Option struct {
	BaseController
}

func (c Option) Index(page int) revel.Result {
	if page <= 0 {
		page = 1
	}

	lists := optionService.GetList(page)
	return c.RenderJSON(info.Res{Ok: true, Code: 200, List: lists})
}

func (c Option) Show(id int) revel.Result {
	option := optionService.GetOne(id)
	return c.RenderJSON(info.Res{Ok: true, Code: 200, Item: option})
}

func (c Option) Store(content *info.Options) revel.Result {
	//Validation
	c.Validation.Required(content.OptionTitle)
	c.Validation.Required(content.OptionName)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Ok: false, Code: 501, Msg: c.Validation.Errors})
	}

	//save
	option := info.Options{OptionTitle: content.OptionTitle,
		OptionName: content.OptionName,
		OptionValue: content.OptionValue,
		OptionGroup: content.OptionGroup,
		OptionRemark: content.OptionRemark,
		OptionStatus: content.OptionStatus,
		DataType: content.DataType}

	ret := optionService.Save(option)
	if ret > 0 {
		return c.ResponseSuccess("add success.")
	} else {
		return c.ResponseError(500, "Add failed.")
	}
}

func (c Option) Update(content *info.Options) revel.Result {

	//save
	option := info.Options{OptionTitle: content.OptionTitle,
		OptionName: content.OptionName,
		OptionValue: content.OptionValue,
		OptionGroup: content.OptionGroup,
		OptionRemark: content.OptionRemark,
		OptionStatus: content.OptionStatus,
		DataType: content.DataType}

	ret := optionService.Update(content.Id, option)
	if ret {
		return c.ResponseSuccess("Update success.")
	} else {
		return c.ResponseError(500, "Update failed.")
	}
}

func (c Option) Destroy(id int) revel.Result {
	ret := optionService.Destroy(id, info.Options{})
	if ret {
		return c.ResponseSuccess("Delete success.")
	} else {
		return c.ResponseError(500, "Delete failed.")
	}
}
