package backend

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/revel/revel"
	"gopkg.in/russross/blackfriday.v2"
)

type Comment struct {
	BaseController
}

func (c Comment) Index(rows int, page int) revel.Result {
	lists, err := commentService.GetListPaging(rows, page)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Comment) Show(id int) revel.Result {
	comment, err := commentService.GetOne(id)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, Item: comment})
}

func (c Comment) Store(comment *info.Comments) revel.Result {
	//Validation
	c.Validation.Required(comment.Name)
	c.Validation.Required(comment.Url)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
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

	_, err := commentService.Save(comments)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
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

	_, err := commentService.Update(content.Id, comment)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Comment) Destroy(id int) revel.Result {
	_, err := commentService.Destroy(id, info.Comments{})
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Delete success.")
}
