package backend

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
)

type Link struct {
	BaseController
}

func (c Link) Index(rows int, page int) revel.Result {
	lists, err := linkService.GetListPaging(rows, page)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Link) Show(id int) revel.Result {
	link, err := linkService.GetOne(id)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, Item: link})
}

func (c Link) Store(content *info.Links) revel.Result {

	//Validation
	c.Validation.Required(content.Name)
	c.Validation.Required(content.Url)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	link := info.Links{Name: content.Name,
		Logo: content.Logo,
		Group: content.Group,
		Url: content.Url,
		Ipaddress: c.ClientIP}

	_, err := linkService.Save(link)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
}

func (c Link) Update(content *info.Links) revel.Result {

	//save
	link := info.Links{Name: content.Name,
		Logo: content.Logo,
		Group: content.Group,
		Url: content.Url,
		Ipaddress: c.ClientIP}

	_, err := linkService.Update(content.Id, link)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Link) Destroy(id int) revel.Result {
	_, err := linkService.Destroy(id, info.Links{})
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Delete success.")
}
