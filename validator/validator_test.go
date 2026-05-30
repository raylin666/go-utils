package validator

import "testing"

// TestValidator 测试数据验证功能
// 测试内容：
//   - 中文错误消息
//   - 自定义字段标签
//   - 验证规则（required、min）
func TestValidator(t *testing.T) {
	// 创建中文验证器
	validate := New(
		WithLocale("zh"),      // 中文错误消息
		WithTagname("label"),  // 自定义字段标签
	)

	// 测试验证通过的场景
	err := validate.Validate(struct {
		Name string `label:"姓名" validate:"required,min=6"`
	}{
		Name: "raylin666", // 符合规则：必填且至少6个字符
	})

	// 验证通过时应该返回 nil
	if err != nil {
		t.Fatalf("验证失败（预期通过）: %v", err)
	}

	t.Log("验证通过: SUCCESS")
}

// TestValidatorFailed 测试验证失败场景
// 测试内容：
//   - 必填字段验证
//   - 最小长度验证
//   - 中文错误消息格式
func TestValidatorFailed(t *testing.T) {
	// 创建中文验证器
	validate := New(WithLocale("zh"))

	// 测试必填字段验证失败
	err := validate.Validate(struct {
		Name string `label:"姓名" validate:"required"`
	}{
		Name: "", // 空字符串，违反必填规则
	})

	// 验证失败时应该返回错误
	if err == nil {
		t.Fatal("验证通过（预期失败）")
	}

	// 检查错误消息是否包含中文
	expectedMsg := "姓名为必填字段"
	if err.Error() != expectedMsg {
		t.Fatalf("错误消息不匹配，期望: %s, 实际: %s", expectedMsg, err.Error())
	}

	t.Logf("验证失败测试通过: %v", err)
}

// TestValidatorMinLength 测试最小长度验证
func TestValidatorMinLength(t *testing.T) {
	validate := New(WithLocale("zh"))

	// 测试最小长度验证失败
	err := validate.Validate(struct {
		Name string `label:"用户名" validate:"min=6"`
	}{
		Name: "abc", // 只有3个字符，违反最小长度规则
	})

	if err == nil {
		t.Fatal("验证通过（预期失败）")
	}

	t.Logf("最小长度验证测试通过: %v", err)
}

// TestValidatorEmail 测试邮箱格式验证
func TestValidatorEmail(t *testing.T) {
	validate := New(WithLocale("zh"))

	// 测试邮箱格式验证失败
	err := validate.Validate(struct {
		Email string `label:"邮箱" validate:"required,email"`
	}{
		Email: "invalid-email", // 无效邮箱格式
	})

	if err == nil {
		t.Fatal("验证通过（预期失败）")
	}

	t.Logf("邮箱验证测试通过: %v", err)
}