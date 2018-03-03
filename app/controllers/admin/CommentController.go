package admin

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
	"gopkg.in/russross/blackfriday.v2"
)

type Comment struct {
	BaseController
}

func (c Comment) Index(page int) revel.Result {
	if page <= 0 {
		page = 1
	}

	lists := commentService.GetList(page)
	return c.RenderJSON(info.Res{Ok: true, Code: 200, List: lists})
}

func (c Comment) Show(id int) revel.Result {
	link := commentService.GetOne(id)
	return c.RenderJSON(info.Res{Ok: true, Code: 200, Item: link})
}

func (c Comment) Store(comment *info.Comments) revel.Result {

	//Validation
	c.Validation.Required(comment.Name)
	c.Validation.Required(comment.Url)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Ok: false, Code: 501, Msg: c.Validation.Errors})
	}

	//save
	htmlContent := blackfriday.Run([]byte(comment.Markdown))
	comments := info.Comments{PostsId: comment.PostsId,
		Name: comment.Name,
		Email: comment.Email,
		Url: comment.Url,
		Content: string(htmlContent),
		Markdown: comment.Markdown,
		Ipaddress: c.ClientIP}

	ret := commentService.Save(comments)
	if ret > 0 {
		return c.ResponseSuccess("add success.")
	} else {
		return c.ResponseError(500, "Add failed.")
	}
}

func (c Comment) Update(content *info.Comments) revel.Result {

	if content.Id <= 0 {
		return c.ResponseError(501, "Params failed.")
	}

	//save
	comment := info.Comments{PostsId: content.PostsId,
		Name: content.Name,
		Email: content.Email,
		Url: content.Url,
		Content: content.Content,
		Markdown: content.Markdown,
		Ipaddress: c.ClientIP}

	ret := commentService.Update(content.Id, comment)
	if ret {
		return c.ResponseSuccess("Update success.")
	} else {
		return c.ResponseError(500, "Update failed.")
	}
}

func (c Comment) Destroy(id int) revel.Result {
	ret := commentService.Destroy(id, info.Comments{})
	if ret {
		return c.ResponseSuccess("Delete success.")
	} else {
		return c.ResponseError(500, "Delete failed.")
	}
}
