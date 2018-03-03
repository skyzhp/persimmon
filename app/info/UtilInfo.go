package info

import "time"

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format("2006-01-02 15:04:05") + `"`), nil
}

type BaiduFanyi struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	From      string `json:"from"`
	To        string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

type CommentMail struct {
	Title     string
	Content   string
	Url       string
	CreatedAt string
}
