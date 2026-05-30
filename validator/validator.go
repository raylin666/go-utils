// Package validator 提供结构体数据验证功能，支持多语言错误消息。
// 功能特性：
//   - 基于 go-playground/validator 实现数据验证
//   - 支持自定义字段名称提取（通过 struct tag）
//   - 支持中英文错误消息翻译
//   - 支持自定义验证规则
package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	govalidator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 编译期接口验证，确保 validator 实现了 Validator 接口
var _ Validator = (*validator)(nil)

// Validator 定义数据验证接口
// 使用场景：
//   - 验证 HTTP 请求参数
//   - 验证配置文件数据
//   - 验证数据库模型字段
type Validator interface {
	// Validate 验证数据结构
	// 参数：
	//   - data: 待验证的数据结构（必须为结构体或结构体指针）
	// 返回值：
	//   - nil: 验证通过
	//   - error: 验证失败时的错误信息（已翻译为对应语言）
	Validate(data interface{}) error
}

// validator 验证器实现
// 字段说明：
//   - locale: 错误消息语言（"en" 英文，"zh" 中文）
//   - tagname: 自定义字段名称的 struct tag 键名
//   - validate: go-playground/validator 实例
//   - translator: 错误消息翻译器
type validator struct {
	locale     string            // 语言设置
	tagname    string            // 字段标签名称
	validate   *govalidator.Validate // 验证器实例
	translator ut.Translator     // 翻译器实例
}

// Option 验证器配置选项函数
// 使用选项模式（Functional Options）进行配置
type Option func(*validator)

// WithLocale 设置错误消息语言
// 参数：
//   - locale: 语言代码，支持 "en"（英文）和 "zh"（中文）
// 使用示例：
//   v := validator.New(validator.WithLocale("zh"))
func WithLocale(locale string) Option {
	return func(v *validator) {
		v.locale = locale
	}
}

// WithTagname 设置自定义字段名称的 struct tag 键名
// 参数：
//   - tagname: struct tag 键名（默认为 "label"）
// 使用示例：
//   type User struct {
//       Name string `label:"姓名" validate:"required"`
//   }
//   v := validator.New(validator.WithTagname("label"))
func WithTagname(tagname string) Option {
	return func(v *validator) {
		v.tagname = tagname
	}
}

// New 创建验证器实例
// 功能说明：
//   - 创建配置好的验证器实例
//   - 支持自定义语言和字段标签名称
//   - 自动注册翻译器和自定义字段名称提取函数
//
// 参数：
//   - opts: 配置选项列表
//
// 返回值：
//   - Validator: 验证器实例
//
// 默认配置：
//   - 语言: 英文（"en"）
//   - 字段标签: "label"
//
// 使用示例：
//   // 创建中文验证器
//   v := validator.New(validator.WithLocale("zh"))
//
//   // 创建自定义标签验证器
//   v := validator.New(
//       validator.WithLocale("zh"),
//       validator.WithTagname("displayName"),
//   )
func New(opts ...Option) Validator {
	// 初始化默认配置
	v := &validator{
		locale:  "en",  // 默认英文
		tagname: "label", // 默认字段标签名称
	}

	// 应用所有配置选项
	for _, opt := range opts {
		opt(v)
	}

	// 注册翻译器的辅助函数
	// 参数：locales.Translator - 语言翻译器实例
	// 返回值：ut.Translator - universal-translator 翻译器实例
	registerTranslatorFn := func(translator locales.Translator) ut.Translator {
		uni := ut.New(translator)
		trans, _ := uni.GetTranslator(v.locale)
		return trans
	}

	// 创建 go-playground/validator 实例
	validate := govalidator.New()

	// 注册自定义字段名称提取函数
	// 功能：从 struct tag 中提取自定义字段名称，用于错误消息显示
	// 示例：`label:"姓名"` -> 错误消息显示 "姓名不能为空" 而非 "Name不能为空"
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// 获取指定 tag 的值，并去除逗号后的选项
		name := strings.SplitN(fld.Tag.Get(v.tagname), ",", 2)[0]
		
		// "-" 表示忽略该字段，不显示在错误消息中
		if name == "-" {
			return ""
		}

		return name
	})

	// 根据语言设置注册对应的翻译器
	var trans ut.Translator
	switch v.locale {
	case "zh":
		// 中文翻译器
		translator := zh.New()
		trans = registerTranslatorFn(translator)
		// 注册中文默认翻译
		_ = zh_translations.RegisterDefaultTranslations(validate, trans)
	default:
		// 英文翻译器（默认）
		translator := en.New()
		trans = registerTranslatorFn(translator)
		// 注册英文默认翻译
		_ = en_translations.RegisterDefaultTranslations(validate, trans)
	}

	// 保存翻译器和验证器实例
	v.translator = trans
	v.validate = validate
	
	return v
}

