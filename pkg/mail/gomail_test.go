package mail

import "testing"

func TestSendMail(t *testing.T) {
	err := Send(&Options{
		MailHost: "smtp.163.com",
		MailPort: 465,
		MailUser: "xxxxxx@163.com",
		MailPass: "xxxxxx",
		MailTo: []string{
			"xxxxxx@qq.com",
		},
		Subject: "测试",
		Body: "我是测试邮箱发送",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log("SUCCESS")
}
