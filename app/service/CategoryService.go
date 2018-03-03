package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
)

type CategoryService struct{}

func (this *CategoryService) GetOne(id int) *info.Categorys {
	category := &info.Categorys{Id: id}
	_, err := db.MasterDB.Get(category)
	if err != nil {
		revel.INFO.Printf("Get category fatal : %s", err)
	}
	return category
}

func (this *CategoryService) GetList(page int) []info.Categorys {
	limit := 20
	start := (page - 1) * limit
	categoryList := make([]info.Categorys, 0)
	err := db.MasterDB.Limit(limit, start).Find(&categoryList)
	if err != nil {
		revel.INFO.Printf("Get categorys list fatal : %s", err)
	}
	return categoryList
}

func (this *CategoryService) Save(category info.Categorys) int64 {
	lastInsertId, err := db.MasterDB.InsertOne(category)
	if err != nil {
		return int64(0)
	}
	return lastInsertId
}

func (this *CategoryService) Update(id int, category info.Categorys) bool {
	_, err := db.MasterDB.Id(id).Update(category)
	if err != nil {
		revel.INFO.Printf("Update category error: %s", err)
		return false
	}
	return true
}

func (this *CategoryService) Destroy(id int, category info.Categorys) bool {
	_, err := db.MasterDB.Id(id).Delete(category)
	if err != nil {
		revel.INFO.Printf("Destroy category error: %s", err)
		return false
	}
	return true
}

func (this *CategoryService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
