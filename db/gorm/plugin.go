// Package gorm 提供 GORM 数据库连接管理功能，支持插件扩展。
// 功能特性：
//   - MySQL 数据库连接管理
//   - 连接池配置（空闲连接、最大连接、生命周期）
//   - 慢查询监控插件
//   - SQL 执行追踪能力
//   - 表名前缀配置
//   - 健康检查支持
package gorm

import (
	"time"

	"gorm.io/gorm"
)

const (
	callBackBeforeName = "db:callback:before"
	callBackAfterName  = "db:callback:after"
	startTime          = "_start_time"
)

// Initialize 初始化GORM插件，注册回调函数
// 在所有SQL操作前后注册钩子函数，用于监控和追踪
func (op *Plugin) Initialize(db *gorm.DB) (err error) {
	if err = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, op.before); err != nil {
		return err
	}
	if err = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, op.before); err != nil {
		return err
	}
	if err = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, op.before); err != nil {
		return err
	}
	if err = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, op.before); err != nil {
		return err
	}
	if err = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, op.before); err != nil {
		return err
	}
	if err = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, op.before); err != nil {
		return err
	}

	if err = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, op.after); err != nil {
		return err
	}
	if err = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, op.after); err != nil {
		return err
	}
	if err = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, op.after); err != nil {
		return err
	}
	if err = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, op.after); err != nil {
		return err
	}
	if err = db.Callback().Row().After("gorm:row").Register(callBackAfterName, op.after); err != nil {
		return err
	}
	if err = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, op.after); err != nil {
		return err
	}

	return nil
}

var _ gorm.Plugin = (*Plugin)(nil)

// Plugin GORM插件，支持SQL执行监控和慢查询检测
// 使用场景：
//   - SQL性能监控
//   - 慢查询告警
//   - SQL执行追踪
//   - 性能优化分析
type Plugin struct {
	// Before SQL执行前回调
	// 参数：
	//   - db: GORM数据库上下文
	//
	// 使用示例：
	//   Before: func(db *gorm.DB) {
	//       log.Printf("开始执行SQL: %s", db.Statement.SQL.String())
	//   }
	Before func(db *gorm.DB)

	// After SQL执行后回调
	// 参数：
	//   - db: GORM数据库上下文
	//   - sql: 完整SQL语句（包含参数值）
	//   - ts: SQL开始执行时间
	//
	// 使用示例：
	//   After: func(db *gorm.DB, sql string, ts time.Time) {
	//       duration := time.Since(ts)
	//       log.Printf("SQL执行完成: %s, 耗时: %v", sql, duration)
	//   }
	After func(db *gorm.DB, sql string, ts time.Time)

	// SlowThreshold 慢查询阈值，超过此时间会触发慢查询告警
	// 建议值：200ms-500ms（根据业务需求调整）
	SlowThreshold time.Duration

	// OnSlowQuery 慢查询回调函数
	// 参数：
	//   - db: GORM数据库上下文
	//   - sql: 慢查询SQL语句
	//   - duration: SQL执行时长
	//
	// 使用示例：
	//   OnSlowQuery: func(db *gorm.DB, sql string, duration time.Duration) {
	//       log.Printf("慢查询告警: %s, 耗时: %v", sql, duration)
	//       metrics.RecordSlowQuery(sql, duration)
	//   }
	OnSlowQuery func(db *gorm.DB, sql string, duration time.Duration)
}

// Name 返回插件名称
func (op *Plugin) Name() string {
	return "db.plugin"
}

// before SQL执行前钩子
// 记录开始时间，用于后续计算执行时长
func (op *Plugin) before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
	if op.Before != nil {
		op.Before(db)
	}
}

// after SQL执行后钩子
// 计算执行时长，检查慢查询，执行回调
func (op *Plugin) after(db *gorm.DB) {
	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}
	// 计算SQL执行时长
	duration := time.Since(ts)
	// 构建SQL语句
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	// 检查慢查询

	if op.SlowThreshold > 0 && duration > op.SlowThreshold {
		if op.OnSlowQuery != nil {
			op.OnSlowQuery(db, sql, duration)
		}
	}

	// 执行After回调
	if op.After != nil {
		op.After(db, sql, ts)
	}
}
