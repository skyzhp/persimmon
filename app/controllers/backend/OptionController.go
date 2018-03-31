package backend

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
)

type Option struct {
	BaseController
}

func (c Option) Index(rows int, page int) revel.Result {
	lists, err := optionService.GetListPaging(rows, page)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Option) Show(id int) revel.Result {
	option, err := optionService.GetOptionById(id)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, Item: option})
}

func (c Option) Store(content *info.Options) revel.Result {
	//Validation
	c.Validation.Required(content.OptionTitle)
	c.Validation.Required(content.OptionName)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	option := info.Options{OptionTitle: content.OptionTitle,
		OptionName: content.OptionName,
		OptionValue: content.OptionValue,
		OptionGroup: content.OptionGroup,
		OptionRemark: content.OptionRemark,
		OptionStatus: content.OptionStatus,
		DataType: content.DataType}

	_, err := optionService.Save(option)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
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

	_, err := optionService.Update(content.Id, option)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Option) Destroy(id int) revel.Result {
	_, err := optionService.Destroy(id, info.Options{})
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Delete success.")
}
