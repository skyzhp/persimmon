package info

type Comments struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	ParentId  int    `json:"parent_id"`
	PostsId   int    `json:"posts_id"`
	Title     string `xorm:"-"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Md5       string `json:"md5" xorm:"-"`
	Url       string `json:"url"`
	Content   string `json:"content"`
	Markdown  string `json:"markdown"`
	ClientIp  uint32 `json:"client_ip"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"updated"`
}