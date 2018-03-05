package backend

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
	"net/url"
)

type Tags struct {
	BaseController
}

func (c Tags) Index(page int) revel.Result {
	if page <= 0 {
		page = 1
	}

	lists := tagService.GetList(page)
	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Tags) Show(id int) revel.Result {
	tag := tagService.GetOne(id)
	return c.RenderJSON(info.Res{Status: 200, Item: tag})
}

func (c Tags) Store(tagName string, tagFlag string) revel.Result {
	//Validation
	c.Validation.Required(tagName)
	c.Validation.Required(tagFlag)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	tag := info.Tags{TagsName: tagName,
		TagsFlag: url.QueryEscape(tagFlag)}

	ret := tagService.SaveOne(tag)
	if ret > 0 {
		return c.ResponseSuccess("add success.")
	} else {
		return c.ResponseError(500, "Add failed.")
	}
}

func (c Tags) Update(id int,tagName string, tagFlag string) revel.Result {
	//save
	tag := info.Tags{TagsName: tagName,
		TagsFlag: url.QueryEscape(tagFlag)}

	ret := tagService.Update(id, tag)
	if ret {
		return c.ResponseSuccess("Update success.")
	} else {
		return c.ResponseError(500, "Update failed.")
	}
}

func (c Tags) Destroy(id int) revel.Result {
	ret := tagService.Destroy(id, info.Tags{})
	if ret {
		return c.ResponseSuccess("Delete success.")
	} else {
		return c.ResponseError(500, "Delete failed.")
	}
}
