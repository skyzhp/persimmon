package info

type Res struct {
	Ok   bool        `json:"ok"`
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	List interface{} `json:"list"`
	Item interface{} `json:"item"`
}

func NewRes(Ok bool, msg string) Res {
	return Res{Ok: Ok, Code: 200, Msg: msg}
}
