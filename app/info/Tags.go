package info

type Tags struct {
	Id        int      `json:"id" xorm:"pk autoincr"`
	TagsName  string   `json:"tags_name"`
	TagsFlag  string   `json:"tags_flag"`
	CreatedAt JsonTime `json:"created_at" xorm:"created"`
	UpdatedAt JsonTime `json:"updated_at" xorm:"created"`
}
