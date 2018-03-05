package home

import (
	"github.com/revel/revel"
)

type Home struct {
	BaseHomeController
}

func (c Home) Index() revel.Result {
	return c.Render("Index")
}

func (c Home) Post(slug string) revel.Result {
	return c.Render("Post")
}

func (c Home) Tag(tag string) revel.Result {
	return c.Render("Tag")
}

func (c Home) Categories(tag string) revel.Result {
	return c.Render("Tag")
}

func (c Home) Feed(tag string) revel.Result {
	return c.Render("Feed")
}