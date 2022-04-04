package protocol

import (
	"github.com/go-playground/form/v4"
	rpcproto "github.com/raylin666/go-utils/rpc/proto"
	server_encoder "github.com/raylin666/go-utils/server/encoder"
	"google.golang.org/protobuf/proto"
	"net/url"
	"reflect"
)

// FormName is form codec name
const FormName = "x-www-form-urlencoded"

func init() {
	decoder := form.NewDecoder()
	decoder.SetTagName("json")
	encoder := form.NewEncoder()
	encoder.SetTagName("json")
	server_encoder.RegisterEncodingCodec(formCodec{encoder: encoder, decoder: decoder})
}

type formCodec struct {
	encoder *form.Encoder
	decoder *form.Decoder
}

func (c formCodec) Marshal(v interface{}) ([]byte, error) {
	var vs url.Values
	var err error
	if m, ok := v.(proto.Message); ok {
		vs, err = rpcproto.EncodeMap(m)
		if err != nil {
			return nil, err
		}
	} else {
		vs, err = c.encoder.Encode(v)
		if err != nil {
			return nil, err
		}
	}
	for k, v := range vs {
		if len(v) == 0 {
			delete(vs, k)
		}
	}
	return []byte(vs.Encode()), nil
}

func (c formCodec) Unmarshal(data []byte, v interface{}) error {
	vs, err := url.ParseQuery(string(data))
	if err != nil {
		return err
	}

	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}
	if m, ok := v.(proto.Message); ok {
		return rpcproto.MapProto(m, vs)
	} else if m, ok := reflect.Indirect(reflect.ValueOf(v)).Interface().(proto.Message); ok {
		return rpcproto.MapProto(m, vs)
	}

	return c.decoder.Decode(v, vs)
}

func (formCodec) Name() string {
	return FormName
}
