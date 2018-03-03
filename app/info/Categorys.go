package info

type Categorys struct {
	Id                  int      `json:"id" xorm:"pk autoincr"`
	CategoryName        string   `json:"category_name"`
	CategoryParent      int      `json:"category_parent"`
	CategoryFlag        string   `json:"category_flag"`
	CategoryDescription string   `json:"category_description"`
	Ipaddress           string   `json:"ipaddress"`
	CreatedAt           JsonTime `json:"created_at" xorm:"created"`
	UpdatedAt           JsonTime `json:"updated_at" xorm:"updated"`
}
