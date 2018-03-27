package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"time"
	"github.com/cong5/persimmon/app/utils"
	"errors"
)

type CategoryService struct {
	page  int
	limit int
}

func (this *CategoryService) GetOne(id int) (*info.Categorys, error) {
	category := &info.Categorys{Id: id}
	_, err := db.MasterDB.Get(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (this *CategoryService) GetOneBySlug(slug string) (*info.Categorys, error) {
	if slug == "" {
		return nil, errors.New("Slug is empty.")
	}
	category := &info.Categorys{}
	_, err := db.MasterDB.Where("category_flag = ?", slug).Get(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (this *CategoryService) GetList(limit int, page int) ([]info.Categorys, error) {
	this.limit = utils.IntDefault(limit, 20)
	this.page = utils.IntDefault(page, 20)

	start := (this.page - 1) * limit
	categorysList := make([]info.Categorys, 0)

	if err := db.MasterDB.Limit(this.limit, start).Find(&categorysList); err != nil {
		return nil, err
	}

	return categorysList, nil
}

func (this *CategoryService) GetListPaging(limit int, page int) (*info.PagingContent, error) {
	dataList, err := this.GetList(limit, page)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountCategory()
	if cErr != nil {
		return nil, err
	}

	totalPage := utils.GetTotalPage(total, this.limit)
	pagingContent := &info.PagingContent{Data: dataList,
		Total: total,
		TotalPage: totalPage,
		CurrentPage: this.page}
	return pagingContent, nil
}

func (this *CategoryService) CountCategory() (int, error) {
	category := new(info.Categorys)
	total, err := db.MasterDB.Count(category)
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (this *CategoryService) Save(category info.Categorys) (int, error) {
	if _, err := db.MasterDB.InsertOne(category); err != nil {
		//revel.INFO.Printf("Save categorys failed : %s", err)
		return 0, err
	}
	return category.Id, nil
}

func (this *CategoryService) Update(id int, category info.Categorys) (bool, error) {
	if _, err := db.MasterDB.Id(id).Update(category); err != nil {
		return false, err
	}
	return true, nil
}

func (this *CategoryService) Destroy(ids []int, category info.Categorys) (bool, error) {
	if _, err := db.MasterDB.In("id", ids).Delete(category); err != nil {
		return false, err
	}
	return true, nil
}

func (this *CategoryService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
