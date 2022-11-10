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
	claims, err := token.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIxODIzNzM2NDM5ODI3ODYiLCJzdWIiOiJtYWNvcyIsImF1ZCI6WyIxNTM2NjIiXSwiZXhwIjoxNjY4MTI1NDgyLCJuYmYiOjE2NjgwMzkwODIsImlhdCI6MTY2ODAzOTA4MiwianRpIjoiMTUzNjYyIn0.khDIKq0oTPOvOS0rCn9lUtvx-pu0i4_hYf4ThUpQ5gs")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(claims)
}