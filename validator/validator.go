package validator

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	govalidator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var _ Validator = (*validator)(nil)

type Validator interface {
	Validate(data interface{}) string
}

type validator struct {
	locale     string
	tagname    string
	validate   *govalidator.Validate
	translator ut.Translator
}

type Option func(*validator)

func WithLocale(locale string) Option {
	return func(v *validator) {
		v.locale = locale
	}
}

func WithTagname(tagname string) Option {
	return func(v *validator) {
		v.tagname = tagname
	}
}

func New(opts ...Option) Validator {
	var v = &validator{
		locale:  "en",
		tagname: "label",
	}
	for _, opt := range opts {
		opt(v)
	}

	// 注册翻译器
	var registerTranslatorFn = func(translator locales.Translator) ut.Translator {
		uni := ut.New(translator)
		trans, _ := uni.GetTranslator(v.locale)
		return trans
	}

	// 注册验证器
	var validate = govalidator.New()
	// 注册一个函数，获取 struct tag 里自定义的字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get(v.tagname), ",", 2)[0]
		if name == "-" {
			return "" // 返回空字符串表示忽略该字段
		}

		return name
	})

	var trans ut.Translator
	switch v.locale {
	case "zh":
		translator := zh.New()
		trans = registerTranslatorFn(translator)
		// 验证器注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(validate, trans)
		break
	default:
		translator := en.New()
		trans = registerTranslatorFn(translator)
		// 验证器注册翻译器
		_ = en_translations.RegisterDefaultTranslations(validate, trans)
	}

	v.translator = trans
	v.validate = validate
	return v
}

// Validate 验证数据
func (v *validator) Validate(data interface{}) string {
	err := v.validate.Struct(data)
	if err != nil {
		for _, err := range err.(govalidator.ValidationErrors) {
			return err.Translate(v.translator)
		}
	}

	return ""
}
