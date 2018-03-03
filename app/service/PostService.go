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

func (this *PostService) GetList(page int) []info.Posts {
	limit := 20
	start := (page - 1) * limit
	postList := make([]info.Posts, 0)
	err := db.MasterDB.Limit(limit, start).Find(&postList)
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

func (this *PostService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
