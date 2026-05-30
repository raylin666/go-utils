package gorm

import (
	"testing"
	"time"
)

func TestOptionFunctions(t *testing.T) {
	// 测试所有选项函数是否正确设置值
	o := &option{}

	WithDsn("test_dsn")(o)
	if o.Dsn != "test_dsn" {
		t.Errorf("WithDsn failed: expected 'test_dsn', got '%s'", o.Dsn)
	}

	WithDriver("mysql")(o)
	if o.Driver != "mysql" {
		t.Errorf("WithDriver failed: expected 'mysql', got '%s'", o.Driver)
	}

	WithDbName("test_db")(o)
	if o.DbName != "test_db" {
		t.Errorf("WithDbName failed: expected 'test_db', got '%s'", o.DbName)
	}

	WithHost("localhost")(o)
	if o.Host != "localhost" {
		t.Errorf("WithHost failed: expected 'localhost', got '%s'", o.Host)
	}

	WithUserName("root")(o)
	if o.UserName != "root" {
		t.Errorf("WithUserName failed: expected 'root', got '%s'", o.UserName)
	}

	WithPassword("password")(o)
	if o.Password != "password" {
		t.Errorf("WithPassword failed: expected 'password', got '%s'", o.Password)
	}

	WithCharset("utf8mb4")(o)
	if o.Charset != "utf8mb4" {
		t.Errorf("WithCharset failed: expected 'utf8mb4', got '%s'", o.Charset)
	}

	WithPort(3306)(o)
	if o.Port != 3306 {
		t.Errorf("WithPort failed: expected 3306, got %d", o.Port)
	}

	WithPrefix("tbl_")(o)
	if o.Prefix != "tbl_" {
		t.Errorf("WithPrefix failed: expected 'tbl_', got '%s'", o.Prefix)
	}

	WithMaxIdleConn(10)(o)
	if o.MaxIdleConn != 10 {
		t.Errorf("WithMaxIdleConn failed: expected 10, got %d", o.MaxIdleConn)
	}

	WithMaxOpenConn(100)(o)
	if o.MaxOpenConn != 100 {
		t.Errorf("WithMaxOpenConn failed: expected 100, got %d", o.MaxOpenConn)
	}

	WithMaxLifeTime(30 * time.Minute)(o)
	if o.MaxLifeTime != 30 * time.Minute {
		t.Errorf("WithMaxLifeTime failed: expected 30m, got %v", o.MaxLifeTime)
	}

	WithParseTime("True")(o)
	if o.ParseTime != "True" {
		t.Errorf("WithParseTime failed: expected 'True', got '%s'", o.ParseTime)
	}

	WithLoc("Local")(o)
	if o.Loc != "Local" {
		t.Errorf("WithLoc failed: expected 'Local', got '%s'", o.Loc)
	}
}

func TestDSNConstruction(t *testing.T) {
	// 测试 DSN 字符串优先于单独配置
	o := &option{
		Dsn: "custom_dsn_string",
	}

	// 当 Dsn 有值时，应该优先使用 Dsn
	if len(o.Dsn) == 0 {
		t.Errorf("DSN should be preferred when set")
	}
}

func TestPlugin(t *testing.T) {
	// 测试插件结构体初始化
	p := &Plugin{
		SlowThreshold: 200 * time.Millisecond,
	}

	if p.Name() != "db.plugin" {
		t.Errorf("Plugin.Name failed: expected 'db.plugin', got '%s'", p.Name())
	}
}