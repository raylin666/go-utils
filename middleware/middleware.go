// Package middleware 提供通用中间件管理系统。
package middleware

// Handler 通用中间件处理函数类型。
// 这是一个抽象类型，具体实现由框架层定义。
type Handler interface{}

// Middleware 通用中间件接口。
// 所有中间件必须实现此接口，提供名称、优先级和处理函数。
type Middleware interface {
	// Name 返回中间件名称。
	// 用于日志记录、调试和中间件识别。
	Name() string

	// Priority 返回中间件优先级。
	// 决定中间件在链中的执行顺序。
	Priority() Priority

	// Handler 返回中间件处理函数。
	// 返回类型为 Handler 接口，具体类型由框架层定义。
	Handler() Handler
}

// Config 中间件配置接口。
// 用于可配置的中间件，提供启用状态检查。
type Config interface {
	// Enabled 返回中间件是否启用。
	// 如果返回 false，中间件管理器将跳过该中间件。
	Enabled() bool
}

// middlewareFunc 函数式中间件实现。
// 提供简化的中间件创建方式，无需定义完整结构体。
type middlewareFunc struct {
	name     string
	priority Priority
	handler  Handler
}

// Name 返回中间件名称。
func (m *middlewareFunc) Name() string {
	return m.name
}

// Priority 返回中间件优先级。
func (m *middlewareFunc) Priority() Priority {
	return m.priority
}

// Handler 返回中间件处理函数。
func (m *middlewareFunc) Handler() Handler {
	return m.handler
}

// NewMiddlewareFunc 创建函数式中间件。
// 提供简化的中间件创建方式。
//
// 参数:
//   - name: 中间件名称
//   - priority: 中间件优先级
//   - handler: 中间件处理函数（具体类型由框架层定义）
//
// 返回:
//   - Middleware: 中间件实例
func NewMiddlewareFunc(name string, priority Priority, handler Handler) Middleware {
	return &middlewareFunc{
		name:     name,
		priority: priority,
		handler:  handler,
	}
}