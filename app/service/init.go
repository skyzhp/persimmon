package service


var authService, AuthS *AuthService
var userService, UserS *UserService
var categoryService, CategoryS *CategoryService
var optionService, OptionS *OptionService
var linkService, LinkS *LinkService
var commentService, commentS *CommentService
var tagService, TagS *TagService
var postService, PostS *PostService
var postTagsService, PostTagsS *PostTagsService
var uploadsService, UploadsS *UploadsService
var baiduFanyiService, baiduFanyiS *BaiduFanyiService

func InitService() {
	AuthS = &AuthService{}
	UserS = &UserService{}
	CategoryS = &CategoryService{}
	OptionS = &OptionService{}
	LinkS = &LinkService{}
	commentS = &CommentService{}
	TagS = &TagService{}
	PostS = &PostService{}
	PostTagsS = &PostTagsService{}
	UploadsS = &UploadsService{}
	baiduFanyiS = &BaiduFanyiService{}

	categoryService = CategoryS
	userService = UserS
	authService = AuthS
	optionService = OptionS
	linkService = LinkS
	commentService = commentS
	tagService = TagS
	postService = PostS
	postTagsService = PostTagsS
	uploadsService = UploadsS
	baiduFanyiService = baiduFanyiS
}