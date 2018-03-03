package home

import (
	"github.com/cong5/persimmon/app/service"
)

var authService, AuthS *service.AuthService
var xmlrpcService, xmlrpcS *service.XmlRpcService
var userService, UserS *service.UserService
var categoryService, CategoryS *service.CategoryService

func InitService() {
	AuthS = &service.AuthService{}
	UserS = &service.UserService{}
	xmlrpcS = &service.XmlRpcService{}

	authService = AuthS
	xmlrpcService = xmlrpcS
	userService = UserS
	categoryService = CategoryS
}

func init() {
	//服务
	InitService()
}
