package mail

import "testing"

func TestSendTextHtml(t *testing.T) {
	m := New(
		WithMailHost("smtp.qq.com"),
		WithMailPort(465),
		WithMailUser("xxxxxx@qq.com"),
		WithMailPass("xxxxxx"),
		)
	err := m.SendTextHtml("测试", "我是测试邮箱发送", []string{"xxxxxx@qq.com"})

	if err != nil {
		t.Fatal(err)
	}

	t.Log("SUCCESS")
}
