package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var _ Token = (*token)(nil)

type Token interface {
	// 生成 Token
	GenerateToken(id string, expireDuration time.Duration, opt ClaimsOptions) (tokenString string, err error)
	// 解析 Token
	ParseToken(tokenString string) (*claims, error)
}

type token struct {
	app    string
	key    string
	secret string
}

type claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func New(app string, key string, secret string) Token {
	return &token{
		app:    app,
		key:    key,
		secret: secret,
	}
}

type ClaimsOptions struct {
	ID       string
	Audience jwt.ClaimStrings
}

// 生成 TOKEN 签名
func (t *token) GenerateToken(id string, expireDuration time.Duration, opt ClaimsOptions) (tokenString string, err error) {
	// The token content.
	// iss: （Issuer）签发者
	// iat: （Issued At）签发时间，用Unix时间戳表示
	// exp: （Expiration Time）过期时间，用Unix时间戳表示
	// aud: （Audience）接收该JWT的一方
	// sub: （Subject）该JWT的主题
	// nbf: （Not Before）不要早于这个时间
	// jti: （JWT ID）用于标识JWT的唯一ID

	var ts = time.Now()
	var registeredClaims = jwt.RegisteredClaims{
		Issuer:    t.key,
		Subject:   t.app,
		NotBefore: jwt.NewNumericDate(ts),
		IssuedAt:  jwt.NewNumericDate(ts),
		ExpiresAt: jwt.NewNumericDate(ts.Add(expireDuration)),
	}

	if opt.ID != "" {
		registeredClaims.ID = opt.ID
	}

	if opt.Audience != nil {
		registeredClaims.Audience = opt.Audience
	}

	claims := claims{id, registeredClaims}
	tokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.secret))
	return
}

// 解析 TOKEN 签名
func (t *token) ParseToken(tokenString string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if err != nil {
		if valid, ok := err.(*jwt.ValidationError); ok {
			if valid.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if valid.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if valid.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
