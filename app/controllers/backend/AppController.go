package backend

import (
	"github.com/revel/revel"
)

type Admin struct {
	BaseController
}

func (c Admin) Index() revel.Result {
	siteUrl := c.Request.Host
	siteName := "Persimmon"
	return c.Render(siteUrl,siteName)
}
