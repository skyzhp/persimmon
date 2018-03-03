package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
)

type LinkService struct{}

func (this *LinkService) GetOne(id int) *info.Links {
	option := &info.Links{Id: id}
	_, err := db.MasterDB.Get(option)
	if err != nil {
		revel.INFO.Printf("Get link fatal : %s", err)
	}
	return option
}

func (this *LinkService) GetList(page int) []info.Links {
	limit := 20
	start := (page - 1) * limit
	optionList := make([]info.Links, 0)
	err := db.MasterDB.Limit(limit, start).Find(&optionList)
	if err != nil {
		revel.INFO.Printf("Get link fatal : %s", err)
	}
	return optionList
}

func (this *LinkService) Save(option info.Links) int64 {
	lastInsertId, err := db.MasterDB.InsertOne(option)
	if err != nil {
		return int64(0)
	}
	return lastInsertId
}

func (this *LinkService) Update(id int, option info.Links) bool {
	_, err := db.MasterDB.Id(id).Update(option)
	if err != nil {
		revel.INFO.Printf("Update link error: %s", err)
		return false
	}
	return true
}

func (this *LinkService) Destroy(id int, option info.Links) bool {
	_, err := db.MasterDB.Id(id).Delete(option)
	if err != nil {
		revel.INFO.Printf("Destroy link error: %s", err)
		return false
	}
	return true
}

func (this *LinkService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
