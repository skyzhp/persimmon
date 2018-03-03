package info

type Links struct {
	Id        int      `json:"id" xorm:"pk autoincr"`
	Name      string   `json:"name"`
	Logo      string   `json:"logo"`
	Group     string   `json:"group"`
	Url       string   `json:"url"`
	Ipaddress string   `json:"ipaddress"`
	CreatedAt JsonTime `json:"created_at" xorm:"created"`
	UpdatedAt JsonTime `json:"updated_at" xorm:"updated"`
}
