// Package crypto 提供密码哈希和验证功能
// 功能特性：
//   - 基于 bcrypt 算法的密码哈希
//   - 支持自定义 cost 参数
//   - 密码验证功能
// 使用场景：
//   - 用户密码存储
//   - 密码验证
package crypto

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrBcryptCostTooLow  = fmt.Errorf("bcrypt cost must be at least %d", bcrypt.MinCost)
	ErrBcryptCostTooHigh = fmt.Errorf("bcrypt cost must be at most %d", bcrypt.MaxCost)
	ErrPasswordEmpty     = fmt.Errorf("password cannot be empty")
)

// BcryptPasswordHash 使用默认 cost 对密码进行哈希
// 参数：
//   - password: 待哈希的密码
// 返回值：
//   - string: 哈希后的密码
//   - error: 哈希失败时的错误信息
func BcryptPasswordHash(password string) (string, error) {
	return BcryptPasswordHashWithCost(password, bcrypt.DefaultCost)
}

// BcryptPasswordHashWithCost 使用自定义 cost 对密码进行哈希
// 功能说明：
//   - 使用 bcrypt 算法对密码进行哈希
//   - cost 参数控制哈希计算复杂度（越大越安全但越慢）
//
// 参数：
//   - password: 待哈希的密码
//   - cost: 哈希复杂度（范围：bcrypt.MinCost 到 bcrypt.MaxCost）
//
// 返回值：
//   - string: 哈希后的密码（可直接存储）
//   - error: 哈希失败时的错误信息
//
// 使用示例：
//   hash, err := crypto.BcryptPasswordHashWithCost("mypassword", 12)
//   if err != nil {
//       log.Fatal(err)
//   }
func BcryptPasswordHashWithCost(password string, cost int) (string, error) {
	if password == "" {
		return "", ErrPasswordEmpty
	}

	if cost < bcrypt.MinCost {
		return "", ErrBcryptCostTooLow
	}

	if cost > bcrypt.MaxCost {
		return "", ErrBcryptCostTooHigh
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("password hashing failed: %w", err)
	}

	return string(bytes), nil
}

// BcryptPasswordVerify 验证密码是否匹配哈希值
// 功能说明：
//   - 验证原始密码与存储的哈希值是否匹配
//   - 用于用户登录验证
//
// 参数：
//   - password: 待验证的原始密码
//   - hash: 存储的密码哈希值
//
// 返回值：
//   - bool: true 表示密码匹配，false 表示不匹配
//
// 使用示例：
//   if crypto.BcryptPasswordVerify("mypassword", storedHash) {
//       // 密码验证成功
//   }
func BcryptPasswordVerify(password, hash string) bool {
	if password == "" || hash == "" {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// BcryptCostRange 返回 bcrypt cost 参数的有效范围
// 返回值：
//   - min: 最小 cost 值
//   - max: 最大 cost 值
//   - defaultCost: 默认 cost 值
func BcryptCostRange() (min, max, defaultCost int) {
	return bcrypt.MinCost, bcrypt.MaxCost, bcrypt.DefaultCost
}