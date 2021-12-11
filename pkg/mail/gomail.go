package mail

import (
	"gopkg.in/gomail.v2"
)

type Options struct {
	MailHost string
	MailPort int
	MailUser string   // 发件人
	MailPass string   // 发件人密码
	MailTo   []string // 收件人 多个用,分割
	Subject  string   // 邮件主题
	Body     string   // 邮件内容
}

// 发送邮件
func Send(option *Options) error {

	mail := gomail.NewMessage()

	// 设置发件人
	mail.SetHeader("From", option.MailUser)

	// 设置发送给多个用户
	mail.SetHeader("To", option.MailTo...)

	// 设置邮件主题
	mail.SetHeader("Subject", option.Subject)

	// 设置邮件正文
	mail.SetBody("text/html", option.Body)

	dialer := gomail.NewDialer(option.MailHost, option.MailPort, option.MailUser, option.MailPass)

	return dialer.DialAndSend(mail)
}
