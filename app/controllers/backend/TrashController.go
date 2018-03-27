package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type Trash struct {
	BaseController
}

func (c Trash) Index(page int, rows int, categoryId int, keywords string) revel.Result {
	lists, err := postTrashService.GetTrashListPaging(categoryId, keywords, rows, page)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Trash) Update(content *info.Posts) revel.Result {
	//save
	posts := info.Posts{}
	_, err := postService.Update(content.Id, posts)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Restore success.")
}

func (c Trash) Destroy(ids []int) revel.Result {
	_, err := postTrashService.Destroy(ids)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Delete success.")
}
