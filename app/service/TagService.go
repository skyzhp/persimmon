package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"net/url"
	"errors"
	"time"
)

type TagService struct{}

func (this *TagService) GetTagById(id int, real bool) (info.Tags, error) {
	tag := info.Tags{}
	cacheKey := utils.CacheKey("TagService", "InfoById", id)
	if err := cache.Get(cacheKey, &tag); err != nil || real {
		_, err := db.MasterDB.Where("id = ?", id).Get(&tag)
		if err != nil {
			return tag, err
		}
		go cache.Set(cacheKey, tag, 30*time.Minute)
	}

	return tag, nil
}

func (this *TagService) GetTagByIdArr(idArr []int, real bool) ([]info.Tags, error) {
	idLen := len(idArr)
	if idLen <= 0 {
		return nil, errors.New("参数不正确")
	}

	tagArr := make([]info.Tags, idLen)

	for k, postId := range idArr {
		post, err := this.GetTagById(postId, real)
		if err == nil {
			tagArr[k] = post
		}
	}

	return tagArr, nil
}

func (this *TagService) GetTagByName(tagName string) (*info.Tags, error) {
	tag := &info.Tags{}
	if _, err := db.MasterDB.Where("tags_name = ?", tagName).Get(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (this *TagService) GetList(limit int, page int, real bool) ([]info.Tags, error) {

	limit = utils.IntDefault(limit, 20)
	page = utils.IntDefault(page, 1)
	start := (page - 1) * limit
	tagsList := make([]info.Tags, 0)
	err := db.MasterDB.Cols("id").Limit(limit, start).Find(&tagsList)
	if err != nil {
		revel.INFO.Printf("Get tag failed : %s", err)
		return nil, err
	}

	idArr := make([]int, len(tagsList))
	for k, v := range tagsList {
		idArr[k] = v.Id
	}

	tagsList, tErr := this.GetTagByIdArr(idArr, real)
	if tErr != nil {
		revel.INFO.Printf("GetTagByIdArr failed: %s", tErr)
		return nil, err
	}

	return tagsList, nil
}

func (this *TagService) GetListPaging(limit int, page int, real bool) (*info.PagingContent, error) {
	dataList, err := this.GetList(limit, page, real)
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

func (this *TagService) GetListByPostId(postId int, real bool) ([]info.Tags, error) {
	tagsList := make([]info.Tags, 0)
	tagIdArr, err := postTagsService.GetTagIdsByPostId(postId)
	if err != nil {
		revel.INFO.Printf("PostTagsService.GetListByPostId Error : %s", err)
		return nil, err
	}

	tagsList, tErr := this.GetTagByIdArr(tagIdArr, real)
	if tErr != nil {
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
			revel.INFO.Printf("Get tag by tags_name failed : %s", err)
			return false, err
		}

		if newTag.Id <= 0 {
			//Not exist, Insert.
			newTag.TagsName = tagsName
			newTag.TagsFlag = url.QueryEscape(tagsName)
			if _, insertErr := db.MasterDB.Insert(&newTag); insertErr != nil {
				revel.INFO.Printf("Get tag by tags_name failed : %s", insertErr)
				continue
			}
			this.GetTagById(newTag.Id, true)
		}

		//construct post tags map[]
		postTags[index].PostsId = postId
		postTags[index].TagsId = newTag.Id
	}

	//Insert post tags array.
	postTagsService.Save(postId, postTags)

	return true, nil
}

func (this *TagService) SaveOne(tag info.Tags) (int, error) {
	if _, err := db.MasterDB.InsertOne(tag); err != nil {
		//revel.INFO.Printf("Get tag by tags_name failed : %s", err)
		return 0, err
	}

	this.GetTagById(tag.Id, true)
	return tag.Id, nil
}

func (this *TagService) Update(id int, tag info.Tags) (bool, error) {
	if _, err := db.MasterDB.Id(id).Update(tag); err != nil {
		//revel.INFO.Printf("Update tag failed: %s", err)
		return false, err
	}

	this.GetTagById(id, true)
	return true, nil
}

func (this *TagService) Destroy(id int, tag info.Tags) (bool, error) {
	if _, err := db.MasterDB.Id(id).Delete(tag); err != nil {
		//revel.INFO.Printf("Destroy tag failed: %s", err)
		return false, err
	}

	cacheKey := utils.CacheKey("TagService", "InfoById", id)
	go cache.Delete(cacheKey)
	return true, nil
}

func (this *TagService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
