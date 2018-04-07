package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type Trash struct {
	BaseController
}

func (c Trash) Index(page int, rows int, categoryId int, keywords string) revel.Result {
	lists, err := postTrashService.GetTrashListPaging(categoryId, keywords, rows, page, false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(info.Res{Status: 200, List: lists})
}

func (c Trash) Update(ids []int) revel.Result {
	//save
	_, err := postTrashService.Restore(ids)
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
