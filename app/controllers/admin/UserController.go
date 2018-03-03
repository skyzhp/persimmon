package admin

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type User struct {
	BaseController
}

func (c User) Index() revel.Result {
	email := c.GetSession("Email")
	user := userService.GetUserByEmail(email)
	return c.RenderJSON(info.Res{Ok: true, Code: 200, Item: user})
}

func (c User) Update(id int, oldPwd string, pwd string, rePwd string) revel.Result {
	//Validation
	c.Validation.Required(oldPwd)
	c.Validation.Required(pwd)
	c.Validation.Required(rePwd)
	if c.Validation.HasErrors() {
		return c.RenderJSON(info.Res{Ok: false, Code: 501, Msg: c.Validation.Errors})
	}
	if pwd != rePwd {
		return c.ResponseError(501, "password and password_confirmation unequal.")
	}

	//Valid old password
	email := c.GetSession("Email")
	user := userService.GetUserByEmail(email)
	if authService.CheckPasswordHash(pwd, user.Password) {
		return c.ResponseError(500, "Old password error.")
	}

	ret := userService.UpdatePassword(id, pwd)
	if ret {
		return c.ResponseSuccess("Update password success.")
	} else {
		return c.ResponseError(500, "Update password failed.")
	}
}
