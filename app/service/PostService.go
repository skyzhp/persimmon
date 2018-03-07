package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
)

type PostService struct{}

func (this *PostService) GetOne(id int) *info.Posts {
	post := &info.Posts{Id: id}
	_, err := db.MasterDB.Get(post)
	if err != nil {
		revel.INFO.Printf("Get post fatal : %s", err)
	}
	return post
}

func (this *PostService) GetList(categoryId int, keywords string, page int, limit int) []info.Posts {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}

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
	err := session.Limit(limit, start).Find(&postList)
	if err != nil {
		revel.INFO.Printf("Get posts list fatal : %s", err)
	}
	return postList
}

func (this *PostService) GetTrashList(page int) []info.Posts {
	limit := 20
	start := (page - 1) * limit
	postList := make([]info.Posts, 0)
	err := db.MasterDB.Unscoped().Limit(limit, start).Find(&postList)
	if err != nil {
		revel.INFO.Printf("Get posts list fatal : %s", err)
	}
	return postList
}

func (this *PostService) Save(post info.Posts) int64 {
	lastInsertId, err := db.MasterDB.InsertOne(post)
	if err != nil {
		return int64(0)
	}
	return lastInsertId
}

func (this *PostService) Update(id int, post info.Posts) bool {
	_, err := db.MasterDB.Id(id).Update(post)
	if err != nil {
		revel.INFO.Printf("Update post error: %s", err)
		return false
	}
	return true
}

func (this *PostService) Trash(id int, post info.Posts) bool {
	_, err := db.MasterDB.Id(id).Delete(post)
	if err != nil {
		revel.INFO.Printf("Destroy post error: %s", err)
		return false
	}
	return true
}

func (this *PostService) Destroy(id int, post info.Posts) bool {
	_, err := db.MasterDB.Id(id).Unscoped().Delete(post)
	if err != nil {
		revel.INFO.Printf("Destroy post error: %s", err)
		return false
	}
	return true
}

func (this *PostService) CountPost() int {
	post := new(info.Posts)
	postNum, err := db.MasterDB.Count(post)
	if err != nil {
		revel.INFO.Printf("Count post error: %s", err)
		return 0
	}
	return int(postNum)
}

func (this *PostService) CountTrashPost() int {
	post := new(info.Posts)
	postNum, err := db.MasterDB.Where("deleted_at != null").Count(post)
	if err != nil {
		revel.INFO.Printf("Count post error: %s", err)
		return 0
	}
	return int(postNum)
}

func (this *PostService) SumViews() int {
	post := new(info.Posts)
	viewNum, err := db.MasterDB.Sum(post, "views")
	if err != nil {
		revel.INFO.Printf("Count post error: %s", err)
		return 0
	}
	return int(viewNum)
}

func (this *PostService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
