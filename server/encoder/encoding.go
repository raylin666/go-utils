package encoder

import (
	"strings"
)

// Codec defines the interface Transport uses to encode and decode messages.  Note
// that implementations of this interface must be thread safe; a Codec's
// methods can be called from concurrent goroutines.
type EncodingCodec interface {
	// Marshal returns the wire format of v.
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v interface{}) error
	// Name returns the name of the Codec implementation. The returned string
	// will be used as part of content type in transmission.  The result must be
	// static; the result cannot change between calls.
	Name() string
}

var encodingRegisteredCodecs = make(map[string]EncodingCodec)

// EncodingCodec registers the provided Codec for use with all Transport clients and
// servers.
func RegisterEncodingCodec(codec EncodingCodec) {
	if codec == nil {
		panic("cannot register a nil Codec")
	}
	if codec.Name() == "" {
		panic("cannot register Codec with empty string result for Name()")
	}
	contentSubtype := strings.ToLower(codec.Name())
	encodingRegisteredCodecs[contentSubtype] = codec
}

// GetEncodingCodec gets a registered Codec by content-subtype, or nil if no Codec is
// registered for the content-subtype.
//
// The content-subtype is expected to be lowercase.
func GetEncodingCodec(contentSubtype string) EncodingCodec {
	return encodingRegisteredCodecs[contentSubtype]
}
