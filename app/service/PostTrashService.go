package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel/cache"
	"github.com/revel/revel"
	"fmt"
	"qiniupkg.com/x/errors.v7"
)

type PostTrashService struct{}

func (this *PostTrashService) GetTrashList(categoryId int, keywords string, limit int, page int, real bool) ([]info.Posts, error) {

	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	start := (page - 1) * limit
	postList := make([]info.Posts, 0)
	if err := db.MasterDB.Where("`deleted_at` IS NOT NULL").Cols("id").Limit(limit, start).Find(&postList); err != nil {
		revel.INFO.Printf("Get trash list failed : %s", err)
		return nil, err
	}

	idArr := make([]int, len(postList))
	for k, v := range postList {
		idArr[k] = v.Id
	}

	if len(idArr) <= 0 {
		return make([]info.Posts, 0), nil
	}

	postList, pErr := postService.GetPostByIdArr(idArr, real)
	if pErr != nil {
		revel.INFO.Printf("postService.GetPostByIdArr Error : %s", pErr)
		return nil, pErr
	}

	return postList, nil
}

func (this *PostTrashService) GetTrashListPaging(categoryId int, keywords string, limit int, page int, real bool) (*info.PagingContent, error) {
	dataList, err := this.GetTrashList(categoryId, keywords, limit, page, real)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountTrashPost()
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

func (this *PostTrashService) Restore(ids []int) (bool, error) {
	if len(ids) < 1 {
		return false, errors.New("params error.")
	}
	idString := utils.IntJoin(ids, ",")
	postsTable := this.Table("posts")
	sql := fmt.Sprintf("UPDATE %s SET deleted_at = NULL WHERE id in (%s)", postsTable, idString)

	if _, err := db.MasterDB.Exec(sql); err != nil {
		revel.INFO.Printf("Restore post failed: %s", err)
		return false, err
	}
	return true, nil
}

//real delete
func (this *PostTrashService) Destroy(ids []int) (bool, error) {
	if len(ids) < 1 {
		return false, errors.New("params error.")
	}
	if _, err := db.MasterDB.In("id", ids).Delete(&info.Posts{}); err != nil {
		revel.INFO.Printf("Destroy post failed: %s", err)
		return false, err
	}

	for _, id := range ids {
		cKey := utils.CacheKey("PostService", "InfoById", id)
		go cache.Delete(cKey)
	}

	return true, nil
}

func (this *PostTrashService) CountTrashPost() (int, error) {
	post := new(info.Posts)
	total, err := db.MasterDB.Where("deleted_at IS NULL OR `deleted_at` = '0001-01-01 00:00:00'").Count(post)
	if err != nil {
		revel.INFO.Printf("CountTrashPost Error: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *PostTrashService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
