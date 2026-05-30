// Package auth 提供 JWT 认证工具，支持 HS256 签名算法。
// 功能特性：
//   - JWT Token 生成与解析
//   - 支持自定义 Claims 配置
//   - 完善的错误类型定义
//
// 安全说明：
//   - HS256 算法要求密钥至少 32 字节以确保安全性
//   - 建议使用强随机密钥，不要使用简单字符串
//   - 密钥应通过环境变量或密钥管理服务传递
package auth

import (
	"errors"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

// 定义具体错误类型，便于调用方精确判断错误原因
// 使用 errors.Is() 进行错误类型判断
var (
	// ErrTokenMalformed Token 格式错误
	// 原因：Token 字符串格式不正确，无法解析
	ErrTokenMalformed = errors.New("token is malformed")

	// ErrTokenExpired Token 已过期
	// 原因：Token 的过期时间已超过当前时间
	ErrTokenExpired = errors.New("token is expired")

	// ErrTokenNotActive Token 未激活
	// 原因：Token 的生效时间未到（nbf 字段）
	ErrTokenNotActive = errors.New("token not active yet")

	// ErrTokenInvalid Token 无效
	// 原因：Token 签名验证失败或其他未知错误
	ErrTokenInvalid = errors.New("token is invalid")

	// ErrAppEmpty 应用名称为空
	// 原因：创建 JWT 时未提供应用名称
	ErrAppEmpty = errors.New("app name cannot be empty")

	// ErrKeyEmpty 签发者为空
	// 原因：创建 JWT 时未提供签发者标识
	ErrKeyEmpty = errors.New("key cannot be empty")

	// ErrSecretTooShort 密钥长度不足
	// 原因：HS256 算法要求密钥至少 32 字节
	// 安全建议：使用强随机密钥（如 openssl rand -base64 32）
	ErrSecretTooShort = errors.New("secret must be at least 32 bytes for HS256 algorithm security")
)

// 编译期接口验证，确保 jwt 实现了 JWT 接口
var _ JWT = (*jwt)(nil)

// JWT 定义 JWT 认证接口
// 使用场景：
//   - 用户登录认证
//   - API 接口鉴权
//   - 单点登录（SSO）
type JWT interface {
	// GenerateToken 生成 JWT Token
	// 功能说明：
	//   - 生成包含用户信息的 JWT Token
	//   - 使用 HS256 算法签名
	//   - 支持自定义 audience（接收方）
	//
	// 参数：
	//   - id: 用户唯一标识符（如用户ID、用户名）
	//   - expireDuration: Token 过期时长（如 24小时）
	//   - audience: 可选的接收方列表（不传则默认为用户ID）
	//
	// 返回值：
	//   - tokenString: 生成的 JWT Token 字符串
	//   - err: 生成失败时的错误信息
	//
	// 使用示例：
	//   jwt, _ := auth.NewJWT("my-app", "issuer", "secret-key-32-bytes")
	//   token, err := jwt.GenerateToken("user-123", time.Hour)
	//   // 或指定 audience
	//   token, err := jwt.GenerateToken("user-123", time.Hour, "service-a", "service-b")
	GenerateToken(id string, expireDuration time.Duration, audience ...string) (string, error)

	// ParseToken 解析 JWT Token
	// 功能说明：
	//   - 解析并验证 JWT Token
	//   - 提取 Claims 信息
	//   - 验证签名、过期时间、生效时间等
	//
	// 参数：
	//   - tokenString: JWT Token 字符串
	//
	// 返回值：
	//   - *jwtClaims: 解析后的 Claims 信息
	//   - error: 解析失败时的错误（可使用 errors.Is 判断具体类型）
	//
	// 使用示例：
	//   claims, err := jwt.ParseToken(tokenString)
	//   if errors.Is(err, auth.ErrTokenExpired) {
	//       // Token 已过期，需要重新登录
	//   }
	//   userID := claims.ID
	ParseToken(tokenString string) (*jwtClaims, error)
}

// jwt JWT 认证实现
// 字段说明：
//   - app: 应用名称（用于 Token Subject）
//   - key: 签发者标识（用于 Token Issuer）
//   - secret: 签名密钥（至少 32 字节）
type jwt struct {
	app    string // 应用名称
	key    string // 签发者标识
	secret string // 签名密钥（敏感信息）
}

// jwtClaims JWT Claims 结构体
// 继承标准 RegisteredClaims，包含以下字段：
//   - Issuer (iss): 签发者
//   - Subject (sub): 主题（应用名称）
//   - Audience (aud): 接收方
//   - ExpiresAt (exp): 过期时间
//   - NotBefore (nbf): 生效时间
//   - IssuedAt (iat): 签发时间
//   - ID (jti): JWT ID（用户唯一标识）
type jwtClaims struct {
	gojwt.RegisteredClaims // 内嵌标准 Claims
}

// NewJWT 创建 JWT 实例
// 功能说明：
//   - 创建配置好的 JWT 认证实例
//   - 验证必要参数（应用名称、签发者、密钥）
//   - 强制密钥长度至少 32 字节（HS256 安全要求）
//
// 参数：
//   - app: 应用名称，不能为空（用于 Token Subject）
//   - key: 签发者标识，不能为空（用于 Token Issuer）
//   - secret: 签名密钥，至少 32 字节（建议使用强随机密钥）
//
// 返回值：
//   - JWT: JWT 实例
//   - error: 参数校验失败时的错误
//     可能错误：ErrAppEmpty、ErrKeyEmpty、ErrSecretTooShort
//
// 安全建议：
//   - 密钥生成：openssl rand -base64 32
//   - 密钥存储：环境变量或密钥管理服务
//   - 密钥更新：定期更换密钥，支持密钥轮转
//
// 使用示例：
//
//	jwt, err := auth.NewJWT(
//	    "my-app",           // 应用名称
//	    "auth-service",     // 签发者标识
//	    "strong-secret-key-at-least-32-bytes", // 强密钥
//	)
//	if err != nil {
//	    if errors.Is(err, auth.ErrSecretTooShort) {
//	        // 密钥长度不足，需要重新生成
//	    }
//	    panic(err)
//	}
func NewJWT(app, key, secret string) (JWT, error) {
	if app == "" {
		return nil, ErrAppEmpty
	}
	if key == "" {
		return nil, ErrKeyEmpty
	}
	if len(secret) < 32 {
		return nil, ErrSecretTooShort
	}
	return &jwt{app: app, key: key, secret: secret}, nil
}

// GenerateToken 生成 JWT Token 签名
// 功能说明：
//   - 生成包含用户信息的 JWT Token
//   - 使用 HS256 算法签名
//   - 包含标准 Claims（iss、sub、aud、exp、nbf、iat、jti）
//
// Token 内容说明：
//   - iss (Issuer): 签发者，由 key 参数指定
//   - iat (Issued At): 签发时间，当前时间
//   - exp (Expiration Time): 过期时间，当前时间 + expireDuration
//   - aud (Audience): 接收方，默认为用户 ID，可自定义
//   - sub (Subject): 主题，应用名称
//   - nbf (Not Before): 生效时间，当前时间
//   - jti (JWT ID): JWT 唯一标识，用户 ID
func (t *jwt) GenerateToken(id string, expireDuration time.Duration, audience ...string) (string, error) {
	ts := time.Now()
	claims := jwtClaims{
		gojwt.RegisteredClaims{
			Issuer:    t.key,
			Subject:   t.app,
			NotBefore: gojwt.NewNumericDate(ts),
			IssuedAt:  gojwt.NewNumericDate(ts),
			ExpiresAt: gojwt.NewNumericDate(ts.Add(expireDuration)),
			ID:        id,
		},
	}
	if len(audience) > 0 {
		claims.Audience = audience
	} else {
		claims.Audience = []string{id}
	}
	return gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims).SignedString([]byte(t.secret))
}

// ParseToken 解析 JWT Token 签名
// 功能说明：
//   - 解析并验证 JWT Token
//   - 提取 Claims 信息
//   - 验证签名、过期时间、生效时间等
//
// 返回值错误类型：
//   - ErrTokenMalformed: Token 格式错误
//   - ErrTokenExpired: Token 已过期
//   - ErrTokenNotActive: Token 未激活
//   - ErrTokenInvalid: Token 无效
func (t *jwt) ParseToken(tokenString string) (*jwtClaims, error) {
	tokenClaims, err := gojwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *gojwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})
	if err != nil {
		if errors.Is(err, gojwt.ErrTokenMalformed) {
			return nil, ErrTokenMalformed
		} else if errors.Is(err, gojwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		} else if errors.Is(err, gojwt.ErrTokenNotValidYet) {
			return nil, ErrTokenNotActive
		}
		return nil, ErrTokenInvalid
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, ErrTokenInvalid
}
