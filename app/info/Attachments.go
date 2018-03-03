package info

type Attachments struct {
	Id        int      `json:"id" xorm:"pk autoincr"`
	Path      string   `json:"path"`
	UserId    int      `json:"user_id"`
	Hash1     string   `json:"hash1"`
	Md5       string   `json:"md5"`
	Ipaddress string   `json:"ipaddress"`
	CreatedAt JsonTime `json:"created_at" xorm:"created"`
	UpdatedAt JsonTime `json:"updated_at" xorm:"updated"`
}
