// Package middleware 提供通用中间件管理系统。
package middleware

// Priority 中间件优先级类型。
// 数值越小优先级越高，中间件按优先级顺序执行。
type Priority int

const (
	// PriorityHighest 最高优先级（数值最小）。
	// 用于必须在最前执行的中间件，如异常恢复、链路追踪。
	PriorityHighest Priority = iota

	// PriorityHigh 高优先级。
	// 用于需要在早期执行的中间件，如 CORS、安全检查。
	PriorityHigh

	// PriorityNormal 正常优先级。
	// 用于常规中间件，如日志、验证、请求处理。
	PriorityNormal

	// PriorityLow 低优先级。
	// 用于业务相关的中间件，如权限检查、限流。
	PriorityLow
)

// String 返回优先级的字符串表示。
func (p Priority) String() string {
	switch p {
	case PriorityHighest:
		return "highest"
	case PriorityHigh:
		return "high"
	case PriorityNormal:
		return "normal"
	case PriorityLow:
		return "low"
	default:
		return "unknown"
	}
}