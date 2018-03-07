package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
)

type CommentService struct{}

func (this *CommentService) GetOne(id int) *info.Comments {
	comment := &info.Comments{Id: id}
	_, err := db.MasterDB.Get(comment)
	if err != nil {
		revel.INFO.Printf("Get comment fatal : %s", err)
	}
	return comment
}

func (this *CommentService) GetList(page int) []info.Comments {
	limit := 20
	start := (page - 1) * limit
	commentList := make([]info.Comments, 0)
	err := db.MasterDB.Limit(limit, start).Find(&commentList)
	if err != nil {
		revel.INFO.Printf("Get comment fatal : %s", err)
	}
	return commentList
}

func (this *CommentService) Save(comment info.Comments) int64 {
	lastInsertId, err := db.MasterDB.InsertOne(comment)
	if err != nil {
		return int64(0)
	}
	return lastInsertId
}

func (this *CommentService) Update(id int, comment info.Comments) bool {
	_, err := db.MasterDB.Id(id).Update(comment)
	if err != nil {
		revel.INFO.Printf("Update comment error: %s", err)
		return false
	}
	return true
}

func (this *CommentService) Destroy(id int, comment info.Comments) bool {
	_, err := db.MasterDB.Id(id).Delete(comment)
	if err != nil {
		revel.INFO.Printf("Destroy comment error: %s", err)
		return false
	}
	return true
}

func (this *CommentService) CountComment() int {
	comment := new(info.Comments)
	postNum, err := db.MasterDB.Count(comment)
	if err != nil {
		revel.INFO.Printf("Count comment error: %s", err)
		return 0
	}
	return int(postNum)
}

func (this *CommentService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
