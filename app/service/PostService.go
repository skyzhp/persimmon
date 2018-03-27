package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"time"
	"github.com/cong5/persimmon/app/utils"
	"errors"
	"github.com/revel/revel"
)

type PostService struct{}

func (this *PostService) GetOne(id int) (*info.Posts, error) {
	post := &info.Posts{Id: id}
	_, err := db.MasterDB.Get(post)
	if err != nil {
		return post, err
	}
	//related tags
	tags, err := tagService.GetListByPostId(post.Id)
	if err != nil {
		revel.INFO.Printf("GetListByPostId failed : %s", tags)
		return nil, err
	}

	return post, nil
}

func (this *PostService) GetOneBySlug(slug string) (*info.Posts, error) {
	post := &info.Posts{}
	if slug == "" {
		return post, errors.New("Slug is empty.")
	}
	_, err := db.MasterDB.Where("flag = ?", slug).Get(post)
	if err != nil {
		return nil, err
	}
	//related tags
	tags, err := tagService.GetListByPostId(post.Id)
	if err != nil {
		revel.INFO.Printf("GetListByPostId failed : %s", tags)
		return nil, err
	}

	return post, nil
}

func (this *PostService) GetList(categoryId int, keywords string, limit int, page int) ([]info.Posts, error) {

	limit = utils.IntDefault(limit,20)
	page = utils.IntDefault(page,1)
	session := db.MasterDB.NewSession()
	start := (page - 1) * limit
	postList := make([]info.Posts, 0)
	if categoryId > 0 {
		session.And("category_id = ?", categoryId)
	}
	if len(keywords) > 0 {
		keywordsStr := "%" + keywords + "%"
		session.And("title like ?", keywordsStr)
	}
	err := session.Desc("id").Limit(limit, start).Find(&postList)
	if err != nil {
		return nil, err
	}

	return this.JoinRelated(postList)
}

func (this *PostService) GetListPaging(categoryId int, keywords string, limit int, page int) *info.PagingContent {
	dataList, _ := this.GetList(categoryId, keywords, limit, page)
	total, _ := this.CountPost()
	totalPage := utils.GetTotalPage(total, limit)
	pagingContent := &info.PagingContent{Data: dataList,
		Total: total,
		TotalPage: totalPage,
		CurrentPage: page}
	return pagingContent
}

func (this *PostService) JoinRelated(posts []info.Posts) ([]info.Posts, error) {
	if len(posts) <= 0 {
		return posts, nil
	}
	categoryList, err := categoryService.GetList(999, 1)
	if err != nil {
		return nil, err
	}
	categories := make(map[int]info.Categorys, len(categoryList))
	for _, category := range categoryList {
		categories[category.Id] = category
	}
	for i, post := range posts {
		posts[i].Categories = categories[post.CategoryId]
		tags, err := tagService.GetListByPostId(post.Id)
		if err == nil {
			posts[i].Tags = tags
		} else {
			revel.INFO.Printf("GetListByPostId failed : %s", err)
		}
	}

	return posts, nil
}

func (this *PostService) Save(post info.Posts) (int, error) {
	if _, err := db.MasterDB.InsertOne(post); err != nil {
		//revel.INFO.Printf("Save post failed : %s", err)
		return 0, err
	}
	return post.Id, nil
}

func (this *PostService) Update(id int, post info.Posts) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(post)
	if err != nil {
		//revel.INFO.Printf("Update post failed: %s", err)
		return false, err
	}
	return true, nil
}

// add to Trash
func (this *PostService) Trash(ids []int, post info.Posts) (bool, error) {
	_, err := db.MasterDB.In("id", ids).Delete(post)
	if err != nil {
		//revel.INFO.Printf("Destroy post failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *PostService) CountPost() (int, error) {
	post := new(info.Posts)
	total, err := db.MasterDB.Count(post)
	if err != nil {
		//revel.INFO.Printf("Count post failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *PostService) SumViews() (int, error) {
	post := new(info.Posts)
	total, err := db.MasterDB.Sum(post, "views")
	if err != nil {
		//revel.INFO.Printf("Count post failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *PostService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
