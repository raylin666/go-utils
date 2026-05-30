# Go-Utils

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

通用、模块化的 Go 语言工具组件库，提供多种可复用的功能模块，简化日常开发工作，提高代码复用率。

---

## 特性

- **模块化设计**：每个模块独立封装，按需引入
- **选项模式**：统一的 Functional Options 配置方式
- **接口抽象**：清晰的接口定义，便于扩展和测试
- **健康检查**：统一的 `HealthChecker` 接口，支持服务监控
- **完善文档**：详细的代码注释和使用示例
- **生产可用**：经过实际项目验证，稳定可靠

---

## 安装

```bash
go get github.com/raylin666/go-utils
```

---

## 模块概览

| 模块 | 功能 | 核心依赖 |
|------|------|----------|
| `auth` | JWT 认证（HS256 签名） | `golang-jwt/jwt/v5` |
| `crypto` | 密码哈希与验证（bcrypt） | `golang.org/x/crypto` |
| `cache/redis` | Redis 客户端封装 | `redis/go-redis/v9` |
| `config` | YAML 配置加载 | `gopkg.in/yaml.v3` |
| `db/gorm` | GORM 数据库客户端 | `gorm.io/gorm` |
| `dingtalk` | 钉钉机器人消息推送 | 内置 HTTP 客户端 |
| `errors` | 错误包装与堆栈追踪 | `pkg/errors` |
| `filesystem` | 目录创建与权限管理 | 无外部依赖 |
| `http` | HTTP 客户端/服务端 | `gojek/heimdall` |
| `logger` | JSON 日志（Zap）+ 日志轮转 | `uber-go/zap` |
| `mail` | SMTP 邮件发送 | `gopkg.in/gomail.v2` |
| `middleware` | HTTP/gRPC 中间件链 | 无外部依赖 |
| `netx` | 网络地址提取 | 无外部依赖 |
| `server` | 服务生命周期管理 | 无外部依赖 |
| `timeutil` | 重试机制 + 超时配置 | 无外部依赖 |
| `upload/qiniu` | 七牛云存储操作 | `qiniu/go-sdk/v7` |
| `validator` | 结构体数据验证 | `playground/validator/v10` |

---

## 快速开始

### 导入模块

```go
import (
    "github.com/raylin666/go-utils/auth"
    "github.com/raylin666/go-utils/cache/redis"
    "github.com/raylin666/go-utils/config"
    "github.com/raylin666/go-utils/crypto"
    "github.com/raylin666/go-utils/db/gorm"
    "github.com/raylin666/go-utils/logger"
    "github.com/raylin666/go-utils/mail"
    "github.com/raylin666/go-utils/timeutil"
    "github.com/raylin666/go-utils/validator"
)
```

---

## 核心模块使用指南

### 🔐 JWT 认证（auth）

提供 JWT Token 生成与解析功能，使用 HS256 算法签名。

**安全要求**：密钥至少 32 字节，建议使用 `openssl rand -base64 32` 生成。

```go
package main

import (
    "errors"
    "time"
    "github.com/raylin666/go-utils/auth"
)

func main() {
    jwt, err := auth.NewJWT(
        "my-app",                           // 应用名称
        "auth-service",                     // 签发者标识
        "strong-secret-key-at-least-32-bytes", // 签名密钥
    )
    if err != nil {
        if errors.Is(err, auth.ErrSecretTooShort) {
            panic("密钥必须至少 32 字节")
        }
        panic(err)
    }

    token, err := jwt.GenerateToken("user-123", 24*time.Hour)
    if err != nil {
        panic(err)
    }
    println("Token:", token)

    claims, err := jwt.ParseToken(token)
    if err != nil {
        if errors.Is(err, auth.ErrTokenExpired) {
            panic("Token 已过期")
        }
        if errors.Is(err, auth.ErrTokenMalformed) {
            panic("Token 格式错误")
        }
        panic(err)
    }
    println("用户 ID:", claims.ID)
}
```

**错误类型**：
- `ErrTokenMalformed` - Token 格式错误
- `ErrTokenExpired` - Token 已过期
- `ErrTokenNotActive` - Token 未激活
- `ErrTokenInvalid` - Token 无效
- `ErrSecretTooShort` - 密钥长度不足

---

### 🔒 密码哈希（crypto）

基于 bcrypt 算法的密码哈希与验证。

