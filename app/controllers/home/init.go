package home

import (
	"github.com/cong5/persimmon/app/service"
	"github.com/revel/revel"
	"fmt"
)

var authService, AuthS *service.AuthService
var userService, UserS *service.UserService
var categoryService, CategoryS *service.CategoryService
var optionService, OptionS *service.OptionService
var linkService, LinkS *service.LinkService
var commentService, commentS *service.CommentService
var tagService, TagS *service.TagService
var postService, PostS *service.PostService
var postTagsService, PostTagsS *service.PostTagsService
var uploadsService, UploadsS *service.UploadsService
var baiduFanyiService, baiduFanyiS *service.BaiduFanyiService
var notificationService, notificationS *service.NotificationService
var postTrashService, postTrashS *service.PostTrashService
var navigationService, navigationS *service.NavigationService

func InitService() {
	AuthS = &service.AuthService{}
	UserS = &service.UserService{}
	CategoryS = &service.CategoryService{}
	OptionS = &service.OptionService{}
	LinkS = &service.LinkService{}
	commentS = &service.CommentService{}
	TagS = &service.TagService{}
	PostS = &service.PostService{}
	postTrashS = &service.PostTrashService{}
	PostTagsS = &service.PostTagsService{}
	UploadsS = &service.UploadsService{}
	baiduFanyiS = &service.BaiduFanyiService{}
	notificationS = &service.NotificationService{}
	navigationS = &service.NavigationService{}

	categoryService = CategoryS
	userService = UserS
	authService = AuthS
	optionService = OptionS
	linkService = LinkS
	commentService = commentS
	tagService = TagS
	postService = PostS
	postTrashService = postTrashS
	postTagsService = PostTagsS
	uploadsService = UploadsS
	baiduFanyiService = baiduFanyiS
	notificationService = notificationS
	navigationService = navigationS
}

func NewTemplateFunc() {
	revel.TemplateFuncs["previous"] = func(val int) string {
		pageNum := val - 1
		if val <= 1 {
			pageNum = 1
		}
		return fmt.Sprintf("/page/%d", pageNum)
	}
	revel.TemplateFuncs["next"] = func(val int) string {
		return fmt.Sprintf("/page/%d", val+1)
	}
}

func init() {
	//服务
	InitService()
	NewTemplateFunc()
}
