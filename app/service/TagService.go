package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
	"net/url"
)

type TagService struct{}

func (this *TagService) GetOne(id int) *info.Tags {
	option := &info.Tags{Id: id}
	_, err := db.MasterDB.Get(option)
	if err != nil {
		revel.INFO.Printf("Get tag fatal : %s", err)
	}
	return option
}

func (this *TagService) GetList(page int) []info.Tags {
	limit := 20
	start := (page - 1) * limit
	optionList := make([]info.Tags, 0)
	err := db.MasterDB.Limit(limit, start).Find(&optionList)
	if err != nil {
		revel.INFO.Printf("Get tag failed : %s", err)
	}
	return optionList
}

func (this *TagService) Save(postId int, tags []string) bool {
	tagsLen := len(tags)
	postTags := make([]info.PostsTags, tagsLen)
	for index := 0; index < tagsLen; index++ {
		//Exist?
		tagsName := tags[index]
		newTag := info.Tags{}
		_, err := db.MasterDB.Where("tags_name=?", tagsName).Get(newTag)
		if err != nil {
			revel.INFO.Printf("Get tag by tags_name failed : %s", err)
		}
		if newTag.Id > 0 {
			continue
		}
		//Not exist, Insert.
		var tagId int64
		var insertErr error
		newTag.TagsName = tagsName
		newTag.TagsFlag = url.QueryEscape(tagsName)
		if tagId, insertErr = db.MasterDB.Insert(newTag); insertErr != nil {
			revel.INFO.Printf("Get tag by tags_name failed : %s", err)
			continue
		}
		//construct post tags map[]
		postTags[index].PostsId = postId
		postTags[index].TagsId = int(tagId)
	}

	//Insert post tags array.
	postTagsService.DeleteByPostID(postId)
	postTagsService.Save(postTags)

	return true
}

func (this *TagService) SaveOne(tag info.Tags) int64 {
	lastInsertId, err := db.MasterDB.InsertOne(tag)
	if err != nil {
		revel.INFO.Printf("Get tag by tags_name failed : %s", err)
		return int64(0)
	}
	return lastInsertId
}

func (this *TagService) Update(id int, option info.Tags) bool {
	_, err := db.MasterDB.Id(id).Update(option)
	if err != nil {
		revel.INFO.Printf("Update tag error: %s", err)
		return false
	}
	return true
}

func (this *TagService) Destroy(id int, option info.Tags) bool {
	_, err := db.MasterDB.Id(id).Delete(option)
	if err != nil {
		revel.INFO.Printf("Destroy tag error: %s", err)
		return false
	}
	return true
}

func (this *TagService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