```go
package main

import (
    "github.com/raylin666/go-utils/crypto"
)

func main() {
    hash, err := crypto.BcryptPasswordHash("mypassword")
    if err != nil {
        panic(err)
    }
    println("哈希值:", hash)

    if crypto.BcryptPasswordVerify("mypassword", hash) {
        println("密码验证成功")
    }

    min, max, defaultCost := crypto.BcryptCostRange()
    println("Cost范围:", min, "-", max, "默认:", defaultCost)
}
```

---

### 📦 Redis 客户端（cache/redis）

完整的 Redis 命令支持，包含健康检查和连接管理。

```go
package main

import (
    "context"
    "time"
    "github.com/raylin666/go-utils/cache/redis"
)

func main() {
    ctx := context.Background()
    
    client, err := redis.NewClient(ctx, &redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
        PoolSize: 10,
    })
    if err != nil {
        panic(err)
    }
    defer client.Close()

    if err := client.HealthCheck(ctx); err != nil {
        panic("Redis 连接异常")
    }

    client.Set(ctx, "key", "value", 10*time.Minute)
    
    val, err := client.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    println("值:", val)

    keys, err := client.SafeScanKeys(ctx, "user:*", 100)
    if err != nil {
        panic(err)
    }
    println("匹配键数量:", len(keys))
}
```

**核心方法**：
- `Set/Get/Del` - 基础键操作
- `HSet/HGet/HGetAll` - 哈希操作
- `LPush/RPush/LPop/RPop` - 列表操作
- `SAdd/SRem/SMembers` - 集合操作
- `ZAdd/ZRem/ZRange` - 有序集合操作
- `Pipeline/TxPipeline` - 管道/事务
- `SafeScanKeys` - 安全键扫描（生产环境推荐）

---

### 🗄️ 数据库连接（db/gorm）

GORM 数据库客户端，支持连接池配置和慢查询监控。

```go
package main

import (
    "context"
    "time"
    "github.com/raylin666/go-utils/db/gorm"
)

func main() {
    client, err := gorm.NewClient(
        gorm.WithHost("localhost"),
        gorm.WithPort(3306),
        gorm.WithUserName("root"),
        gorm.WithPassword("password"),
        gorm.WithDbName("test"),
        gorm.WithCharset("utf8mb4"),
        gorm.WithMaxIdleConn(10),
        gorm.WithMaxOpenConn(100),
        gorm.WithMaxLifeTime(30 * time.Minute),
    )
    if err != nil {
        panic(err)
    }
    defer client.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := client.HealthCheck(ctx); err != nil {
        panic("数据库连接异常")
    }

    err = client.WithPluginBeforeHandler(nil, func(db *gorm.DB, sql string, ts time.Time) {
        duration := time.Since(ts)
        if duration > 200*time.Millisecond {
            println("慢查询:", sql, "耗时:", duration)
        }
    })

    db := client.DB()
    db.Find(&users)
}
```

**配置选项**：
- `WithDsn` - 使用完整 DSN 配置
- `WithHost/WithPort/WithUserName/WithPassword/WithDbName` - 分离参数配置
- `WithMaxIdleConn` - 最大空闲连接（默认 10）
- `WithMaxOpenConn` - 最大打开连接（默认 100）
- `WithMaxLifeTime` - 连接生命周期（默认 30 分钟）
- `WithPrefix` - 表名前缀

---

### 📝 YAML 配置（config）

YAML 配置文件加载，支持文件路径和字符串两种方式。

```go
package main

import (
    "github.com/raylin666/go-utils/config"
)

type AppConfig struct {
    Server struct {
        Port int `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        Host string `yaml:"host"`
        Port int    `yaml:"port"`
    } `yaml:"database"`
}

func main() {
    var cfg AppConfig
    err := config.LoadYaml("config.yaml", &cfg)
    if err != nil {
        panic(err)
    }
    println("端口:", cfg.Server.Port)

    yamlStr := "server:\n  port: 8080"
    var cfg2 AppConfig
    err = config.LoadYamlByString(yamlStr, &cfg2)
}
```

---

### 📊 日志记录（logger）

基于 Zap 的高性能 JSON 日志，支持日志轮转和压缩。

```go
package main

import (
    "github.com/raylin666/go-utils/logger"
    "go.uber.org/zap"
)

