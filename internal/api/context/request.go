package context

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go-server/internal/constant"
	"reflect"
)

// 请求数据验证
func (ctx *Context) RequestValidate(validate interface{}) bool {
	// 注册翻译器
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	// 注册验证器
	valid := validator.New()

	//注册一个函数，获取struct tag里自定义的label作为字段名
	valid.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(valid, trans)
	if err != nil {
		ctx.Error(constant.StatusBsValidationHandleError)
		return false
	}

	err = valid.Struct(validate)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ctx.ResponseBuilder.WithMessage(err.Translate(trans))
			ctx.Error(constant.StatusUnprocessableEntity)
			return false
		}
	}

	return true
}
