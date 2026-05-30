// Package mail provides SMTP email sending utilities via gomail.
// It supports HTML email content and multiple recipients.
// Supports timeout configuration for connection and sending operations.
package mail

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

var _ Mail = (*mail)(nil)

var (
	ErrHostEmpty      = fmt.Errorf("mail server address cannot be empty")
	ErrPortInvalid    = fmt.Errorf("mail port must be in range 1-65535")
	ErrUserEmpty      = fmt.Errorf("mail username cannot be empty")
	ErrToEmpty        = fmt.Errorf("mail recipient cannot be empty")
	ErrSubjectEmpty   = fmt.Errorf("mail subject cannot be empty")
	ErrTimeoutInvalid = fmt.Errorf("timeout must be greater than 0")
	ErrSendTimeout    = fmt.Errorf("mail send timeout")
	ErrConnectTimeout = fmt.Errorf("mail connection timeout")
)

type Mail interface {
	Message(subject string, to []string) *gomail.Message
	SendTextHtml(subject, body string, to []string) error
	SendTextHtmlWithContext(ctx context.Context, subject, body string, to []string) error
	Close() error
}

type Option func(*option)

type option struct {
	host               string
	port               int
	user               string
	pass               string
	timeout            time.Duration
	insecureSkipVerify bool
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

func WithMailTimeout(timeout time.Duration) Option {
	return func(o *option) {
		o.timeout = timeout
	}
}

func WithMailInsecureSkipVerify(skip bool) Option {
	return func(o *option) {
		o.insecureSkipVerify = skip
	}
}

func New(opts ...Option) (Mail, error) {
	var m = new(mail)
	var o = &option{
		timeout: 30 * time.Second,
	}
	for _, opt := range opts {
		opt(o)
	}

	if o.host == "" {
		return nil, ErrHostEmpty
	}
	if o.port <= 0 || o.port > 65535 {
		return nil, ErrPortInvalid
	}
	if o.user == "" {
		return nil, ErrUserEmpty
	}
	if o.timeout <= 0 {
		return nil, ErrTimeoutInvalid
	}

	m.option = o
	return m, nil
}

type mail struct {
	*option
	conn net.Conn
}

func (m *mail) message(subject string, to []string) *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", m.user)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	return message
}

func (m *mail) Message(subject string, to []string) *gomail.Message {
	return m.message(subject, to)
}

func (m *mail) SendTextHtml(subject, body string, to []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	return m.SendTextHtmlWithContext(ctx, subject, body, to)
}

func (m *mail) SendTextHtmlWithContext(ctx context.Context, subject, body string, to []string) error {
	if len(to) == 0 {
		return ErrToEmpty
	}
	if subject == "" {
		return ErrSubjectEmpty
	}

	return m.sendWithContext(ctx, subject, body, to)
}

func (m *mail) sendWithContext(ctx context.Context, subject, body string, to []string) error {
	addr := fmt.Sprintf("%s:%d", m.host, m.port)

	var d net.Dialer
	d.Timeout = m.timeout

	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("%w: %v", ErrConnectTimeout, err)
		}
		return fmt.Errorf("failed to connect to mail server: %w", err)
	}
	m.conn = conn

	client, err := smtp.NewClient(conn, m.host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}

	if err = client.StartTLS(&tls.Config{
		ServerName:         m.host,
		InsecureSkipVerify: m.insecureSkipVerify,
	}); err != nil {
		if !strings.Contains(err.Error(), "already using TLS") {
			return fmt.Errorf("failed to start TLS: %w", err)
		}
	}

	auth := smtp.PlainAuth("", m.user, m.pass, m.host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %w", err)
	}

	if err = client.Mail(m.user); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			return fmt.Errorf("failed to set recipient [%s]: %w", recipient, err)
		}
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to prepare mail data: %w", err)
	}

	headers := fmt.Sprintf("Subject: %s\r\nFrom: %s\r\nTo: %s\r\nMIME-version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n",
		subject, m.user, strings.Join(to, ","))
	_, err = w.Write([]byte(headers))
	if err != nil {
		return fmt.Errorf("failed to write mail headers: %w", err)
	}

	_, err = w.Write([]byte(body))
	if err != nil {
		return fmt.Errorf("failed to write mail body: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close mail data writer: %w", err)
	}

	err = client.Quit()
	if err != nil {
		return fmt.Errorf("failed to close SMTP connection: %w", err)
	}

	return nil
}

func (m *mail) Close() error {
	if m.conn != nil {
		return m.conn.Close()
	}
	return nil
}
