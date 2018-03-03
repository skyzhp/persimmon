package info

type Comments struct {
	Id        int      `json:"id" xorm:"pk autoincr"`
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
