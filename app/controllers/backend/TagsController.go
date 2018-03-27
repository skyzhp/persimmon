package backend

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
	"net/url"
)

type Tags struct {
	BaseController
}

func (c Tags) Index(rows int, page int) revel.Result {
	lists, err := tagService.GetListPaging(rows, page)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Tags) Show(id int) revel.Result {
	tag, err := tagService.GetOne(id)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

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

	_, err := tagService.SaveOne(tag)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
}

func (c Tags) Update(id int, tagName string, tagFlag string) revel.Result {
	//save
	tag := info.Tags{TagsName: tagName,
		TagsFlag: url.QueryEscape(tagFlag)}

	_, err := tagService.Update(id, tag)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Tags) Destroy(id int) revel.Result {
	_, err := tagService.Destroy(id, info.Tags{})
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Delete success.")
}
