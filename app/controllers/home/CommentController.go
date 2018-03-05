package home

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type HomeComment struct {
	BaseHomeController
}

func (c HomeComment) List(postId int) revel.Result {
	return c.Render("List")
}

func (c HomeComment) Post(comment info.Comments) revel.Result {
	return c.Render("Post")
}
