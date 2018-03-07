package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"gopkg.in/russross/blackfriday.v2"
)

type Posts struct {
	BaseController
}

func (c Posts) Index(page int, rows int, categoryId int, keywords string) revel.Result {
	lists := postService.GetList(categoryId, keywords, page, rows)
	//categoryList := categoryService.GetList(1000, 1)
	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Posts) Show(id int) revel.Result {
	post := postService.GetOne(id)
	return c.RenderJSON(info.Res{Status: 200, Item: post})
}

func (c Posts) Store(post *info.Post) revel.Result {
	//Validation
	c.Validation.Required(post.Title)
	c.Validation.Required(post.Flag)
	c.Validation.Required(post.Markdown)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	htmlContent := blackfriday.Run([]byte(post.Markdown))
	posts := info.Posts{Title: post.Title,
		Flag: post.Flag,
		Thumb: post.Thumb,
		CategoryId: post.CategoryId,
		UserId: post.UserId,
		Markdown: post.Markdown,
		Content: string(htmlContent),
		Ipaddress: c.ClientIP}

	postId := postService.Save(posts)
	if postId > 0 {
		//Insert post tags.
		tagService.Save(int(postId), post.Tags)
		return c.ResponseSuccess("add success.")
	} else {
		return c.ResponseError(500, "add failed.")
	}
}

func (c Posts) Update(post info.Post) revel.Result {
	//save
	htmlContent := blackfriday.Run([]byte(post.Markdown))
	posts := info.Posts{Title: post.Title,
		Flag: post.Flag,
		Thumb: post.Thumb,
		CategoryId: post.CategoryId,
		UserId: post.UserId,
		Markdown: post.Markdown,
		Content: string(htmlContent),
		Ipaddress: c.ClientIP}

	ret := postService.Update(post.Id, posts)
	if ret {
		//Update post tags
		tagService.Save(int(post.Id), post.Tags)
		return c.ResponseSuccess("Update success.")
	} else {
		return c.ResponseError(500, "Update failed.")
	}
}

func (c Posts) Destroy(id int) revel.Result {
	ret := postService.Trash(id, info.Posts{})
	if ret {
		return c.ResponseSuccess("Delete success.")
	} else {
		return c.ResponseError(500, "Delete failed.")
	}
}
