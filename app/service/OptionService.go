package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
	"github.com/cong5/persimmon/app/utils"
)

type OptionService struct{}

func (this *OptionService) GetOne(id int) (*info.Options, error) {
	option := &info.Options{Id: id}
	_, err := db.MasterDB.Get(option)
	if err != nil {
		//revel.INFO.Printf("Get option failed : %s", err)
		return nil, err
	}
	return option, nil
}

func (this *OptionService) GetList(limit int, page int) ([]info.Options, error) {

	limit = utils.IntDefault(limit, 20)
	page = utils.IntDefault(page, 1)
	start := (page - 1) * limit
	optionList := make([]info.Options, 0)
	if err := db.MasterDB.Limit(limit, start).Find(&optionList); err != nil {
		//revel.INFO.Printf("Get option failed : %s", err)
		return nil, err
	}

	return optionList, nil
}

func (this *OptionService) GetListPaging(limit int, page int) (*info.PagingContent, error) {
	dataList, err := this.GetList(limit, page)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountOption()
	if cErr != nil {
		return nil, err
	}

	totalPage := utils.GetTotalPage(total, limit)
	pagingContent := &info.PagingContent{Data: dataList,
		Total: total,
		TotalPage: totalPage,
		CurrentPage: page}
	return pagingContent, nil
}

func (this *OptionService) CountOption() (int, error) {
	option := new(info.Options)
	total, err := db.MasterDB.Count(option)
	if err != nil {
		revel.INFO.Printf("Count option failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *OptionService) Save(option info.Options) (int, error) {
	if _, err := db.MasterDB.InsertOne(option); err != nil {
		return 0, err
	}
	return option.Id, nil
}

func (this *OptionService) Update(id int, option info.Options) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(option)
	if err != nil {
		//revel.INFO.Printf("Update option failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *OptionService) Destroy(id int, option info.Options) (bool, error) {
	_, err := db.MasterDB.Id(id).Delete(option)
	if err != nil {
		//revel.INFO.Printf("Destroy option failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *OptionService) GetValueByName(optionName string) (string, error) {
	option := &info.Options{}
	_, err := db.MasterDB.Where("option_name = ?", optionName).Get(option)
	if err != nil {
		//revel.INFO.Printf("Get value by name failed : %s", err)
		return "", err
	}
	return option.OptionValue, nil
}

func (this *OptionService) UpdateByName(optionName string, optionValue string) (bool, error) {
	option := info.Options{OptionValue: optionValue}
	_, err := db.MasterDB.Where("option_name=?", optionName).Cols("option_value").Update(option)
	if err != nil {
		//revel.INFO.Printf("Update by name option failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *OptionService) GetAllOption() ([]info.Options, error) {
	options := make([]info.Options, 0)
	err := db.MasterDB.Where("option_status !='hidden'").Find(&options)
	if err != nil {
		//revel.INFO.Printf("Get all option failed: %s", err)
		return nil, err
	}

	return options, nil
}

func (this *OptionService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
