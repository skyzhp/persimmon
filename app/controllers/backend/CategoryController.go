package backend

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/utils"
)

type Category struct {
	BaseController
}

func (c Category) Index(rows int, page int) revel.Result {
	lists, err := categoryService.GetListPaging(rows, page, false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Category) Show(id int) revel.Result {
	category, err := categoryService.GetCategoryById(id, false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, Item: category})
}

func (c Category) Store(content *info.Categorys) revel.Result {

	//Validation
	c.Validation.Required(content.CategoryName)
	c.Validation.Required(content.CategoryFlag)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	clientIP := utils.Ip2long(c.ClientIP)
	category := info.Categorys{CategoryName: content.CategoryName,
		CategoryDescription: content.CategoryDescription,
		CategoryFlag: content.CategoryFlag,
		CategoryParent: content.CategoryParent,
		ClientIp: clientIP}

	_, err := categoryService.Save(category)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
}

func (c Category) Update(content *info.Categorys) revel.Result {
	//save
	clientIP := utils.Ip2long(c.ClientIP)
	category := info.Categorys{CategoryName: content.CategoryName,
		CategoryDescription: content.CategoryDescription,
		CategoryFlag: content.CategoryFlag,
		CategoryParent: content.CategoryParent,
		ClientIp: clientIP}

	_, err := categoryService.Update(content.Id, category)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Category) Destroy(ids []int) revel.Result {
	_, err := categoryService.Destroy(ids, info.Categorys{})
	if err != nil {
		return c.ResponseError(500, "Delete failed.")
	}

	return c.ResponseSuccess("Delete success.")
}
