package context

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/internal/constant"
	"go-server/internal/model"
	"time"
)

// 上下文结构体
type Context struct {
	*gin.Context

	// 认证 Key
	Authorization string

	// 响应数据包内容
	ResponseBuilder ResponseBuilder

	Model struct {
		JwtSecretModel *model.JwtSecretModel
		JwtUsersModel  *model.JwtUsersModel
	}
}

// 上下文处理函数
type ContextHandlerFunc func(ctx *Context)

// 设置 认证 Key
func (ctx *Context) SetContextAuthorization(value string) {
	ctx.Set("Authorization", value)
}

// 获取 认证 Key
func (ctx *Context) GetContextAuthorization() string {
	return ctx.Authorization
}

// 上下文处理
func ContextHandler(handler ContextHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context = new(Context)
		context.Context = ctx
		context.ResponseBuilder.WithRequestStartTime(time.Now())
		if authorization, ok := ctx.Keys["Authorization"]; ok {
			context.Authorization = authorization.(string)
		}

		context.Model.JwtSecretModel = model.NewJwtSecretModel()
		context.Model.JwtUsersModel = model.NewJwtUsersModel()

		handler(context)
	}
}

// API 成功响应
func (ctx *Context) Success(data interface{}) {
	ctx.ResponseBuilder.WithCode(constant.StatusOK)
	ctx.ResponseBuilder.WithData(data)
	ctx.handlerResponse()
}

// API 失败响应
func (ctx *Context) Error(code int) {
	ctx.ResponseBuilder.WithCode(code)
	ctx.ResponseBuilder.WithData(H{})
	ctx.handlerResponse()
}

// 处理响应数据包内容
func (ctx *Context) handlerResponse() {
	// 处理响应状态码
	ctx.ResponseBuilder.WithCode(ctx.ResponseBuilder.Code)
	if ctx.ResponseBuilder.Code <= 600 && ctx.ResponseBuilder.Code > 0 {
		ctx.ResponseBuilder.WithHttpCode(ctx.ResponseBuilder.Code)
	} else if ctx.ResponseBuilder.Code > 600 {
		ctx.ResponseBuilder.WithHttpCode(constant.StatusOK)
	} else {
		ctx.ResponseBuilder.WithHttpCode(ctx.ResponseBuilder.HttpCode)
	}

	ctx.ResponseBuilder.WithMessage(ctx.ResponseBuilder.Message)

	ctx.ResponseBuilder.WithRequestEndTime(time.Now())

	ctx.ResponseBuilder.WithResponseTime(ctx.ResponseBuilder.GetRequestEndTime().Sub(ctx.ResponseBuilder.GetRequestStartTime()))

	ctx.ResponseBuilder.Data = H{
		"code":         ctx.ResponseBuilder.Code,
		"message":      ctx.ResponseBuilder.Message,
		"data":         ctx.ResponseBuilder.Data,
		"responseTime": fmt.Sprintf("%s", ctx.ResponseBuilder.GetResponseTime()),
	}

	switch ctx.ResponseBuilder.Format {
	case constant.FormatXML:
		ctx.builderResponseXML()
	case constant.FormatYAML:
		ctx.builderResponseYAML()
	case constant.FormatJSON:
		ctx.builderResponseJSONP()
	default:
		ctx.builderResponseJSON()
	}
}

func (ctx *Context) builderResponseJSON() {
	ctx.Context.JSON(ctx.ResponseBuilder.HttpCode, ctx.ResponseBuilder.Data)
}

func (ctx *Context) builderResponseXML() {
	ctx.Context.XML(ctx.ResponseBuilder.HttpCode, ctx.ResponseBuilder.Data)
}

func (ctx *Context) builderResponseYAML() {
	ctx.Context.YAML(ctx.ResponseBuilder.HttpCode, ctx.ResponseBuilder.Data)
}

func (ctx *Context) builderResponseJSONP() {
	ctx.Context.JSONP(ctx.ResponseBuilder.HttpCode, ctx.ResponseBuilder.Data)
}
