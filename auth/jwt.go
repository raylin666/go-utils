package auth

import (
	"errors"
	gojwt "github.com/golang-jwt/jwt/v4"
	"time"
)

var _ JWT = (*jwt)(nil)

type JWT interface {
	// 生成 Token
	GenerateToken(id string, expireDuration time.Duration, opt JWTClaimsOptions) (tokenString string, err error)
	// 解析 Token
	ParseToken(tokenString string) (*jwtClaims, error)
}

type jwt struct {
	app    string
	key    string
	secret string
}

type jwtClaims struct {
	ID string `json:"id"`
	gojwt.RegisteredClaims
}

func NewJWT(app string, key string, secret string) JWT {
	return &jwt{
		app:    app,
		key:    key,
		secret: secret,
	}
}

type JWTClaimsOptions struct {
	ID       string
	Audience gojwt.ClaimStrings
}

// 生成 TOKEN 签名
func (t *jwt) GenerateToken(id string, expireDuration time.Duration, opt JWTClaimsOptions) (tokenString string, err error) {
	// The token content.
	// iss: （Issuer）签发者
	// iat: （Issued At）签发时间，用Unix时间戳表示
	// exp: （Expiration Time）过期时间，用Unix时间戳表示
	// aud: （Audience）接收该JWT的一方
	// sub: （Subject）该JWT的主题
	// nbf: （Not Before）不要早于这个时间
	// jti: （JWT ID）用于标识JWT的唯一ID

	var ts = time.Now()
	var registeredClaims = gojwt.RegisteredClaims{
		Issuer:    t.key,
		Subject:   t.app,
		NotBefore: gojwt.NewNumericDate(ts),
		IssuedAt:  gojwt.NewNumericDate(ts),
		ExpiresAt: gojwt.NewNumericDate(ts.Add(expireDuration)),
	}

	if opt.ID != "" {
		registeredClaims.ID = opt.ID
	}

	if opt.Audience != nil {
		registeredClaims.Audience = opt.Audience
	}

	claims := jwtClaims{id, registeredClaims}
	tokenString, err = gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims).SignedString([]byte(t.secret))
	return
}

// 解析 TOKEN 签名
func (t *jwt) ParseToken(tokenString string) (*jwtClaims, error) {
	tokenClaims, err := gojwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *gojwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if err != nil {
		if valid, ok := err.(*gojwt.ValidationError); ok {
			if valid.Errors&gojwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if valid.Errors&gojwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if valid.Errors&gojwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
