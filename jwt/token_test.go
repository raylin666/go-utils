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
	var duration time.Duration = 86400000000
	tokenString, err := token.GenerateToken("153662", duration)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tokenString)
}

func TestParseToken(t *testing.T) {
	token := New(app, key, secret)
	claims, err := token.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIxNTM2NjIiLCJleHAiOjE2MzM2NjQ1MjksImlhdCI6MTYzMzY2NDQ0MiwiaXNzIjoiMTgyMzczNjQzOTgyNzg2IiwibmJmIjoxNjMzNjY0NDQyLCJzdWIiOiJtYWNvcyJ9._B8WHbxXB6AtuqkBOvwEeFPSV3j9LuaMgiLVCgRQqbY")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(claims)
}