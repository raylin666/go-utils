// Package cmd 提供Redis命令结果包装器
// 功能特性：
//   - 统一的命令结果类型
//   - 原生Redis命令封装
//   - 类型安全的返回值
package cmd

import (
	"github.com/redis/go-redis/v9"
)

// StringSliceCmd 字符串切片命令结果包装器
// 用于返回字符串列表的Redis命令，如 KEYS、MGET、LRANGE 等
type StringSliceCmd struct {
	CMD *redis.StringSliceCmd
}

// NewStringSliceCMD 创建字符串切片命令包装器
// 参数：
//   - cmd: 原生Redis字符串切片命令
// 返回值：
//   - *StringSliceCmd: 包装后的命令对象
func NewStringSliceCMD(cmd *redis.StringSliceCmd) *StringSliceCmd {
	return &StringSliceCmd{cmd}
}

// Result 获取命令执行结果
// 返回值：
//   - []string: 字符串列表
//   - error: 执行错误
//
// 使用示例：
//   keys, err := client.Keys(ctx, "user:*").Result()
//   if err != nil {
//       log.Printf("获取键列表失败: %v", err)
//   }
func (c *StringSliceCmd) Result() ([]string, error) {
	return c.CMD.Result()
}

// String 返回命令的字符串表示
func (c *StringSliceCmd) String() string {
	return c.CMD.String()
}

// FullName 返回命令的完整名称
func (c *StringSliceCmd) FullName() string {
	return c.CMD.FullName()
}

// Args 返回命令参数列表
func (c *StringSliceCmd) Args() []interface{} {
	return c.CMD.Args()
}

// Val 获取命令返回值（不返回错误）
// 注意：如果命令执行失败，返回空切片
func (c *StringSliceCmd) Val() []string {
	return c.CMD.Val()
}

// Err 返回命令执行错误
func (c *StringSliceCmd) Err() error {
	return c.CMD.Err()
}

// ScanSlice 将结果扫描到容器中
// 参数：
//   - container: 目标容器（必须是切片指针）
// 返回值：
//   - error: 扫描错误
//
// 使用示例：
//   var users []User
//   err := client.Keys(ctx, "user:*").ScanSlice(&users)
func (c *StringSliceCmd) ScanSlice(container interface{}) error {
	return c.CMD.ScanSlice(container)
}

// SetVal 设置命令返回值
// 用于测试或模拟场景
func (c *StringSliceCmd) SetVal(val []string) {
	c.CMD.SetVal(val)
}

// SetErr 设置命令错误
// 用于测试或模拟场景
func (c *StringSliceCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}