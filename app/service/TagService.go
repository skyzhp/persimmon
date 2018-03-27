package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
	"net/url"
	"fmt"
	"github.com/cong5/persimmon/app/utils"
)

type TagService struct{}

func (this *TagService) GetOne(id int) (*info.Tags, error) {
	tag := &info.Tags{Id: id}
	if _, err := db.MasterDB.Get(tag); err != nil {
		//revel.INFO.Printf("Get tag failed : %s", err)
		return nil, err
	}
	return tag, nil
}

func (this *TagService) GetList(limit int, page int) ([]info.Tags, error) {

	limit = utils.IntDefault(limit, 20)
	page = utils.IntDefault(page, 1)
	start := (page - 1) * limit
	tagsList := make([]info.Tags, 0)
	err := db.MasterDB.Limit(limit, start).Find(&tagsList)
	if err != nil {
		//revel.INFO.Printf("Get tag failed : %s", err)
		return nil, err
	}
	return tagsList, nil
}

func (this *TagService) GetListPaging(limit int, page int) (*info.PagingContent, error) {
	dataList, err := this.GetList(limit, page)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountTags()
	if cErr != nil {
		return nil, cErr
	}

	totalPage := utils.GetTotalPage(total, limit)
	pagingContent := &info.PagingContent{Data: dataList,
		Total: total,
		TotalPage: totalPage,
		CurrentPage: page}
	return pagingContent, nil
}

func (this *TagService) CountTags() (int, error) {
	tag := new(info.Tags)
	total, err := db.MasterDB.Count(tag)
	if err != nil {
		//revel.INFO.Printf("Count tag failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *TagService) GetListByPostId(postId int) ([]info.Tags, error) {
	tagsList := make([]info.Tags, 0)

	tagsTable := db.MasterDB.TableMapper.Obj2Table("tags")
	postsTagsTable := db.MasterDB.TableMapper.Obj2Table("posts_tags")
	joinWhere := fmt.Sprintf("%s.tags_id = %s.id", postsTagsTable, tagsTable)
	where := fmt.Sprintf("%s.posts_id = ?", postsTagsTable)

	err := db.MasterDB.Table(tagsTable).Join("LEFT OUTER", postsTagsTable, joinWhere).Where(where, postId).Find(&tagsList)
	if err != nil {
		//revel.INFO.Printf("Get tag failed : %s", err)
		return nil, err
	}
	return tagsList, nil
}

func (this *TagService) Save(postId int, tags []string) (bool, error) {
	tagsLen := len(tags)
	postTags := make([]info.PostsTags, tagsLen)
	for index := 0; index < tagsLen; index++ {
		//Exist?
		tagsName := tags[index]
		newTag := info.Tags{}
		if _, err := db.MasterDB.Where("tags_name=?", tagsName).Get(&newTag); err != nil {
			//revel.INFO.Printf("Get tag by tags_name failed : %s", err)
			return false, err
		}
		if newTag.Id < 0 {
			//Not exist, Insert.
			newTag.TagsName = tagsName
			newTag.TagsFlag = url.QueryEscape(tagsName)
			if _, insertErr := db.MasterDB.Insert(&newTag); insertErr != nil {
				revel.INFO.Printf("Get tag by tags_name failed : %s", insertErr)
				continue
			}
		}

		//construct post tags map[]
		postTags[index].PostsId = postId
		postTags[index].TagsId = newTag.Id
	}

	//Insert post tags array.
	postTagsService.DeleteByPostID(postId)
	postTagsService.Save(postTags)

	return true, nil
}

func (this *TagService) SaveOne(tag info.Tags) (int, error) {
	if _, err := db.MasterDB.InsertOne(tag); err != nil {
		//revel.INFO.Printf("Get tag by tags_name failed : %s", err)
		return 0, err
	}

	return tag.Id, nil
}

func (this *TagService) Update(id int, tag info.Tags) (bool, error) {
	if _, err := db.MasterDB.Id(id).Update(tag); err != nil {
		//revel.INFO.Printf("Update tag failed: %s", err)
		return false, err
	}

	return true, nil
}

func (this *TagService) Destroy(id int, tag info.Tags) (bool, error) {
	if _, err := db.MasterDB.Id(id).Delete(tag); err != nil {
		//revel.INFO.Printf("Destroy tag failed: %s", err)
		return false, err
	}

	return true, nil
}

func (this *TagService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
