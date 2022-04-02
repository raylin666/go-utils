package gorm

import "time"

type Option func(opt *option)

type option struct {
	Dsn			string
	Driver      string
	DbName      string
	Host        string
	UserName    string
	Password    string
	Charset     string
	Port        int
	Prefix      string
	MaxIdleConn int
	MaxOpenConn int
	MaxLifeTime time.Duration
	ParseTime   string
	Loc         string
}

func WithDsn(dsn string) Option {
	return func(opt *option) {
		opt.Dsn = dsn
	}
}

func WithDriver(driver string) Option {
	return func(opt *option) {
		opt.Driver = driver
	}
}

func WithDbName(dbName string) Option {
	return func(opt *option) {
		opt.DbName = dbName
	}
}

func WithHost(host string) Option {
	return func(opt *option) {
		opt.Host = host
	}
}

func WithUserName(userName string) Option {
	return func(opt *option) {
		opt.UserName = userName
	}
}

func WithPassword(password string) Option {
	return func(opt *option) {
		opt.Password = password
	}
}

func WithCharset(charset string) Option {
	return func(opt *option) {
		opt.Charset = charset
	}
}

func WithPort(port int) Option {
	return func(opt *option) {
		opt.Port = port
	}
}

func WithPrefix(prefix string) Option {
	return func(opt *option) {
		opt.Prefix = prefix
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(opt *option) {
		opt.MaxIdleConn = maxIdleConn
	}
}

func WithMaxOpenConn(maxOpenConn int) Option {
	return func(opt *option) {
		opt.MaxOpenConn = maxOpenConn
	}
}

func WithMaxLifeTime(maxLifeTime time.Duration) Option {
	return func(opt *option) {
		opt.MaxLifeTime = maxLifeTime
	}
}

func WithParseTime(parseTime string) Option {
	return func(opt *option) {
		opt.ParseTime = parseTime
	}
}

func WithLoc(loc string) Option {
	return func(opt *option) {
		opt.Loc = loc
	}
}