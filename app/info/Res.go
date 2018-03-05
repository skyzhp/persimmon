package info

type Res struct {
	Status int         `json:"status"`
	Info   interface{} `json:"info"`
	List   interface{} `json:"list"`
	Item   interface{} `json:"item"`
}