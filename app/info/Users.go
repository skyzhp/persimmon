package info

type Users struct {
	Id        int      `json:"id" xorm:"pk autoincr"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Avatar    string   `json:"avatar"`
	CreatedAt JsonTime `json:"created_at" xorm:"created"`
	UpdatedAt JsonTime `json:"updated_at" xorm:"updated"`
}
