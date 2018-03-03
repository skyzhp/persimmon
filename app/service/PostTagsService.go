package service

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/info"
)

type PostTagsService struct{}

func (this *PostTagsService) Save(postsTags []info.PostsTags) bool {
	_, err := db.MasterDB.Insert(&postsTags)
	if err != nil {
		revel.INFO.Printf("Save post tag failed: %s", err)
		return false
	}

	return true
}

func (this *PostTagsService) DeleteByPostID(postId int) bool {
	post := info.PostsTags{}
	_, err := db.MasterDB.Id(postId).Delete(post)
	if err != nil {
		revel.INFO.Printf("Destroy post tag error: %s", err)
		return false
	}
	return true
}
