package mail

import (
	"context"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	_, err := New(
		WithMailHost("smtp.qq.com"),
		WithMailPort(465),
		WithMailUser("xxxxxx@qq.com"),
		WithMailPass("xxxxxx"),
		WithMailTimeout(30*time.Second),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Mail instance created successfully")
}

func TestNewWithInvalidParams(t *testing.T) {
	_, err := New(WithMailHost(""))
	if err != ErrHostEmpty {
		t.Errorf("expected ErrHostEmpty, got %v", err)
	}

	_, err = New(WithMailHost("smtp.test.com"), WithMailPort(0))
	if err != ErrPortInvalid {
		t.Errorf("expected ErrPortInvalid, got %v", err)
	}

	_, err = New(WithMailHost("smtp.test.com"), WithMailPort(465), WithMailUser(""))
	if err != ErrUserEmpty {
		t.Errorf("expected ErrUserEmpty, got %v", err)
	}

	_, err = New(
		WithMailHost("smtp.test.com"),
		WithMailPort(465),
		WithMailUser("test@test.com"),
		WithMailTimeout(-1*time.Second),
	)
	if err != ErrTimeoutInvalid {
		t.Errorf("expected ErrTimeoutInvalid, got %v", err)
	}
}

func TestSendTextHtmlWithEmptyParams(t *testing.T) {
	m, err := New(
		WithMailHost("smtp.qq.com"),
		WithMailPort(465),
		WithMailUser("xxxxxx@qq.com"),
		WithMailPass("xxxxxx"),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = m.SendTextHtml("", "body", []string{"test@test.com"})
	if err != ErrSubjectEmpty {
		t.Errorf("expected ErrSubjectEmpty, got %v", err)
	}

	err = m.SendTextHtml("subject", "body", []string{})
	if err != ErrToEmpty {
		t.Errorf("expected ErrToEmpty, got %v", err)
	}
}

func TestSendTextHtmlWithContext(t *testing.T) {
	m, err := New(
		WithMailHost("smtp.qq.com"),
		WithMailPort(465),
		WithMailUser("xxxxxx@qq.com"),
		WithMailPass("xxxxxx"),
		WithMailTimeout(10*time.Second),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = m.SendTextHtmlWithContext(ctx, "测试", "我是测试邮箱发送", []string{"xxxxxx@qq.com"})
	if err != nil {
		t.Logf("Send error (expected in test env): %v", err)
	}
}

func TestSendTextHtml(t *testing.T) {
	m, err := New(
		WithMailHost("smtp.qq.com"),
		WithMailPort(465),
		WithMailUser("xxxxxx@qq.com"),
		WithMailPass("xxxxxx"),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = m.SendTextHtml("测试", "我是测试邮箱发送", []string{"xxxxxx@qq.com"})
	if err != nil {
		t.Logf("Send error (expected in test env): %v", err)
	}

	t.Log("Test completed")
}

func TestClose(t *testing.T) {
	m, err := New(
		WithMailHost("smtp.qq.com"),
		WithMailPort(465),
		WithMailUser("xxxxxx@qq.com"),
		WithMailPass("xxxxxx"),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = m.Close()
	if err != nil {
		t.Errorf("Close error: %v", err)
	}
}