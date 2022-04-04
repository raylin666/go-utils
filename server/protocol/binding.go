package protocol

import (
	"google.golang.org/grpc/encoding"
	"net/http"
	"net/url"
)

// BindQuery bind vars parameters to target.
func BindQuery(vars url.Values, target interface{}) error {
	return encoding.GetCodec(FormName).Unmarshal([]byte(vars.Encode()), target)
}

// BindForm bind form parameters to target.
func BindForm(req *http.Request, target interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	return encoding.GetCodec(FormName).Unmarshal([]byte(req.Form.Encode()), target)
}
