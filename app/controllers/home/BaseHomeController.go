package home

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"strings"
	"bytes"
	"fmt"
	"github.com/cong5/persimmon/app/utils"
)

type BaseHomeController struct {
	*revel.Controller
}

func (c BaseHomeController) ResponseSuccess(msg string) revel.Result {
	return c.RenderJSON(info.Res{Status: 200, Info: msg})
}

func (c BaseHomeController) ResponseError(code int, msg string) revel.Result {
	return c.RenderJSON(info.Res{Status: code, Info: msg})
}

func (c BaseHomeController) RenderTemplateStr(templatePath string) string {
	// Get the Template.
	// 返回 GoTemplate{tmpl, loader}
	template, err := revel.MainTemplateLoader.Template(templatePath)
	if err != nil {
	}

	tpl := &revel.RenderTemplateResult{
		Template: template,
		ViewArgs: c.ViewArgs, // 把args给它
	}

	var buffer bytes.Buffer
	tpl.Template.Render(&buffer, c.ViewArgs)
	return buffer.String()
}

func (c BaseHomeController) SendCommentMail(postId int, comment info.Comments, host string) {

	tplPath := "app/views/mails/comments.html"

	post, err := postService.GetOne(postId)
	if err != nil {
		revel.INFO.Printf("find post by id failed: %s", err)
	}

	c.ViewArgs["host"] = host
	c.ViewArgs["comment"] = comment
	c.ViewArgs["subject"] = fmt.Sprintf("“%s” 有新的评论", post.Title)
	tplData := c.RenderTemplateStr(tplPath)

	revel.INFO.Printf("%s", tplData)
	/*
	this.SendMail(subject, body)
	if err != nil {
		return ret, err
	}*/
}

func (c BaseHomeController) Input(key string) string {
	values := c.Params.Values
	if values == nil {
		return ""
	}

	valStr := values[key]
	if len(valStr) == 0 {
		return ""
	}
	return strings.Join(valStr, "")
}

func (c BaseHomeController) GetGlobalInfo() {
	allOption, err := optionService.GetAllOption()
	if err == nil {
		for _, v := range allOption {
			c.ViewArgs[v.OptionName] = v.OptionValue
		}
	} else {
		revel.INFO.Printf("GetAllOption failed: %s", err)
	}

	//nav
	navi, navErr := navigationService.GetNavigation()
	if navErr != nil {
		revel.INFO.Printf("GetNavigation failed: %s", err)
	}
	c.ViewArgs["navigations"] = navi
}

func (c BaseHomeController) PostList(categoryId int, limit int, page int) {
	c.GetGlobalInfo()

	limit = utils.IntDefault(limit, 10)
	page = utils.IntDefault(page, 1)
	categoryId = utils.IntDefault(categoryId, 0)
	posts, err := postService.GetList(categoryId, "", limit, page)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	posts = utils.SubstrContent(posts)
	total, cErr := postService.CountPost()
	if cErr != nil {
		c.Flash.Error(cErr.Error())
	}

	totalPage := utils.GetTotalPage(total, limit)

	c.ViewArgs["posts"] = posts
	c.ViewArgs["page"] = page
	c.ViewArgs["total"] = total
	c.ViewArgs["totalPage"] = totalPage
}