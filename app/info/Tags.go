package info

type Tags struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	TagsName  string `json:"tags_name"`
	TagsFlag  string `json:"tags_flag"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"created"`
}
