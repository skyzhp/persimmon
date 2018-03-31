package home

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/utils"
	"gopkg.in/russross/blackfriday.v2"
)

type HomeComment struct {
	BaseHomeController
}

func (c HomeComment) List(postId int, page int) revel.Result {
	page = utils.IntDefault(page, 1)
	comments, err := commentService.GetCommentByPostId(postId, 50, page)
	if err != nil {
		return c.AjaxError(500, err.Error())
	}

	return c.RenderJSON(comments)
}

func (c HomeComment) Post(comment info.Comments) revel.Result {

	c.Validation.Required(comment.Name)
	c.Validation.Required(comment.Email)
	c.Validation.Required(comment.Markdown)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 501, Info: c.Validation.Errors})
	}

	htmlContent := blackfriday.Run([]byte(comment.Markdown))
	comment.Content = string(htmlContent)

	_, err := commentService.Save(comment)
	if err != nil {
		return c.AjaxError(500, err.Error())
	}

	go notificationService.SendCommentNotice(comment.PostsId, comment.Id, c.Request.Host)

	return c.AjaxSuccess("评论成功")
}