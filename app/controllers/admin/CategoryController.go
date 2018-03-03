package admin

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
)

type Category struct {
	BaseController
}

func (c Category) Index(page int) revel.Result {
	if page <= 0 {
		page = 1
	}

	lists := categoryService.GetList(page)
	return c.RenderJSON(info.Res{Ok: true, Code: 200, List: lists})
}

func (c Category) Show(id int) revel.Result {
	category := categoryService.GetOne(id)
	return c.RenderJSON(info.Res{Ok: true, Code: 200, Item: category})
}

func (c Category) Store(content *info.Categorys) revel.Result {

	//Validation
	c.Validation.Required(content.CategoryName)
	c.Validation.Required(content.CategoryFlag)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Ok: false, Code: 501, Msg: c.Validation.Errors})
	}

	//save
	category := info.Categorys{CategoryName: content.CategoryName,
		CategoryDescription: content.CategoryDescription,
		CategoryFlag: content.CategoryFlag,
		CategoryParent: content.CategoryParent,
		Ipaddress: c.ClientIP}

	ret := categoryService.Save(category)
	if ret > 0 {
		return c.ResponseSuccess("add success.")
	} else {
		return c.ResponseError(500, "Add failed.")
	}
}

func (c Category) Update(content *info.Categorys) revel.Result {
	//save
	category := info.Categorys{CategoryName: content.CategoryName,
		CategoryDescription: content.CategoryDescription,
		CategoryFlag: content.CategoryFlag,
		CategoryParent: content.CategoryParent,
		Ipaddress: c.ClientIP}

	ret := categoryService.Update(content.Id, category)
	if ret {
		return c.ResponseSuccess("Update success.")
	} else {
		return c.ResponseError(500, "Update failed.")
	}
}

func (c Category) Destroy(id int) revel.Result {
	ret := categoryService.Destroy(id, info.Categorys{})
	if ret {
		return c.ResponseSuccess("Delete success.")
	} else {
		return c.ResponseError(500, "Delete failed.")
	}
}
