package context

import (
	"encoding/xml"
	"go-server/internal/constant"
	"time"
)

// 响应数据包内容
type H map[string]interface{}

// 构建响应结构体
type ResponseBuilder struct {
	// HTTP 状态码
	HttpCode int
	// 业务状态码
	Code int
	// 状态码提示信息
	Message string
	// 响应内容信息
	Data interface{}
	// 请求开始时间
	RequestStartTime time.Time
	// 请求结束时间
	RequestEndTime time.Time
	// 响应总时长
	ResponseTime time.Duration
	// 响应格式
	Format string
}

// MarshalXML allows type H to be used with xml.Marshal.
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (response *ResponseBuilder) WithHttpCode(code int) *ResponseBuilder {
	if code == 0 {
		response.HttpCode = constant.StatusOK
	} else {
		response.HttpCode = code
	}

	return response
}

func (response *ResponseBuilder) GetHttpCode() int {
	return response.HttpCode
}

func (response *ResponseBuilder) WithCode(code int) *ResponseBuilder {
	if code == 0 {
		response.Code = constant.StatusOK
	} else {
		response.Code = code
	}

	return response
}

func (response *ResponseBuilder) GetCode() int {
	return response.Code
}

func (response *ResponseBuilder) WithMessage(message string) *ResponseBuilder {
	if message == "" {
		response.Message = constant.GetStatusText(response.Code)
	} else {
		response.Message = message
	}

	return response
}

func (response *ResponseBuilder) GetMessage() string {
	return response.Message
}

func (response *ResponseBuilder) WithData(data interface{}) *ResponseBuilder {
	response.Data = data
	return response
}

func (response *ResponseBuilder) GetData() interface{} {
	return response.Data
}

func (response *ResponseBuilder) WithResponseTime(duration time.Duration) *ResponseBuilder {
	response.ResponseTime = duration
	return response
}

func (response *ResponseBuilder) GetResponseTime() time.Duration {
	return response.ResponseTime
}

func (response *ResponseBuilder) WithFormat(format string) *ResponseBuilder {
	switch format {
	case constant.FormatJSON:
	case constant.FormatJSONP:
	case constant.FormatXML:
	case constant.FormatYAML:
	default:
		format = constant.FormatJSON
	}
	response.Format = format
	return response
}

func (response *ResponseBuilder) GetFormat() string {
	return response.Format
}

func (response *ResponseBuilder) WithRequestStartTime(duration time.Time) *ResponseBuilder {
	response.RequestStartTime = duration
	return response
}

func (response *ResponseBuilder) GetRequestStartTime() time.Time {
	return response.RequestStartTime
}

func (response *ResponseBuilder) WithRequestEndTime(duration time.Time) *ResponseBuilder {
	response.RequestEndTime = duration
	return response
}

func (response *ResponseBuilder) GetRequestEndTime() time.Time {
	return response.RequestEndTime
}