func main() {
    log, err := logger.NewJSONLogger(
        logger.WithInfoLevel(),
        logger.WithField("app", "my-service"),
        logger.WithField("environment", "production"),
        logger.WithPathFileRotation("/var/log/app.log", logger.PathFileRotationOption{
            MaxSize:    100,  // MB
            MaxBackups: 10,   // 保留文件数
            MaxAge:     30,   // 保留天数
            LocalTime:  true,
            Compress:   true,
        }),
    )
    if err != nil {
        panic(err)
    }
    defer log.Close()

    log.Info("服务启动", 
        zap.String("version", "1.0.0"),
        zap.Int("port", 8080))
    
    log.Error("操作失败", 
        zap.String("operation", "query"),
        zap.Error(err))

    logger.SetDebugLevel(log)
    log.Debug("调试信息")
}
```

**日志级别**：
- `WithDebugLevel` - 调试级别
- `WithInfoLevel` - 信息级别（默认）
- `WithWarnLevel` - 警告级别
- `WithErrorLevel` - 错误级别

---

### 📧 邮件发送（mail）

SMTP 邮件发送，支持超时控制和 TLS 加密。

```go
package main

import (
    "context"
    "time"
    "github.com/raylin666/go-utils/mail"
)

func main() {
    m, err := mail.New(
        mail.WithMailHost("smtp.qq.com"),
        mail.WithMailPort(465),
        mail.WithMailUser("your@qq.com"),
        mail.WithMailPass("your-password"),
        mail.WithMailTimeout(30*time.Second),
    )
    if err != nil {
        panic(err)
    }
    defer m.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    err = m.SendTextHtmlWithContext(ctx, 
        "测试邮件", 
        "<h1>Hello World</h1>", 
        []string{"recipient@example.com"})
    if err != nil {
        panic(err)
    }
}
```

---

### ✅ 数据验证（validator）

结构体数据验证，支持中英文错误消息。

```go
package main

import (
    "fmt"
    "github.com/raylin666/go-utils/validator"
)

type User struct {
    Name  string `label:"姓名" validate:"required"`
    Email string `label:"邮箱" validate:"required,email"`
    Age   int    `label:"年龄" validate:"gte=0,lte=150"`
    Phone string `label:"电话" validate:"omitempty,len=11"`
}

func main() {
    v := validator.New(validator.WithLocale("zh"))

    user := User{Name: "", Email: "invalid", Age: 200}
    
    err := v.Validate(user)
    if err != nil {
        fmt.Println("验证失败:", err.Error())
    }

    err = v.ValidateAll(user)
    if err != nil {
        fmt.Println("所有错误:", err.Error())
    }
}
```

**验证规则**：
- `required` - 必填字段
- `email` - 邮箱格式
- `gte/lte` - 数值范围
- `min/max` - 字符串长度
- `omitempty` - 可选字段

---

### 🔄 重试机制（timeutil）

指数退避重试策略，支持超时控制和自定义条件。

```go
package main

import (
    "context"
    "time"
    "github.com/raylin666/go-utils/timeutil"
)

func main() {
    ctx := context.Background()

    err := timeutil.RetryWithBackoff(ctx, timeutil.DefaultRetryConfig, func() error {
        return connectDB()
    })
    if err != nil {
        if timeutil.IsRetryError(err) {
            attempts := timeutil.GetRetryAttempts(err)
            println("重试失败，尝试次数:", attempts)
        }
        panic(err)
    }

    config := timeutil.RetryConfig{
        MaxAttempts:     5,
        InitialDelay:    200 * time.Millisecond,
        MaxDelay:        10 * time.Second,
        Multiplier:      2.0,
        RetryCondition: func(err error) bool {
            return isNetworkError(err)
        },
    }
    err = timeutil.RetryWithBackoff(ctx, config, func() error {
        return callAPI()
    })

    err = timeutil.RetryWithJitter(ctx, config, func() error {
        return distributedOperation()
    })
}
```

**重试策略**：
- `RetryWithBackoff` - 指数退避（推荐）
- `RetryWithFixedDelay` - 固定延迟
- `RetryWithJitter` - 抖动策略（分布式场景）

---

### ⏱️ 超时配置（timeutil）

统一的超时常量和 Context 创建工具。

```go
package main

import (
    "context"
    "github.com/raylin666/go-utils/timeutil"
)

func main() {
    ctx, cancel := timeutil.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    ctx2, cancel2 := timeutil.WithConnectTimeout(context.Background())
    defer cancel2()

    ctx3, cancel3 := timeutil.WithQueryTimeout(context.Background())
    defer cancel3()

    println("连接超时:", timeutil.DefaultConnectTimeout)
    println("查询超时:", timeutil.DefaultQueryTimeout)
}
```

**超时常量**：
- `DefaultConnectTimeout` - 5s（连接）
- `DefaultReadTimeout` - 10s（读取）
- `DefaultQueryTimeout` - 15s（查询）
- `DefaultOperationTimeout` - 30s（通用操作）

---

### 🌐 HTTP 客户端（http）

支持熔断、重试的 HTTP 客户端。

```go
package main

