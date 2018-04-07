package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel/cache"
	"github.com/revel/revel"
	"errors"
	"time"
)

type LinkService struct{}

func (this *LinkService) GetLinkById(id int, real bool) (info.Links, error) {
	link := info.Links{}

	if id <= 0 {
		return link, errors.New("Param error.")
	}

	cKey := utils.CacheKey("LinkService", "InfoById", id)
	if real {
		go cache.Delete(cKey)
	}

	if err := cache.Get(cKey, &link); err != nil {
		_, err := db.MasterDB.Where("id = ?", id).Get(&link)
		if err != nil {
			revel.INFO.Printf("Get link failed : %s", err)
			return link, err
		}

		go cache.Set(cKey, link, 30*time.Minute)
	}

	return link, nil
}

func (this *LinkService) GetLinksByIdArr(idArr []int, real bool) ([]info.Links, error) {
	idLen := len(idArr)
	if idLen <= 0 {
		return nil, errors.New("参数不正确")
	}

	linkArr := make([]info.Links, idLen)

	for k, postId := range idArr {
		link, err := this.GetLinkById(postId, real)
		if err == nil {
			linkArr[k] = link
		}
	}

	return linkArr, nil
}

func (this *LinkService) GetList(limit int, page int, real bool) ([]info.Links, error) {
	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	start := (page - 1) * limit
	linkIdArr := make([]info.Links, 0)

	err := db.MasterDB.Cols("id").Limit(limit, start).Find(&linkIdArr)
	if err != nil {
		revel.INFO.Printf("Get link failed : %s", err)
		return nil, err
	}

	linkLen := len(linkIdArr)
	if linkLen <= 0 {
		return make([]info.Links, 0), nil
	}

	idArr := make([]int, linkLen)
	for k, v := range linkIdArr {
		idArr[k] = v.Id
	}

	linkArr, lErr := this.GetLinksByIdArr(idArr, real)
	if lErr != nil {
		revel.INFO.Printf("GetLinksByIdArr : %s", lErr)
		return nil, err
	}

	return linkArr, nil
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

	this.GetLinkById(link.Id, true)
	return link.Id, nil
}

func (this *LinkService) Update(id int, link info.Links) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(link)
	if err != nil {
		//revel.INFO.Printf("Update link failed: %s", err)
		return false, err
	}

	this.GetLinkById(id, true)
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