// Validate 验证数据结构
// 功能说明：
//   - 验证结构体字段是否符合 validate tag 定义规则
//   - 返回已翻译的错误消息（根据配置的语言）
//   - 支持多种验证规则（required、email、min、max 等）
//
// 参数：
//   - data: 待验证的数据，必须为结构体或结构体指针
//
// 返回值：
//   - nil: 所有字段验证通过
//   - error: 第一个验证失败的字段错误信息
//     错误消息已翻译为配置的语言，格式如："姓名为必填字段"
//
// 验证规则示例：
//   type User struct {
//       Name  string `label:"姓名" validate:"required"`           // 必填
//       Email string `label:"邮箱" validate:"required,email"`     // 必填且邮箱格式
//       Age   int    `label:"年龄" validate:"gte=0,lte=150"`      // 0-150之间
//       Phone string `label:"电话" validate:"omitempty,len=11"`   // 可选，但必须11位
//   }
//
// 使用示例：
//   v := validator.New(validator.WithLocale("zh"))
//   user := User{Name: "", Email: "invalid", Age: 200}
//   if err := v.Validate(user); err != nil {
//       fmt.Println("验证失败:", err.Error()) // 输出: "姓名为必填字段"
//   }
func (v *validator) Validate(data interface{}) error {
	// 执行结构体验证
	err := v.validate.Struct(data)
	if err != nil {
		// 使用安全的类型断言，避免 panic
		// ValidationErrors 是 go-playground/validator 的特定错误类型
		if verr, ok := err.(govalidator.ValidationErrors); ok {
			// 返回第一个验证失败的字段错误（已翻译）
			// 多字段失败时，只返回第一个错误，避免错误消息过长
			if len(verr) > 0 {
				return fmt.Errorf("%s", verr[0].Translate(v.translator))
			}
		}
		
		// 非 ValidationErrors 类型错误（如传入非结构体）
		// 返回原始错误信息
		return fmt.Errorf("validation failed: %w", err)
	}

	// 所有字段验证通过
	return nil
}

// ValidateAll 验证所有字段并返回所有错误
// 功能说明：
//   - 与 Validate 功能相同，但返回所有验证失败的字段错误
//   - 适用于需要一次性展示所有错误的场景
//
// 参数：
//   - data: 待验证的数据结构
//
// 返回值：
//   - nil: 所有字段验证通过
//   - error: 包含所有验证失败字段的聚合错误
//
// 使用示例：
//   v := validator.New(validator.WithLocale("zh"))
//   user := User{Name: "", Email: "invalid", Age: 200}
//   if err := v.ValidateAll(user); err != nil {
//       fmt.Println("所有验证错误:", err.Error())
//       // 输出: "姓名为必填字段; 邮箱格式错误; 年龄必须小于或等于150"
//   }
func (v *validator) ValidateAll(data interface{}) error {
	err := v.validate.Struct(data)
	if err != nil {
		if verr, ok := err.(govalidator.ValidationErrors); ok {
			if len(verr) > 0 {
				// 聚合所有错误消息
				var errMsgs []string
				for _, e := range verr {
					errMsgs = append(errMsgs, e.Translate(v.translator))
				}
				return fmt.Errorf("%s", strings.Join(errMsgs, "; "))
			}
		}
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}