import (
    "net/http"
    "time"
    "github.com/raylin666/go-utils/http"
)

func main() {
    client := http.NewClient(
        http.WithClientHTTPTimeout(30*time.Second),
        http.WithClientRetryCount(3),
    )

    resp, err := client.GET("https://api.example.com/data", http.Header{
        "Authorization": []string{"Bearer token"},
    })
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    resp, err = client.POST("https://api.example.com/create", 
        bytes.NewBuffer(jsonData), 
        http.Header{"Content-Type": []string{"application/json"})
}
```

---

### 🖥️ HTTP 服务端（http）

HTTP 服务器管理，支持 TLS 和优雅关闭。

```go
package main

import (
    "context"
    "net/http"
    "github.com/raylin666/go-utils/http"
)

func main() {
    hs := &http.Server{Addr: ":8080"}
    
    srv := http.NewServer(hs,
        http.WithServerAddress(":8080"),
        http.WithServerNetwork("tcp"),
    )

    endpoint, err := srv.Endpoint()
    if err != nil {
        panic(err)
    }
    println("服务端点:", endpoint.String())

    ctx := context.Background()
    if err := srv.Start(ctx); err != nil {
        panic(err)
    }
}
```

---

### 🔔 钉钉机器人（dingtalk）

钉钉机器人消息推送，支持多种消息类型。

```go
package main

import (
    "github.com/raylin666/go-utils/dingtalk"
)

func main() {
    robot := dingtalk.NewRobot("your-access-token")

    resp, err := robot.SendTextMessage(dingtalk.RobotTextMessageType{
        Content: "服务告警：数据库连接异常",
    })

    resp, err = robot.SendMarkdownMessage(dingtalk.RobotMarkdownMessageType{
        Title: "服务状态报告",
        Text:  "## 服务状态\n- **状态**: 正常\n- **时间**: 2024-01-01",
    })

    resp, err = robot.SendLinkMessage(dingtalk.RobotLinkMessageType{
        Title:      "监控详情",
        Text:       "点击查看详细监控数据",
        MessageURL: "https://monitor.example.com",
        PicURL:     "https://example.com/pic.png",
    })
}
```

---

### ☁️ 七牛云存储（upload/qiniu）

七牛云存储操作，支持上传、下载、管理。

```go
package main

import (
    "github.com/raylin666/go-utils/upload/qiniu"
    "github.com/qiniu/go-sdk/v7/storage"
)

func main() {
    q := qiniu.New(
        "your-access-key",
        "your-secret-key",
        "your-bucket",
        "huanan",
        nil,
    )

    ret, err := q.FormUploaderPutFile("/path/to/file.txt", "file.txt")
    if err != nil {
        panic(err)
    }

    publicURL := q.MakePublicURL("https://cdn.example.com", "file.txt")
    
    privateURL := q.MakePrivateURL("https://cdn.example.com", "file.txt", 3600)

    info, err := q.GetFileInfo("file.txt")
    if err != nil {
        panic(err)
    }
    println("文件大小:", info.Fsize)

    err = q.Delete("file.txt")
}
```

---

### 🌍 网络工具（netx）

网络地址提取和 IP 获取。

```go
package main

import (
    "github.com/raylin666/go-utils/netx"
)

func main() {
    addr, err := netx.ExtractAddress("0.0.0.0:8080", nil)
    if err != nil {
        panic(err)
    }
    println("地址:", addr)

    host, port, err := netx.ExtractHostPort("localhost:8080")
    if err != nil {
        panic(err)
    }
    println("主机:", host, "端口:", port)

    ip := netx.GetLocalServerIp()
    println("本机 IP:", ip)
}
```

---

### 📁 文件系统（filesystem）

目录创建与权限管理。

```go
package main

import (
    "github.com/raylin666/go-utils/filesystem"
)

func main() {
    err := filesystem.CreateDirectory("/var/log/myapp")
    if err != nil {
        panic(err)
    }

    err = filesystem.CreateDirectoryWithPerm("/tmp/shared", 0777)
    if err != nil {
        panic(err)
    }
}
```

---

### 🔗 中间件链（middleware）

HTTP/gRPC 中间件链式调用。

```go
package main

import (
    "context"
    "net/http"
    "github.com/raylin666/go-utils/middleware"
)

func main() {
    handler := middleware.HTTPChain(
        loggingMiddleware,
        authMiddleware,
        corsMiddleware,
    )(finalHandler)

    http.Handle("/", handler)
    http.ListenAndServe(":8080", nil)
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        println("请求:", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

---

### 🛡️ 优雅关闭（server/system）

服务优雅关闭，支持信号处理。

```go
package main

import (
    "github.com/raylin666/go-utils/server/system"
)

func main() {
    hook := system.NewShutdown()
    
    hook.WithSignals(syscall.SIGINT, syscall.SIGTERM)
    
    hook.Close(
        func() { println("关闭数据库连接") },
        func() { println("关闭 Redis 连接") },
        func() { println("关闭日志文件") },
    )
}
```

---

### ⚠️ 错误处理（errors）

错误包装与堆栈追踪。

```go
package main

import (
    "github.com/raylin666/go-utils/errors"
)

func main() {
    err := errors.New("基础错误")
    
    err = errors.Wrap(err, "附加信息")
    
    err = errors.Errorf("格式化错误: %s", "detail")
    
    err = errors.WithStack(err)
    
    println(err.Error())
}
```

---

## 设计原则

### 选项模式（Functional Options）

所有模块使用一致的选项模式进行配置：

```go
client, err := NewClient(
    WithOption1("value1"),
    WithOption2("value2"),
    WithOption3("value3"),
)
```

### 接口抽象

每个模块定义清晰的接口，便于扩展和模拟测试：

```go
type Client interface {
    DoSomething() error
}

var _ Client = (*client)(nil)
```

### 健康检查接口

所有连接类模块实现统一的健康检查接口：

```go
type HealthChecker interface {
    HealthCheck(ctx context.Context) error
    IsConnected() bool
}
```

---

## 项目结构

```
go-utils/
├── auth/               # JWT 认证
│   ├── jwt.go          # JWT 实现
│   └── jwt_test.go     # 测试
├── cache/
│   └── redis/          # Redis 客户端
│       ├── client.go   # 客户端实现
│       └── cmd/        # 命令封装
├── config/             # YAML 配置
│   └── yaml.go
├── crypto/             # 密码哈希
│   └── bcrypt.go
├── db/
│   └── gorm/           # GORM 客户端
│       ├── client.go
│       ├── option.go   # 配置选项
│       └── plugin.go   # SQL 监控插件
├── dingtalk/           # 钉钉机器人
│   ├── robot.go
│   └── robot_messagetype.go
├── errors/             # 错误处理
│   └── error.go
├── filesystem/         # 文件系统
│   └── directory.go
├── http/               # HTTP 组件
│   ├── client.go       # HTTP 客户端
│   └── server.go       # HTTP 服务端
├── logger/             # 日志模块
│   ├── logger.go
│   ├── option.go       # 配置选项
│   └── meta.go
├── mail/               # 邮件模块
│   └── gomail.go
├── middleware/         # 中间件链
│   └── middleware.go
├── netx/               # 网络工具
│   └── host.go
├── server/             # 服务管理
│   ├── endpoint.go
│   └── system/
│       └── shutdown.go
├── timeutil/           # 时间工具
│   ├── retry.go        # 重试机制
│   └── timeout.go      # 超时配置
├── upload/
│   └── qiniu/          # 七牛云
│       └── qiniu.go
├── validator/          # 数据验证
│   └ validator.go
├── health.go           # 健康检查接口
├── go.mod
├── go.sum
└── README.md
```

---

## 安全建议

1. **JWT 密钥**：使用至少 32 字节的强密钥，通过 `openssl rand -base64 32` 生成
2. **敏感配置**：密码等敏感信息通过环境变量传递，不要硬编码
3. **文件权限**：日志文件权限 `0640`，目录权限 `0750`
4. **超时控制**：所有网络操作使用 context 进行超时控制
5. **错误处理**：使用 `errors.Is()` 判断具体错误类型

---

## 测试

```bash
go test ./...

go test ./auth/...
go test ./db/gorm/...

go test -cover ./...

go test -bench=. ./timeutil/...
```

---

## 依赖版本

| 依赖 | 版本 |
|------|------|
| gorm.io/gorm | v1.25.12 |
| redis/go-redis/v9 | v9.7.0 |
| uber-go/zap | v1.27.0 |
| golang-jwt/jwt/v5 | v5.3.1 |
| playground/validator/v10 | v10.26.0 |
| qiniu/go-sdk/v7 | v7.25.6 |
| gojek/heimdall/v7 | v7.0.2 |
| gopkg.in/gomail.v2 | v2.0.0 |
| golang.org/x/crypto | v0.39.0 |
| pkg/errors | v0.9.1 |
| gopkg.in/yaml.v3 | v3.0.1 |

---

## 许可证

MIT License

---

## 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request