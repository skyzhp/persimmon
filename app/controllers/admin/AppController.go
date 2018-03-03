package admin

import (
	"github.com/revel/revel"
)

type Admin struct {
	BaseController
}

func (c Admin) Index() revel.Result {
	return c.Render()
}
