// Package filesystem 提供文件系统操作功能
// 功能特性：
//   - 目录创建和管理
//   - 支持自定义权限设置
// 使用场景：
//   - 日志目录创建
//   - 数据目录初始化
package filesystem

import (
	"fmt"
	"os"
)

// CreateDirectory 检查目录是否存在，不存在则创建
// 功能说明：
//   - 检查指定路径的目录是否存在
//   - 如果目录不存在，则递归创建所有必要的父目录
//   - 使用更安全的权限设置 0750（目录所有者可读写执行，组用户可读执行，其他用户无权限）
//
// 参数：
//   - dir: 目录路径，支持相对路径和绝对路径
//
// 返回值：
//   - nil: 目录已存在或创建成功
//   - error: 目录检查或创建失败时的错误信息
//     可能的错误类型：
//     - 权限不足（Permission denied）
//     - 路径无效（Invalid path）
//     - 磁盘空间不足（No space left）
//     - 父目录不存在且无法创建
//
// 使用示例：
//   err := filesystem.CreateDirectory("/var/log/myapp")
//   if err != nil {
//       log.Fatalf("创建目录失败: %v", err)
//   }
func CreateDirectory(dir string) error {
	stat, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if createErr := os.MkdirAll(dir, 0750); createErr != nil {
				return fmt.Errorf("failed to create directory [%s]: %w", dir, createErr)
			}
			return nil
		}

		return fmt.Errorf("failed to check directory status [%s]: %w", dir, err)
	}

	if !stat.IsDir() {
		return fmt.Errorf("path [%s] exists but is not a directory", dir)
	}

	return nil
}

// CreateDirectoryWithPerm 使用自定义权限创建目录
// 功能说明：
//   - 与 CreateDirectory 功能相同，但支持自定义权限设置
//   - 适用于需要特殊权限的场景（如共享目录、临时目录等）
//
// 参数：
//   - dir: 目录路径
//   - perm: 目录权限（如 0755、0777 等）
//
// 返回值：
//   - nil: 创建成功
//   - error: 创建失败时的错误信息
//
// 使用示例：
//   // 创建共享目录，允许所有用户读写
//   err := filesystem.CreateDirectoryWithPerm("/tmp/shared", 0777)
func CreateDirectoryWithPerm(dir string, perm os.FileMode) error {
	stat, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if createErr := os.MkdirAll(dir, perm); createErr != nil {
				return fmt.Errorf("failed to create directory [%s]: %w", dir, createErr)
			}
			return nil
		}
		return fmt.Errorf("failed to check directory status [%s]: %w", dir, err)
	}

	if !stat.IsDir() {
		return fmt.Errorf("path [%s] exists but is not a directory", dir)
	}

	return nil
}