package gorm

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var _ Client = (*client)(nil)

type Client interface {
	Options() *option
	DB() *gorm.DB
	SqlDB() *sql.DB
	WithPluginBeforeHandler(before func(db *gorm.DB), after func(db *gorm.DB, sql string, ts time.Time)) error
}

type client struct {
	*option
	db    *gorm.DB
	sqlDb *sql.DB
}

func NewClient(opts ...Option) (Client, error) {
	var c = new(client)
	var o = new(option)
	for _, opt := range opts {
		opt(o)
	}
	c.option = o

	var dsn string
	if len(o.Dsn) > 0 {
		dsn = o.Dsn
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
			o.UserName,
			o.Password,
			o.Host,
			o.Port,
			o.DbName,
			o.Charset,
			o.ParseTime,
			o.Loc)
	}

	conn, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   o.Prefix, // 设置表前缀
				SingularTable: true,     // 全局禁用表名复数
			},
		})

	if err != nil {
		return nil, err
	}

	sqlDb, _ := conn.DB()
	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDb.SetMaxIdleConns(o.MaxIdleConn)
	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDb.SetMaxOpenConns(o.MaxOpenConn)
	// 设置最大连接超时
	sqlDb.SetConnMaxLifetime(time.Minute * o.MaxLifeTime)

	c.db = conn
	c.sqlDb = sqlDb
	return c, nil
}

func (c *client) Options() *option {
	return c.option
}

func (c *client) DB() *gorm.DB {
	return c.db
}

func (c *client) SqlDB() *sql.DB {
	return c.sqlDb
}

func (c *client) WithPluginBeforeHandler(before func(db *gorm.DB), after func(db *gorm.DB, sql string, ts time.Time)) error {
	// 插件处理
	return c.db.Use(&Plugin{before, after})
}
