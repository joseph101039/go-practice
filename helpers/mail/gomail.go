package mail

import (
	"errors"
	"goroutine/helpers/env"

	"gopkg.in/gomail.v2"
)

const (
	htmlType = "text/html; charset=utf-8"
	textType = "text/plain; charset=utf-8"
)

// Gmail require env MAIL_USERNAME, MAIL_PASSWORD
// example:
// var gmail Gmail
// g := (&gmail).New()
// err := g.SetSubject(). ...  . Send()
type Gmail struct {
	subject       string   // 主題
	content       string   // 內容
	contentType   string   // 內容類型
	recipients    []string // 	收件者
	sender        string   // 	寄件者
	attachedFiles []string
	dialer        *gomail.Dialer
}

func (g *Gmail) New() error {
	g.contentType = htmlType // default

	username := env.Get("MAIL_USERNAME")
	g.sender = username
	if username == "" {
		return errors.New("env MAIL_USERNAME is not set")
	}

	password := env.Get("MAIL_PASSWORD")
	if password == "" {
		return errors.New("env MAIL_PASSWORD is not set")
	}

	g.dialer = gomail.NewDialer("smtp.gmail.com", 587, username, password)
	return nil
}

// Send 送出郵件
func (g *Gmail) Send() error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", g.sender)
	if len(g.recipients) == 0 {
		return errors.New("無指定收件者, 請呼叫 setRecipients")
	}
	msg.SetHeader("To", g.recipients...)
	msg.SetHeader("Subject", g.subject)
	msg.SetBody(g.contentType, g.content)

	for _, file := range g.attachedFiles {
		msg.Attach(file)
	}

	// Send the email to Bob, Cora and Dan.
	if err := g.dialer.DialAndSend(msg); err != nil {
		panic(err)
	}

	return nil
}

// SetRecipients 設定收件者
func (g *Gmail) SetRecipients(recipients []string) *Gmail {
	g.recipients = recipients
	return g
}

func (g *Gmail) SetSubject(subject string) *Gmail {
	g.subject = subject
	return g
}

func (g *Gmail) SetContent(content string) *Gmail {
	g.content = content
	return g
}

// SetHtmlContentType 內容為 html
func (g *Gmail) SetHtmlContentType() *Gmail {
	g.contentType = htmlType
	return g
}

// SetTextContentType 內容為存純文字
func (g *Gmail) SetTextContentType() *Gmail {
	g.contentType = textType
	return g
}

func (g *Gmail) AttacheFiles(files []string) *Gmail {
	g.attachedFiles = files
	return g
}
