package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
	"time"
	"github.com/cong5/persimmon/app/utils"
	"fmt"
)

type CommentService struct{}

func (this *CommentService) GetOne(id int) (*info.Comments, error) {
	comment := &info.Comments{Id: id}
	_, err := db.MasterDB.Get(comment)
	if err != nil {
		//revel.INFO.Printf("Get comment failed : %s", err)
		return nil, err
	}
	return comment, nil
}

func (this *CommentService) GetList(limit int, page int) ([]info.JoinComments, error) {
	limit = utils.IntDefault(limit, 20)
	page = utils.IntDefault(page, 1)

	start := (page - 1) * limit
	commentList := make([]info.JoinComments, 0)
	commentTable := db.MasterDB.TableMapper.Obj2Table("comments")
	postTable := db.MasterDB.TableMapper.Obj2Table("posts")
	joinWhere := fmt.Sprintf("%s.posts_id = %s.id", commentTable, postTable)
	commentColumns := fmt.Sprintf("%s.*", commentTable)
	postColumns := fmt.Sprintf("%s.title", postTable)
	err := db.MasterDB.Table(commentTable).Join("LEFT OUTER", postTable, joinWhere).Limit(limit, start).Cols(commentColumns, postColumns).Find(&commentList)
	if err != nil {
		//revel.INFO.Printf("Get comment failed : %s", err)
		return nil, err
	}

	return commentList, nil
}

func (this *CommentService) GetListPaging(limit int, page int) (*info.PagingContent, error) {
	dataList, err := this.GetList(limit, page)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountComment(0)
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

func (this *CommentService) GetCommentByPostId(postId int, limit int, page int) (*info.PagingContent, error) {
	limit = utils.IntDefault(limit, 20)
	page = utils.IntDefault(page, 1)

	comments := make([]info.Comments, 0)
	start := (page - 1) * limit
	err := db.MasterDB.Where("posts_id = ?", postId).Limit(limit, start).Find(&comments)
	if err != nil {
		//revel.INFO.Printf("Get comment by post id failed : %s", err)
		return nil, err
	}

	//对email进行Md5签名,内容@解析
	comments = utils.CommentRelolver(comments)

	total, cErr := this.CountComment(postId)
	if cErr != nil {
		return nil, cErr
	}

	totalPage := utils.GetTotalPage(total, limit)

	pagingContent := &info.PagingContent{Data: comments,
		Total: total,
		TotalPage: totalPage,
		CurrentPage: page}
	return pagingContent, nil
}

func (this *CommentService) Save(comment info.Comments) (int, error) {
	if _, err := db.MasterDB.InsertOne(&comment); err != nil {
		revel.INFO.Printf("Save comment failed : %s", err)
		return 0, err
	}
	return comment.Id, nil
}

func (this *CommentService) Update(id int, comment info.Comments) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(comment)
	if err != nil {
		//revel.INFO.Printf("Update comment failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *CommentService) Destroy(id int, comment info.Comments) (bool, error) {
	_, err := db.MasterDB.Id(id).Delete(comment)
	if err != nil {
		//revel.INFO.Printf("Destroy comment failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *CommentService) CountComment(postId int) (int, error) {
	comment := new(info.Comments)
	dbSession := db.MasterDB.NewSession()
	if postId > 0 {
		dbSession.Where("posts_id = ?", postId)
	}
	total, err := dbSession.Count(comment)
	if err != nil {
		//revel.INFO.Printf("Count comment failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *CommentService) GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
