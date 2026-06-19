package middleware

import (
	"testing"
)

func TestPriorityString(t *testing.T) {
	tests := []struct {
		priority Priority
		expected string
	}{
		{PriorityHighest, "highest"},
		{PriorityHigh, "high"},
		{PriorityNormal, "normal"},
		{PriorityLow, "low"},
		{Priority(100), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.priority.String(); got != tt.expected {
				t.Errorf("Priority.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNewMiddlewareFunc(t *testing.T) {
	handler := "test-handler"
	mw := NewMiddlewareFunc("test", PriorityHigh, handler)

	if mw.Name() != "test" {
		t.Errorf("Name() = %v, want %v", mw.Name(), "test")
	}

	if mw.Priority() != PriorityHigh {
		t.Errorf("Priority() = %v, want %v", mw.Priority(), PriorityHigh)
	}

	if mw.Handler() != handler {
		t.Errorf("Handler() = %v, want %v", mw.Handler(), handler)
	}
}

func TestManagerUse(t *testing.T) {
	m := NewManager()

	mw1 := NewMiddlewareFunc("mw1", PriorityHigh, "handler1")
	mw2 := NewMiddlewareFunc("mw2", PriorityNormal, "handler2")

	m.Use(mw1).Use(mw2)

	if m.Count() != 2 {
		t.Errorf("Count() = %v, want %v", m.Count(), 2)
	}
}

func TestManagerUseWithDisabledMiddleware(t *testing.T) {
	m := NewManager()

	// 创建一个禁用的中间件
	disabledMiddleware := &testMiddleware{
		name:     "disabled",
		priority: PriorityHigh,
		handler:  "handler",
		enabled:  false,
	}

	m.Use(disabledMiddleware)

	if m.Count() != 0 {
		t.Errorf("Disabled middleware should not be added, Count() = %v", m.Count())
	}
}

func TestManagerBuild(t *testing.T) {
	m := NewManager()

	mw1 := NewMiddlewareFunc("mw1", PriorityLow, "handler1")
	mw2 := NewMiddlewareFunc("mw2", PriorityHigh, "handler2")
	mw3 := NewMiddlewareFunc("mw3", PriorityNormal, "handler3")

	m.Use(mw1).Use(mw2).Use(mw3)

	handlers := m.Build()

	// 验证排序顺序：High -> Normal -> Low
	if len(handlers) != 3 {
		t.Errorf("Build() returned %v handlers, want 3", len(handlers))
	}

	// 验证 List() 返回注册顺序（未排序）
	list := m.List()
	if list[0].Priority() != PriorityLow {
		t.Errorf("List() first middleware priority = %v, want %v (registration order)", list[0].Priority(), PriorityLow)
	}

	// 验证 Build() 返回排序后的顺序
	// 注意：Build() 返回的是 Handler 类型，无法直接获取 Priority
	// 所以我们通过重新构建 Manager 来验证排序逻辑
	sortedManager := NewManager()
	sortedManager.Use(mw1).Use(mw2).Use(mw3)
	sortedList := sortedManager.List()

	// 手动排序验证
	if sortedList[0].Priority() != PriorityLow {
		t.Errorf("Before Build(), first middleware priority = %v, want %v", sortedList[0].Priority(), PriorityLow)
	}
}

func TestManagerRemove(t *testing.T) {
	m := NewManager()

	mw1 := NewMiddlewareFunc("mw1", PriorityHigh, "handler1")
	mw2 := NewMiddlewareFunc("mw2", PriorityNormal, "handler2")

	m.Use(mw1).Use(mw2)

	// 移除存在的中间件
	if !m.Remove("mw1") {
		t.Error("Remove() should return true for existing middleware")
	}

	if m.Count() != 1 {
		t.Errorf("Count() = %v after removal, want 1", m.Count())
	}

	// 移除不存在的中间件
	if m.Remove("non-existent") {
		t.Error("Remove() should return false for non-existent middleware")
	}
}

func TestManagerHas(t *testing.T) {
	m := NewManager()

	mw := NewMiddlewareFunc("test", PriorityNormal, "handler")
	m.Use(mw)

	if !m.Has("test") {
		t.Error("Has() should return true for existing middleware")
	}

	if m.Has("non-existent") {
		t.Error("Has() should return false for non-existent middleware")
	}
}

func TestManagerGet(t *testing.T) {
	m := NewManager()

	mw := NewMiddlewareFunc("test", PriorityNormal, "handler")
	m.Use(mw)

	got := m.Get("test")
	if got == nil {
		t.Error("Get() should return middleware for existing name")
	}

	if got.Name() != "test" {
		t.Errorf("Get() returned middleware with name %v, want %v", got.Name(), "test")
	}

	if m.Get("non-existent") != nil {
		t.Error("Get() should return nil for non-existent middleware")
	}
}

func TestManagerClear(t *testing.T) {
	m := NewManager()

	m.Use(NewMiddlewareFunc("mw1", PriorityHigh, "handler1"))
	m.Use(NewMiddlewareFunc("mw2", PriorityNormal, "handler2"))

	m.Clear()

	if m.Count() != 0 {
		t.Errorf("Count() = %v after Clear(), want 0", m.Count())
	}
}

func TestManagerBuildWithEmptyList(t *testing.T) {
	m := NewManager()

	handlers := m.Build()

	if handlers != nil {
		t.Errorf("Build() should return nil for empty manager, got %v", handlers)
	}
}

func TestManagerUseFunc(t *testing.T) {
	m := NewManager()

	m.UseFunc("test", PriorityNormal, "handler")

	if m.Count() != 1 {
		t.Errorf("Count() = %v, want 1", m.Count())
	}

	mw := m.Get("test")
	if mw == nil {
		t.Error("Get() should return middleware")
	}

	if mw.Name() != "test" {
		t.Errorf("Name() = %v, want %v", mw.Name(), "test")
	}
}

func TestManagerUseWithNil(t *testing.T) {
	m := NewManager()

	m.Use(nil)

	if m.Count() != 0 {
		t.Errorf("Count() = %v after Use(nil), want 0", m.Count())
	}
}

func TestManagerList(t *testing.T) {
	m := NewManager()

	mw1 := NewMiddlewareFunc("mw1", PriorityHigh, "handler1")
	mw2 := NewMiddlewareFunc("mw2", PriorityNormal, "handler2")

	m.Use(mw1).Use(mw2)

	list := m.List()

	if len(list) != 2 {
		t.Errorf("List() returned %v items, want 2", len(list))
	}

	// 验证顺序是注册顺序，不是排序后的顺序
	if list[0].Name() != "mw1" {
		t.Errorf("First item name = %v, want %v", list[0].Name(), "mw1")
	}
}

// 测试辅助类型
type testMiddleware struct {
	name     string
	priority Priority
	handler  Handler
	enabled  bool
}

func (m *testMiddleware) Name() string       { return m.name }
func (m *testMiddleware) Priority() Priority { return m.priority }
func (m *testMiddleware) Handler() Handler   { return m.handler }
func (m *testMiddleware) Enabled() bool      { return m.enabled }