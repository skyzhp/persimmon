package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel/cache"
	"github.com/revel/revel"
	"time"
)

type LinkService struct{}

func (this *LinkService) GetOne(id int) (*info.Links, error) {
	link := &info.Links{Id: id}
	_, err := db.MasterDB.Get(link)
	if err != nil {
		//revel.INFO.Printf("Get link failed : %s", err)
		return nil, err
	}
	return link, nil
}

func (this *LinkService) GetList(limit int, page int, real bool) ([]info.Links, error) {
	limit = utils.IntDefault(limit, 20)
	page = utils.IntDefault(page, 1)
	start := (page - 1) * limit
	linkList := make([]info.Links, 0)

	cacheKey := utils.CacheKey("LinkService", "Links")
	if err := cache.Get(cacheKey, &linkList); err != nil || real {
		err := db.MasterDB.Limit(limit, start).Find(&linkList)
		if err != nil {
			revel.INFO.Printf("Get link failed : %s", err)
			return nil, err
		}
		go cache.Set(cacheKey, linkList, 30*time.Minute)
	}

	return linkList, nil
}

func (this *LinkService) GetListPaging(limit int, page int) (*info.PagingContent, error) {
	dataList, err := this.GetList(limit, page, false)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountLinks()
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

func (this *LinkService) CountLinks() (int, error) {
	link := new(info.Links)
	total, err := db.MasterDB.Count(link)
	if err != nil {
		//revel.INFO.Printf("Count link failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *LinkService) Save(link info.Links) (int, error) {
	if _, err := db.MasterDB.InsertOne(link); err != nil {
		//revel.INFO.Printf("Save link failed : %s", err)
		return 0, err
	}
	return link.Id, nil
}

func (this *LinkService) Update(id int, link info.Links) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(link)
	if err != nil {
		//revel.INFO.Printf("Update link failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *LinkService) Destroy(id int, link info.Links) (bool, error) {
	_, err := db.MasterDB.Id(id).Delete(link)
	if err != nil {
		//revel.INFO.Printf("Destroy link failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *LinkService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
