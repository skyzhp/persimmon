package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type Auth struct {
	BaseController
}

func (c Auth) Login(email string, password string) revel.Result {
	//Validation
	c.Validation.Required(email)
	c.Validation.Required(password)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Status: 400, Info: c.Validation.Errors})
	}

	//Login Auth
	userInfo := userService.GetUserByEmail(email)
	AuthResult := authService.Login(password, userInfo)
	if AuthResult {
		c.SetSession(userInfo)
		userInfo.Password = "Don't tell you."
		return c.RenderJSON(info.Res{Status: 200, Info: "Login success.", Item: userInfo})
	} else {
		return c.ResponseError(200, "Username or password error.")
	}

}

func (c Auth) Logout() revel.Result {
	c.ClearSession()
	return c.Redirect("/")
}
