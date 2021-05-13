package database

import (
	"fmt"
	"go-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

const (
	DefaultDatabaseConnection = "default"
)

var (
	db = new(Database)
)

type Database struct {
	Map map[string]*gorm.DB
}

func InitDatabase() {
	var (
		err  error
		conn *gorm.DB
	)

	conf := config.Get().Database
	db.Map = make(map[string]*gorm.DB, len(conf))

	for key, value := range conf {
		var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
			value.UserName,
			value.Password,
			value.Host,
			value.Port,
			value.DbName,
			value.Charset,
			value.ParseTime,
			value.Loc)

		conn, err = gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   value.Prefix, // 设置表前缀
					SingularTable: true,         // 全局禁用表名复数
				},
			})

		if err != nil {
			panic(err)
		}

		sqlDb, _ := conn.DB()
		// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
		sqlDb.SetMaxIdleConns(value.MaxIdleConn)
		// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
		sqlDb.SetMaxOpenConns(value.MaxOpenConn)
		// 设置最大连接超时
		sqlDb.SetConnMaxLifetime(time.Minute * value.MaxLifeTime)

		if value.OpenPlugin {
			// 使用插件
			_ = conn.Use(&TracePlugin{})
		}

		db.Map[strings.ToLower(key)] = conn
	}
}

// 获取链接
func GetDB(connection string) *gorm.DB {
	return db.Map[strings.ToLower(connection)]
}

// 关闭链接
func Close(connection string) error {
	sqlDb, _ := db.Map[strings.ToLower(connection)].DB()
	return sqlDb.Close()
}

// 关闭所有链接
func CloseAll() error {
	for _, connection := range db.Map {
		sqlDb, _ := connection.DB()
		if err := sqlDb.Close(); err != nil {
			return err
		}
	}

	return nil
}
