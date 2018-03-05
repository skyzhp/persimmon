package backend

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
)

type Link struct {
	BaseController
}

func (c Link) Index(page int) revel.Result {
	if page <= 0 {
		page = 1
	}

	lists := linkService.GetList(page)
	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Link) Show(id int) revel.Result {
	link := linkService.GetOne(id)
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

	ret := linkService.Save(link)
	if ret > 0 {
		return c.ResponseSuccess("add success.")
	} else {
		return c.ResponseError(500, "Add failed.")
	}
}

func (c Link) Update(content *info.Links) revel.Result {

	//save
	link := info.Links{Name: content.Name,
		Logo: content.Logo,
		Group: content.Group,
		Url: content.Url,
		Ipaddress: c.ClientIP}

	ret := linkService.Update(content.Id, link)
	if ret {
		return c.ResponseSuccess("Update success.")
	} else {
		return c.ResponseError(500, "Update failed.")
	}
}

func (c Link) Destroy(id int) revel.Result {
	ret := linkService.Destroy(id, info.Links{})
	if ret {
		return c.ResponseSuccess("Delete success.")
	} else {
		return c.ResponseError(500, "Delete failed.")
	}
}
