// Package middleware 提供通用中间件管理系统。
package middleware

import (
	"sort"
)

// Manager 中间件管理器。
// 负责收集、排序和构建中间件链。
// 这是一个通用实现，不依赖任何特定框架。
type Manager struct {
	middlewares []Middleware
}

// NewManager 创建新的中间件管理器。
//
// 返回:
//   - *Manager: 中间件管理器实例
func NewManager() *Manager {
	return &Manager{
		middlewares: make([]Middleware, 0),
	}
}

// Use 添加中间件到管理器。
// 支持链式调用，可以连续添加多个中间件。
//
// 参数:
//   - middleware: 要添加的中间件实例
//
// 返回:
//   - *Manager: 中间件管理器实例（支持链式调用）
func (m *Manager) Use(middleware Middleware) *Manager {
	if middleware == nil {
		return m
	}

	// 检查是否启用（如果实现了 Config 接口）
	if config, ok := middleware.(Config); ok && !config.Enabled() {
		return m
	}

	m.middlewares = append(m.middlewares, middleware)
	return m
}

// UseFunc 使用函数方式添加中间件。
// 提供简化的中间件添加方式，无需创建完整结构体。
//
// 参数:
//   - name: 中间件名称
//   - priority: 中间件优先级
//   - handler: 中间件处理函数
//
// 返回:
//   - *Manager: 中间件管理器实例（支持链式调用）
func (m *Manager) UseFunc(name string, priority Priority, handler Handler) *Manager {
	return m.Use(NewMiddlewareFunc(name, priority, handler))
}

// Build 构建中间件链。
// 按优先级排序中间件，并返回处理函数列表。
// 优先级数值越小，执行顺序越靠前。
//
// 返回:
//   - []Handler: 排序后的中间件处理函数列表
func (m *Manager) Build() []Handler {
	if len(m.middlewares) == 0 {
		return nil
	}

	// 复制中间件列表，避免修改原列表
	sorted := make([]Middleware, len(m.middlewares))
	copy(sorted, m.middlewares)

	// 按优先级排序（数值越小优先级越高）
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Priority() < sorted[j].Priority()
	})

	// 转换为处理函数列表
	handlers := make([]Handler, len(sorted))
	for i, mw := range sorted {
		handlers[i] = mw.Handler()
	}

	return handlers
}

// List 返回所有已注册的中间件列表。
// 用于调试、日志记录或检查中间件配置。
//
// 返回:
//   - []Middleware: 中间件列表（按注册顺序）
func (m *Manager) List() []Middleware {
	return m.middlewares
}

// Count 返回已注册的中间件数量。
//
// 返回:
//   - int: 中间件数量
func (m *Manager) Count() int {
	return len(m.middlewares)
}

// Clear 清空所有已注册的中间件。
// 用于重新配置中间件链。
func (m *Manager) Clear() {
	m.middlewares = make([]Middleware, 0)
}

// Remove 移除指定名称的中间件。
// 如果找到并移除成功返回 true，否则返回 false。
//
// 参数:
//   - name: 要移除的中间件名称
//
// 返回:
//   - bool: 是否成功移除
func (m *Manager) Remove(name string) bool {
	for i, mw := range m.middlewares {
		if mw.Name() == name {
			m.middlewares = append(m.middlewares[:i], m.middlewares[i+1:]...)
			return true
		}
	}
	return false
}

// Has 检查是否包含指定名称的中间件。
//
// 参数:
//   - name: 要检查的中间件名称
//
// 返回:
//   - bool: 是否包含该中间件
func (m *Manager) Has(name string) bool {
	for _, mw := range m.middlewares {
		if mw.Name() == name {
			return true
		}
	}
	return false
}

// Get 获取指定名称的中间件。
// 如果找到返回中间件实例，否则返回 nil。
//
// 参数:
//   - name: 要获取的中间件名称
//
// 返回:
//   - Middleware: 中间件实例（如果找到）
func (m *Manager) Get(name string) Middleware {
	for _, mw := range m.middlewares {
		if mw.Name() == name {
			return mw
		}
	}
	return nil
}