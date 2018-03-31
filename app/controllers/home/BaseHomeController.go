package home

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/utils"
	"net/http"
	"fmt"
)

const (
	PageSize  = 10
	FirstPage = 1
)

type BaseHomeController struct {
	*revel.Controller
}

type myRenderXML string

func (r myRenderXML) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "text/xml")
	resp.GetWriter().Write([]byte(r))
}

func (c BaseHomeController) AjaxSuccess(msg string) revel.Result {
	return c.RenderJSON(info.Res{Status: 200, Info: msg})
}

func (c BaseHomeController) AjaxError(code int, msg string) revel.Result {
	return c.RenderJSON(info.Res{Status: code, Info: msg})
}


func (c BaseHomeController) GetGlobalInfo() {
	allOption, err := optionService.GetAllOption(false)
	if err == nil {
		for _, v := range allOption {
			c.ViewArgs[v.OptionName] = v.OptionValue
		}
	} else {
		revel.INFO.Printf("GetAllOption failed: %s", err)
	}

	//nav
	navi, navErr := navigationService.GetNavigation(false)
	if navErr != nil {
		revel.INFO.Printf("GetNavigation failed: %s", err)
	}
	c.ViewArgs["navigations"] = navi
}

func (c BaseHomeController) ViewAssign(posts []info.Posts, total int, limit int, page int) {
	c.GetGlobalInfo()

	posts = utils.SubstrContent(posts)
	totalPage := utils.GetTotalPage(total, limit)

	prefix := "page"
	if tagName := c.Params.Route.Get("name"); tagName != "" {
		prefix = fmt.Sprintf("tag/%s", tagName)
	}

	c.ViewArgs["posts"] = posts
	c.ViewArgs["page"] = page
	c.ViewArgs["total"] = total
	c.ViewArgs["totalPage"] = totalPage
	c.ViewArgs["showPage"] = postService.ShowPage(prefix, page, totalPage)
}
