package auth

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
	token := NewJWT(app, key, secret)
	var duration = 86400 * time.Second
	tokenString, err := token.GenerateToken("153662", duration, JWTClaimsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tokenString)
}

func TestParseToken(t *testing.T) {
	token := NewJWT(app, key, secret)
	claims, err := token.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1MzY2MiIsImlzcyI6IjE4MjM3MzY0Mzk4Mjc4NiIsInN1YiI6Im1hY29zIiwiZXhwIjoxNjQ4OTUxODc0LCJuYmYiOjE2NDg4NjU0NzQsImlhdCI6MTY0ODg2NTQ3NH0.WDes8MQ6u0EnUm0xTEZnxZrYayadzNK8mGOzuR0nB0I")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(claims)
}