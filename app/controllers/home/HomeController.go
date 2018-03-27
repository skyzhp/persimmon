package home

import (
	"github.com/revel/revel"
)

type Home struct {
	BaseHomeController
}

func (c Home) Index() revel.Result {
	limit := 10
	c.PostList(0, limit, 1)
	return c.Render()
}

func (c Home) Page(page int) revel.Result {
	limit := 10
	c.PostList(0, limit, page)
	return c.Render()
}

func (c Home) Post(slug string) revel.Result {
	c.GetGlobalInfo()
	post, err := postService.GetOneBySlug(slug)
	if slug == "" || err != nil {
		return c.NotFound("很抱歉，没有找到这个页面.")
	}

	return c.Render(post)
}

func (c Home) Tag(tag string) revel.Result {
	c.GetGlobalInfo()
	return c.Render()
}

func (c Home) Categories(slug string, page int) revel.Result {
	limit := 10
	category, cErr := categoryService.GetOneBySlug(slug)
	if cErr != nil {
		c.Flash.Error(cErr.Error())
	}
	c.PostList(category.Id, limit, page)

	return c.Render()
}

func (c Home) Friends() revel.Result {
	c.GetGlobalInfo()
	links, err := linkService.GetList(9999, 1)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	return c.Render(links)
}

func (c Home) Feed() revel.Result {
	c.GetGlobalInfo()
	return c.Render()
}

func (c Home) SiteMap() revel.Result {
	c.GetGlobalInfo()
	return c.Render()
}
