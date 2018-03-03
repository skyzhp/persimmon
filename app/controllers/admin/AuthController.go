package admin

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type Auth struct {
	BaseController
}

func (c Auth) Login(email string, pwd string) revel.Result {
	//Validation
	c.Validation.Required(email)
	c.Validation.Required(pwd)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Ok: false, Code: 400, Msg: c.Validation.Errors})
	}

	//Login Auth
	userInfo := userService.GetUserByEmail(email)
	AuthResult := authService.Login(pwd, userInfo)
	if AuthResult {
		c.SetSession(userInfo)
		return c.ResponseSuccess("Login success.")
	} else {
		return c.ResponseError(200, "Login failed.")
	}

}

func (c Auth) Logout() revel.Result {
	c.ClearSession()
	return c.Redirect("/")
}
