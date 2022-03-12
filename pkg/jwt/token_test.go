package jwt

import (
	"testing"
	"time"
)

const (
	app = "macos"
	key = "182373643982786"
	secret = "CJf938D88fgdD3774ccFic37f"
)

func TestGenerateToken(t *testing.T) {
	token := New(app, key, secret)
	var duration = 86400 * time.Second
	tokenString, err := token.GenerateToken("153662", duration)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tokenString)
}

func TestParseToken(t *testing.T) {
	token := New(app, key, secret)
	claims, err := token.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTUzNjYyIiwiaXNzIjoiMTgyMzczNjQzOTgyNzg2Iiwic3ViIjoibWFjb3MiLCJleHAiOjE2NDcxNDc1NjMsIm5iZiI6MTY0NzA2MTE2MywiaWF0IjoxNjQ3MDYxMTYzfQ.7ksobgZHI2FB9dQT5Zk9b2tfdnuAuA_Odx-SnRiXNe4")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(claims)
}