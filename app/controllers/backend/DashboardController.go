package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type Dashboard struct {
	BaseController
}

func (c Dashboard) Index() revel.Result {
	siteUrl := c.Request.Host
	siteName := "Persimmon"
	return c.Render(siteUrl, siteName)
}

func (c Dashboard) MetaCount() revel.Result {
	post := postService.CountPost()
	postTrash := postService.CountTrashPost()
	comment := commentService.CountComment()
	viewsCount := postService.SumViews()
	metaInfo := info.MetaInfo{Posts: post,
		PostTrash: postTrash,
		Comments: comment,
		UserViews: viewsCount}
	return c.RenderJSON(info.Res{Status: 200, Item: metaInfo})
}
