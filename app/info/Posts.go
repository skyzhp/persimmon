package info

type Posts struct {
	Id         int       `json:"id" xorm:"pk autoincr"`
	Flag       string    `json:"flag"`
	Title      string    `json:"title"`
	Thumb      string    `json:"thumb"`
	CategoryId int       `json:"categoryId"`
	Categories Categorys `json:"categories" xorm:"-"`
	UserId     int       `json:"user_id"`
	Content    string    `json:"content"`
	Markdown   string    `json:"markdown"`
	Views      int       `json:"views"`
	Comments   int       `json:"comments"`
	Ipaddress  string    `json:"ipaddress"`
	Tags       []Tags    `json:"tags" xorm:"-"`
	DeletedAt  JsonTime  `json:"deleted_at" xorm:"deleted"`
	CreatedAt  JsonTime  `json:"created_at" xorm:"created"`
	UpdatedAt  JsonTime  `json:"updated_at" xorm:"updated"`
}

type Post struct {
	Id         int
	Flag       string
	Title      string
	Thumb      string
	CategoryId int
	UserId     int
	Content    string
	Markdown   string
	Tags       []string
}
