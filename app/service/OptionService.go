package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
)

type OptionService struct{}

func (this *OptionService) GetOne(id int) *info.Options {
	option := &info.Options{Id: id}
	_, err := db.MasterDB.Get(option)
	if err != nil {
		revel.INFO.Printf("Get option fatal : %s", err)
	}
	return option
}

func (this *OptionService) GetList(page int) []info.Options {
	limit := 20
	start := (page - 1) * limit
	optionList := make([]info.Options, 0)
	err := db.MasterDB.Limit(limit, start).Find(&optionList)
	if err != nil {
		revel.INFO.Printf("Get option fatal : %s", err)
	}
	return optionList
}

func (this *OptionService) Save(option info.Options) int64 {
	lastInsertId, err := db.MasterDB.InsertOne(option)
	if err != nil {
		return int64(0)
	}
	return lastInsertId
}

func (this *OptionService) Update(id int, option info.Options) bool {
	_, err := db.MasterDB.Id(id).Update(option)
	if err != nil {
		revel.INFO.Printf("Update option error: %s", err)
		return false
	}
	return true
}

func (this *OptionService) Destroy(id int, option info.Options) bool {
	_, err := db.MasterDB.Id(id).Delete(option)
	if err != nil {
		revel.INFO.Printf("Destroy option error: %s", err)
		return false
	}
	return true
}

func (this *OptionService) GetValueByName(optionName string) string {
	option := info.Options{OptionName: optionName}
	_, err := db.MasterDB.Get(option)
	if err != nil {
		revel.INFO.Printf("Get value by name fatal : %s", err)
	}
	return option.OptionValue
}

func (this *OptionService) UpdateByName(optionName string, optionValue string) bool {
	option := info.Options{OptionValue: optionValue}
	_, err := db.MasterDB.Where("option_name=?", optionName).Cols("option_value").Update(option)
	if err != nil {
		revel.INFO.Printf("Update by name option error: %s", err)
		return false
	}
	return true
}

func (this *OptionService) GetAllOption() info.Options {
	option := info.Options{}
	err := db.MasterDB.Where("option_status !='hidden'").Find(option)
	if err != nil {
		revel.INFO.Printf("Update by name option error: %s", err)
	}
	return option
}

func (this *OptionService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
