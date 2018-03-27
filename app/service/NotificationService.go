package service

import (
	"github.com/revel/revel"
	"net/smtp"
	"strings"
	"fmt"
	"os"
)

type NotificationService struct{}

func (this *NotificationService) SendMail(subject string, body string) (bool, error) {
	host := revel.Config.StringDefault("email.host", "")
	port := revel.Config.StringDefault("email.port", "")
	account := revel.Config.StringDefault("email.account", "")
	password := revel.Config.StringDefault("email.password", "")
	to := revel.Config.StringDefault("email.to", "")

	auth := smtp.PlainAuth("", account, password, host)
	contentType := "Content-Type: text/html; charset=UTF-8"
	message := []byte("To: " + to + "\r\nFrom: " + account + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	serverAddr := fmt.Sprintf("%s:%s", host, port)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(serverAddr, auth, account, sendTo, message)
	if err != nil {
		revel.INFO.Printf("SendMail Error: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (this *NotificationService) SendCommentNotice(postId int, commentId int, host string) (bool, error) {

	tpl, tplErr := os.Open("app/views/mails/comments.html") // For read access.
	if tplErr != nil {
		return false, tplErr
	}
	data := make([]byte, 2*1024)
	_, readErr := tpl.Read(data)
	if readErr != nil {
		return false, tplErr
	}

	post, err := postService.GetOne(postId)
	if err != nil {
		return false, err
	}

	comment, cmErr := commentService.GetOne(commentId)
	if cmErr != nil {
		return false, cmErr
	}

	createdAt, _ := comment.CreatedAt.MarshalJSON()
	subject := fmt.Sprintf("“%s” 有新的评论", post.Title)
	body := fmt.Sprintf(string(data), subject, comment.Url, comment.Name, comment.Content, string(createdAt), host)
	//revel.INFO.Println(body)

	ret, err := this.SendMail(subject, body)
	if err != nil {
		return ret, err
	}

	return true, nil
}
