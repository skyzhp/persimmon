package info

type Comments struct {
	Id        int      `json:"id" xorm:"pk autoincr"`
	ParentId  int      `json:"parent_id"`
	PostsId   int      `json:"posts_id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Md5       string   `json:"md5" xorm:"-"`
	Url       string   `json:"url"`
	Content   string   `json:"content"`
	Markdown  string   `json:"markdown"`
	Ipaddress string   `json:"ipaddress"`
	CreatedAt JsonTime `json:"created_at" xorm:"created"`
	UpdatedAt JsonTime `json:"updated_at" xorm:"updated"`
}

type JoinComments struct {
	Id        int      `json:"id" xorm:"pk autoincr"`
	ParentId  int      `json:"parent_id"`
	Title     string   `json:"title"`
	PostsId   int      `json:"posts_id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Url       string   `json:"url"`
	Content   string   `json:"content"`
	Markdown  string   `json:"markdown"`
	Ipaddress string   `json:"ipaddress"`
	CreatedAt JsonTime `json:"created_at" xorm:"created"`
	UpdatedAt JsonTime `json:"updated_at" xorm:"updated"`
}
