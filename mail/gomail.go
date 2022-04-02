package mail

import (
	"gopkg.in/gomail.v2"
)

var _ Mail = (*mail)(nil)

type Mail interface {
	// 自定义构建消息, 直接提供 gomail.Message 调用
	Message(subject string, to []string) *gomail.Message
	// subject:邮件主题 | body:邮件内容 | to:收件人 多个用,分割
	SendTextHtml(subject, body string, to []string) error
}

type mail struct {
	*option
}

type Option func(*option)

type option struct {
	host string
	port int
	user string   // 发件人
	pass string   // 发件人密码
}

func WithMailHost(host string) Option {
	return func(o *option) {
		o.host = host
	}
}

func WithMailPort(port int) Option {
	return func(o *option) {
		o.port = port
	}
}

func WithMailUser(user string) Option {
	return func(o *option) {
		o.user = user
	}
}

func WithMailPass(pass string) Option {
	return func(o *option) {
		o.pass = pass
	}
}

func New(opts ...Option) Mail {
	var mail = new(mail)
	var o = new(option)
	for _, opt := range opts {
		opt(o)
	}

	mail.option = o
	return mail
}

func (m *mail) message(subject string, to []string) *gomail.Message {
	message := gomail.NewMessage()

	// 设置发件人
	message.SetHeader("From", m.user)

	// 设置发送给多个用户
	message.SetHeader("To", to...)

	// 设置邮件主题
	message.SetHeader("Subject", subject)

	return message
}

func (m *mail) Message(subject string, to []string) *gomail.Message {
	return m.message(subject, to)
}

func (m *mail) send(message *gomail.Message) error {
	dialer := gomail.NewDialer(m.host, m.port, m.user, m.pass)
	return dialer.DialAndSend(message)
}

// SendTextHtml 发送文本HTML邮件
func (m *mail) SendTextHtml(subject, body string, to []string) error  {
	var message = m.message(subject, to)

	// 设置邮件正文
	message.SetBody("text/html", body)

	return m.send(message)
}