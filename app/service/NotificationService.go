package service

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/utils"
	"html/template"
	"net/smtp"
	"strings"
	"bytes"
	"fmt"
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

	tplFile := "app/views/mails/comments.html"
	comment, cmErr := commentService.GetCommentById(commentId, false)
	if cmErr != nil {
		revel.INFO.Printf("GetCommentById Error: %s", cmErr)
		return false, cmErr
	}

	subject := fmt.Sprintf("“%s” 有新的评论", comment.Title)
	t, pErr := template.ParseFiles(tplFile)
	if pErr != nil {
		revel.INFO.Printf("Template ParseFiles Error: %s", pErr)
		return false, pErr
	}

	data := map[string]interface{}{
		"Title":     comment.Title,
		"Url":       comment.Url,
		"Name":      comment.Name,
		"Content":   template.HTML(comment.Content),
		"CreatedAt": utils.Date("2006-01-02 15:04:05", comment.CreatedAt),
		"Host":      host}

	var tpl bytes.Buffer
	if exeErr := t.Execute(&tpl, data); exeErr != nil {
		revel.INFO.Printf("Template Execute Error: %s", exeErr)
		return false, exeErr
	}

	body := tpl.String()
	ret, err := this.SendMail(subject, body)
	if err != nil {
		return false, err
	}

	return ret, nil
}
