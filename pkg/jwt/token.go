package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var _ Token = (*token)(nil)

type Token interface {
	// 生成 Token
	GenerateToken(userId string, expireDuration time.Duration) (tokenString string, err error)
	// 解析 Token
	ParseToken(tokenString string) (*claims, error)
}

type token struct {
	app    string
	key    string
	secret string
}

type claims struct {
	UserID string
	jwt.StandardClaims
}

func New(app string, key string, secret string) Token {
	return &token{
		app:    app,
		key:    key,
		secret: secret,
	}
}

// 生成 TOKEN 签名
func (t *token) GenerateToken(userId string, expireDuration time.Duration) (tokenString string, err error) {
	// The token content.
	// iss: （Issuer）签发者
	// iat: （Issued At）签发时间，用Unix时间戳表示
	// exp: （Expiration Time）过期时间，用Unix时间戳表示
	// aud: （Audience）接收该JWT的一方
	// sub: （Subject）该JWT的主题
	// nbf: （Not Before）不要早于这个时间
	// jti: （JWT ID）用于标识JWT的唯一ID

	claims := claims{
		userId,
		jwt.StandardClaims{
			Issuer:    t.key,
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			Subject:   t.app,
		},
	}
	tokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.secret))
	return
}

// 解析 TOKEN 签名
func (t *token) ParseToken(tokenString string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
