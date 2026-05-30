package gorm

import "time"

// Option 数据库配置选项函数
// 使用选项模式（Functional Options）进行配置
// 优点：灵活配置、可扩展、默认值友好
type Option func(opt *option)

// option 数据库连接配置结构体
// 字段说明：
//   - Dsn: 数据源名称（Data Source Name），完整连接字符串
//   - Driver: 数据库驱动类型（如 mysql、postgres）
//   - DbName: 数据库名称
//   - Host: 数据库主机地址
//   - UserName: 数据库用户名
//   - Password: 数据库密码（敏感信息，建议通过环境变量传递）
//   - Charset: 字符集编码（如 utf8mb4）
//   - Port: 数据库端口
//   - Prefix: 表名前缀
//   - MaxIdleConn: 最大空闲连接数
//   - MaxOpenConn: 最大打开连接数
//   - MaxLifeTime: 连接最大生命周期（分钟）
//   - ParseTime: 是否解析时间字段
//   - Loc: 时区设置
type option struct {
	Dsn         string        // 数据源名称（完整连接字符串）
	Driver      string        // 数据库驱动类型
	DbName      string        // 数据库名称
	Host        string        // 数据库主机地址
	UserName    string        // 数据库用户名
	Password    string        // 数据库密码（敏感信息）
	Charset     string        // 字符集编码
	Port        int           // 数据库端口
	Prefix      string        // 表名前缀
	MaxIdleConn int           // 最大空闲连接数
	MaxOpenConn int           // 最大打开连接数
	MaxLifeTime time.Duration // 连接最大生命周期
	ParseTime   string        // 是否解析时间
	Loc         string        // 时区设置
}

// WithDsn 设置数据源名称（DSN）
// 功能说明：
//   - 使用完整的连接字符串配置数据库连接
//   - 优先级高于单独配置 Host/Port/UserName 等
//
// 参数：
//   - dsn: 数据源名称，格式如 "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
//
// 使用示例：
//   client, err := gorm.NewClient(
//       gorm.WithDsn("root:password@tcp(localhost:3306)/test?charset=utf8mb4"),
//   )
func WithDsn(dsn string) Option {
	return func(opt *option) {
		opt.Dsn = dsn
	}
}

// WithDriver 设置数据库驱动类型
// 功能说明：
//   - 指定数据库驱动（如 mysql、postgres）
//   - 当前版本仅支持 MySQL
//
// 参数：
//   - driver: 驱动类型字符串
func WithDriver(driver string) Option {
	return func(opt *option) {
		opt.Driver = driver
	}
}

// WithDbName 设置数据库名称
// 参数：
//   - dbName: 数据库名称（必须已存在）
func WithDbName(dbName string) Option {
	return func(opt *option) {
		opt.DbName = dbName
	}
}

// WithHost 设置数据库主机地址
// 参数：
//   - host: 主机地址（IP 或域名），如 "localhost" 或 "192.168.1.100"
func WithHost(host string) Option {
	return func(opt *option) {
		opt.Host = host
	}
}

// WithUserName 设置数据库用户名
// 参数：
//   - userName: 数据库用户名
func WithUserName(userName string) Option {
	return func(opt *option) {
		opt.UserName = userName
	}
}

// WithPassword 设置数据库密码
// 安全建议：
//   - 密码为敏感信息，建议通过环境变量传递
//   - 不要在代码中硬编码密码
//   - 不要在日志中输出密码
//
// 参数：
//   - password: 数据库密码
//
// 使用示例：
//   password := os.Getenv("DB_PASSWORD")
//   client, err := gorm.NewClient(gorm.WithPassword(password))
func WithPassword(password string) Option {
	return func(opt *option) {
		opt.Password = password
	}
}

// WithCharset 设置字符集编码
// 功能说明：
//   - 设置数据库连接字符集
//   - 推荐使用 utf8mb4（支持完整 UTF-8 字符集，包括 emoji）
//
// 参数：
//   - charset: 字符集编码，默认 utf8mb4
func WithCharset(charset string) Option {
	return func(opt *option) {
		opt.Charset = charset
	}
}

// WithPort 设置数据库端口
// 参数：
//   - port: 数据库端口，MySQL 默认 3306
func WithPort(port int) Option {
	return func(opt *option) {
		opt.Port = port
	}
}

// WithPrefix 设置表名前缀
// 功能说明：
//   - 为所有表名添加统一前缀
//   - 用于多项目共用同一数据库时的表名隔离
//
// 参数：
//   - prefix: 表名前缀，如 "app_"
//
// 使用示例：
//   // 设置前缀后，表名 "users" 实际为 "app_users"
//   client, err := gorm.NewClient(gorm.WithPrefix("app_"))
func WithPrefix(prefix string) Option {
	return func(opt *option) {
		opt.Prefix = prefix
	}
}

// WithMaxIdleConn 设置最大空闲连接数
// 功能说明：
//   - 设置连接池中最大空闲连接数
//   - 空闲连接可被复用，减少连接创建开销
//   - 建议值：10-20（根据并发量调整）
//
// 参数：
//   - maxIdleConn: 最大空闲连接数，默认 10
//
// 性能建议：
//   - 低并发场景：5-10
//   - 中并发场景：10-20
//   - 高并发场景：20-50（但需配合 MaxOpenConn）
func WithMaxIdleConn(maxIdleConn int) Option {
	return func(opt *option) {
		opt.MaxIdleConn = maxIdleConn
	}
}

// WithMaxOpenConn 设置最大打开连接数
// 功能说明：
//   - 设置连接池最大连接数（包括活跃和空闲连接）
//   - 防止并发过高导致 "too many connections" 错误
//   - 建议值：100-200（需小于数据库服务器 max_connections 配置）
//
// 参数：
//   - maxOpenConn: 最大打开连接数，默认 100
//
// 性能建议：
//   - 低并发场景：50-100
//   - 中并发场景：100-200
//   - 高并发场景：200-500（需确认数据库服务器配置）
func WithMaxOpenConn(maxOpenConn int) Option {
	return func(opt *option) {
		opt.MaxOpenConn = maxOpenConn
	}
}

// WithMaxLifeTime 设置连接最大生命周期
// 功能说明：
//   - 设置连接在池中的最大存活时间
//   - 防止长时间使用同一连接导致的问题（如数据库服务器重启）
//   - 建议值：30分钟（需小于数据库服务器 wait_timeout 配置）
//
// 参数：
//   - maxLifeTime: 连接最大生命周期（time.Duration），默认 30 分钟
//
// 使用示例：
//   client, err := gorm.NewClient(
//       gorm.WithMaxLifeTime(30 * time.Minute),
//   )
func WithMaxLifeTime(maxLifeTime time.Duration) Option {
	return func(opt *option) {
		opt.MaxLifeTime = maxLifeTime
	}
}

// WithParseTime 设置是否解析时间字段
// 功能说明：
//   - 控制是否将数据库时间字段自动解析为 Go time.Time 类型
//   - 推荐开启（True），便于时间字段操作
//
// 参数：
//   - parseTime: 是否解析时间，默认 "True"
func WithParseTime(parseTime string) Option {
	return func(opt *option) {
		opt.ParseTime = parseTime
	}
}

// WithLoc 设置时区
// 功能说明：
//   - 设置数据库连接时区
//   - 推荐使用 Local（本地时区）或 UTC（统一时区）
//
// 参数：
//   - loc: 时区设置，默认 "Local"
func WithLoc(loc string) Option {
	return func(opt *option) {
		opt.Loc = loc
	}
}