package info

type PostsTags struct {
	PostsId   int      `json:"posts_id" xorm:"pk autoincr"`
	TagsId    int      `json:"tags_id"`
	CreatedAt JsonTime `json:"created_at" xorm:"created"`
	UpdatedAt JsonTime `json:"updated_at" xorm:"updated"`
}