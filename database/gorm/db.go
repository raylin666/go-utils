package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Options struct {
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
	OpenPlugin  bool
}

type PluginConfig struct {
	Before func(gormDb *gorm.DB)
	After  func(gormDb *gorm.DB, sql string, ts time.Time)
}

type DB struct {
	*gorm.DB
}

func New(opts Options, pgc PluginConfig) (*DB, error) {
	var (
		dsn  string
		err  error
		conn *gorm.DB
	)

	if len(opts.Dsn) > 0 {
		dsn = opts.Dsn
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
			opts.UserName,
			opts.Password,
			opts.Host,
			opts.Port,
			opts.DbName,
			opts.Charset,
			opts.ParseTime,
			opts.Loc)
	}

	conn, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   opts.Prefix, // 设置表前缀
				SingularTable: true,         // 全局禁用表名复数
			},
		})

	if err != nil {
		return nil, err
	}

	sqlDb, _ := conn.DB()
	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDb.SetMaxIdleConns(opts.MaxIdleConn)
	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDb.SetMaxOpenConns(opts.MaxOpenConn)
	// 设置最大连接超时
	sqlDb.SetConnMaxLifetime(time.Minute * opts.MaxLifeTime)

	if opts.OpenPlugin {
		// 使用插件
		_ = conn.Use(&TracePlugin{
			Before: pgc.Before,
			After:  pgc.After,
		})
	}

	return &DB{conn}, nil
}
