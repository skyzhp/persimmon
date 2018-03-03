package admin

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type Trash struct {
	BaseController
}

func (c Trash) Index(page int) revel.Result {
	if page <= 0 {
		page = 1
	}

	lists := postService.GetTrashList(page)
	return c.RenderJSON(info.Res{Ok: true, Code: 200, List: lists})
}

func (c Trash) Update(content *info.Posts) revel.Result {
	//save
	posts := info.Posts{}
	ret := postService.Update(content.Id, posts)
	if ret {
		return c.ResponseSuccess("Recover success.")
	} else {
		return c.ResponseError(500, "Recover failed.")
	}
}

func (c Trash) Destroy(id int) revel.Result {
	ret := postService.Destroy(id, info.Posts{})
	if ret {
		return c.ResponseSuccess("Delete success.")
	} else {
		return c.ResponseError(500, "Delete failed.")
	}
}
