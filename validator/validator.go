package validator

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Options struct {
	Locale string
	Translator ut.Translator
}

type Validator struct {
	Options

	Valid *validator.Validate
}

func New(locale string) *Validator {
	var translator locales.Translator
	switch locale {
	case "en":
		translator = en.New()
	break
	default:
		translator = zh.New()
	}

	// 注册翻译器
	uni := ut.New(translator)
	trans, _ := uni.GetTranslator(locale)

	// 注册验证器
	valid := validator.New()

	// 注册一个函数，获取struct tag里自定义的label作为字段名
	valid.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 验证器注册翻译器
	switch locale {
	case "en":
		_ = en_translations.RegisterDefaultTranslations(valid, trans)
		break
	default:
		_ = zh_translations.RegisterDefaultTranslations(valid, trans)
	}

	var validatorInstance = new(Validator)
	validatorInstance.Locale = locale
	validatorInstance.Translator = trans
	validatorInstance.Valid = valid
	return validatorInstance
}

// Validate 验证数据
func (v *Validator) Validate(data interface{}) string {
	err := v.Valid.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err.Translate(v.Translator)
		}
	}

	return ""
}
