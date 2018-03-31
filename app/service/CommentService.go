package service

import (
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/utils"
	"errors"
	"time"
)

type CommentService struct{}

func (this *CommentService) GetCommentById(id int, real bool) (info.Comments, error) {
	comments := info.Comments{}
	cacheKey := utils.CacheKey("CommentService", "InfoById", id)
	if err := cache.Get(cacheKey, &comments); err != nil || real {
		_, err := db.MasterDB.Where("id = ?", id).Get(&comments)
		if err != nil {
			return comments, err
		}
		go cache.Set(cacheKey, comments, 30*time.Minute)
	}

	//post title
	post, _ := postService.GetPostById(comments.PostsId, false)
	comments.Title = post.Title

	return comments, nil
}

func (this *CommentService) GetCommentByIdArr(idArr []int, real bool) ([]info.Comments, error) {
	idLen := len(idArr)
	if idLen <= 0 {
		return nil, errors.New("参数不正确")
	}

	commentArr := make([]info.Comments, idLen)

	for k, commentId := range idArr {
		comment, err := this.GetCommentById(commentId, real)
		if err == nil {
			commentArr[k] = comment
		}
	}

	return commentArr, nil
}

func (this *CommentService) GetList(postId int, limit int, page int, real bool) ([]info.Comments, error) {
	limit = utils.IntDefault(limit, 20)
	page = utils.IntDefault(page, 1)
	postId = utils.IntDefault(postId, 0)
	start := (page - 1) * limit
	commentIdArr := make([]info.Comments, 0)
	dbSession := db.MasterDB.NewSession()

	if postId > 0 {
		dbSession.Where("posts_id = ?", postId)
	}

	err := dbSession.Cols("id").Limit(start, limit).Find(&commentIdArr)
	if err != nil {
		revel.INFO.Printf("Get comment failed : %s", err)
		return nil, err
	}

	idArr := make([]int, len(commentIdArr))
	for k, v := range commentIdArr {
		idArr[k] = v.Id
	}

	commentArr, cErr := this.GetCommentByIdArr(idArr, false)
	if cErr != nil {
		revel.INFO.Printf("GetCommentByIdArr Error : %s", cErr)
		return nil, cErr
	}

	return commentArr, nil
}

func (this *CommentService) GetListPaging(limit int, page int, real bool) (*info.PagingContent, error) {
	dataList, err := this.GetList(0, limit, page, real)
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

func (this *CommentService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
