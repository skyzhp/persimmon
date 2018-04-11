package home

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/utils"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/adtac/go-akismet/akismet"
)

type HomeComment struct {
	BaseHomeController
}

func (c HomeComment) List(postId int, page int) revel.Result {
	if postId <= 0 {
		return c.AjaxError(501, "Post ID empty.")
	}

	page = utils.IntDefault(page > 0, page, 1)
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

	//垃圾评论检测
	domain := c.Request.Host
	commentType := utils.StringDefault(comment.ParentId <= 0, "comment", "reply")
	akismetKey := revel.Config.StringDefault("akismet.key", "")
	isSpam, _ := akismet.Check(&akismet.Comment{Blog: "https://" + domain,
		UserIP: c.ClientIP,
		UserAgent: c.Request.UserAgent(),
		CommentType: commentType,
		CommentAuthor: comment.Name,
		CommentAuthorEmail: comment.Email,
		CommentAuthorURL: comment.Url,
		CommentContent: comment.Content,
	}, akismetKey)
	status := utils.IntDefault(isSpam, 2, 1)

	//save
	clientIP := utils.Ip2long(c.ClientIP)
	comments := info.Comments{PostsId: comment.PostsId,
		Name: comment.Name,
		Email: comment.Email,
		Url: comment.Url,
		Content: comment.Content,
		Markdown: comment.Markdown,
		ClientIp: clientIP,
		Status: status}

	if comments.Content == "" {
		htmlContent := blackfriday.Run([]byte(comment.Markdown))
		comments.Content = string(htmlContent)
	}

	commentId, err := commentService.Save(comments)
	if err != nil {
		return c.AjaxError(500, err.Error())
	}

	go notificationService.SendCommentNotice(comment.PostsId, commentId, c.Request.Host)

	return c.AjaxSuccess("评论成功")
}
