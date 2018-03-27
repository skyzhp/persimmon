package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"time"
	"fmt"
	"github.com/cong5/persimmon/app/utils"
)

type PostTrashService struct{}

func (this *PostTrashService) GetTrashList(categoryId int, keywords string, limit int, page int) ([]info.Posts, error) {

	limit = utils.IntDefault(limit, 20)
	page = utils.IntDefault(page, 1)
	start := (page - 1) * limit
	postList := make([]info.Posts, 0)
	if err := db.MasterDB.Unscoped().Limit(limit, start).Find(&postList); err != nil {
		//revel.INFO.Printf("Get trash list failed : %s", err)
		return nil, err
	}

	return this.JoinRelated(postList)
}

func (this *PostTrashService) GetTrashListPaging(categoryId int, keywords string, limit int, page int) (*info.PagingContent, error) {
	dataList, err := this.GetTrashList(categoryId, keywords, limit, page)
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

func (this *PostTrashService) JoinRelated(posts []info.Posts) ([]info.Posts, error) {
	if len(posts) <= 0 {
		return posts, nil
	}
	categoryList, catErr := categoryService.GetList(999, 1)
	if catErr != nil {
		return nil, catErr
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
		}
	}

	return posts, nil
}

func (this *PostTrashService) Restore(ids []int) (bool, error) {
	idString := utils.IntImplode(ids, ",")
	postsTable := db.MasterDB.TableMapper.Obj2Table("posts")
	sql := fmt.Sprintf("UPDATE %s SET updated_at = '0001-01-01 00:00:00' WHERE id in (%s)", postsTable, idString)

	if _, err := db.MasterDB.Exec(sql); err != nil {
		//revel.INFO.Printf("Restore post failed: %s", err)
		return false, err
	}
	return true, nil
}

//real delete
func (this *PostTrashService) Destroy(ids []int) (bool, error) {
	if _, err := db.MasterDB.In("id", ids).Unscoped().Delete(&info.Posts{}); err != nil {
		//revel.INFO.Printf("Destroy post failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *PostTrashService) CountTrashPost() (int, error) {
	post := new(info.Posts)
	total, err := db.MasterDB.Where("deleted_at IS NULL OR `deleted_at` = '0001-01-01 00:00:00'").Count(post)
	if err != nil {
		//revel.INFO.Printf("Count post failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *PostTrashService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
