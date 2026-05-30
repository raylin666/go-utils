package auth

import (
	"testing"
	"time"
)

const (
	app    = "macos"
	key    = "182373643982786"
	secret = "test-secret-key-for-unit-test-only-32b"
)

func TestGenerateToken(t *testing.T) {
	token, err := NewJWT(app, key, secret)
	if err != nil {
		t.Fatal(err)
	}
	duration := 86400 * time.Second
	tokenString, err := token.GenerateToken("153662", duration)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tokenString)
}

func TestParseToken(t *testing.T) {
	token, err := NewJWT(app, key, secret)
	if err != nil {
		t.Fatal(err)
	}
	duration := 86400 * time.Second
	tokenString, err := token.GenerateToken("153662", duration)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := token.ParseToken(tokenString)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(claims)
	t.Logf("Token ID: %s", claims.ID)
}

func TestGenerateTokenWithAudience(t *testing.T) {
	token, err := NewJWT(app, key, secret)
	if err != nil {
		t.Fatal(err)
	}
	duration := 86400 * time.Second
	tokenString, err := token.GenerateToken("153662", duration, "service-a", "service-b")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tokenString)
}

func BenchmarkGenerateToken(b *testing.B) {
	jwt, err := NewJWT(app, key, secret)
	if err != nil {
		b.Fatal(err)
	}
	duration := 86400 * time.Second
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jwt.GenerateToken("user-123", duration)
	}
}

func BenchmarkParseToken(b *testing.B) {
	jwt, err := NewJWT(app, key, secret)
	if err != nil {
		b.Fatal(err)
	}
	duration := 86400 * time.Second
	tokenString, _ := jwt.GenerateToken("user-123", duration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jwt.ParseToken(tokenString)
	}
}