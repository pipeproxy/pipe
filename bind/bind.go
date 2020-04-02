// DO NOT EDIT! Code generated.

package bind

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wzshiming/funcfg/kinder"
	"github.com/wzshiming/funcfg/types"
	"github.com/wzshiming/funcfg/unmarshaler"
)

// ========= Begin Common =========
//

var kindKey = "@Kind"

var defTypes = types.NewTypes()

// Unmarshal parses the encoded data and stores the result
func Unmarshal(config []byte, v interface{}) error {
	u := unmarshaler.Unmarshaler{
		Ctx:  context.Background(),
		Get:  defTypes.Get,
		Kind: kinder.Kind,
	}
	return u.Unmarshal(config, v)
}

// Marshal returns the encoding of v
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Component is basic definition of Component
type Component interface {
	isComponent()
}

// RawComponent is store raw bytes of Component
type RawComponent []byte

func (RawComponent) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawComponent) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawComponent) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawComponent: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

func appendKV(k, v string, data []byte) []byte {
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"%s\":%q}", k, v))
		} else {
			data = append([]byte(fmt.Sprintf("{\"%s\":%q,", k, v)), data[1:]...)
		}
	}
	return data
}

//
// ========= End Common =========

// ========= Begin tls.TLS =========
//

// TLS tls.TLS
type TLS interface {
	isTLS()
	Component
}

// RawTLS is store raw bytes of TLS
type RawTLS []byte

func (RawTLS) isTLS()       {}
func (RawTLS) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawTLS) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawTLS) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawTLS: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End tls.TLS =========

// ========= Begin acme@tls.TLS =========
//

const kindAcmeTLSConfig = "acme@tls.TLS"

// AcmeTLSConfig acme@tls.TLS
type AcmeTLSConfig struct {
	Domains  []string
	CacheDir string
}

func init() {
	_ = defTypes.Register(
		kindAcmeTLSConfig,
		func(r *AcmeTLSConfig) TLS {
			return r
		},
	)
}

func (AcmeTLSConfig) isTLS()       {}
func (AcmeTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m AcmeTLSConfig) MarshalJSON() ([]byte, error) {
	type t AcmeTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindAcmeTLSConfig, data)
	return data, nil
}

//
// ========= End acme@tls.TLS =========

// ========= Begin http.Handler =========
//

// HTTPHandler http.Handler
type HTTPHandler interface {
	isHTTPHandler()
	Component
}

// RawHTTPHandler is store raw bytes of HTTPHandler
type RawHTTPHandler []byte

func (RawHTTPHandler) isHTTPHandler() {}
func (RawHTTPHandler) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawHTTPHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawHTTPHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawHTTPHandler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End http.Handler =========

// ========= Begin add_request_header@net/http.Handler =========
//

const kindAddRequestHeaderNetHTTPHandlerConfig = "add_request_header@net/http.Handler"

// AddRequestHeaderNetHTTPHandlerConfig add_request_header@net/http.Handler
type AddRequestHeaderNetHTTPHandlerConfig struct {
	Key   string
	Value string
}

func init() {
	_ = defTypes.Register(
		kindAddRequestHeaderNetHTTPHandlerConfig,
		func(r *AddRequestHeaderNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (AddRequestHeaderNetHTTPHandlerConfig) isHTTPHandler() {}
func (AddRequestHeaderNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m AddRequestHeaderNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t AddRequestHeaderNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindAddRequestHeaderNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End add_request_header@net/http.Handler =========

// ========= Begin add_response_header@net/http.Handler =========
//

const kindAddResponseHeaderNetHTTPHandlerConfig = "add_response_header@net/http.Handler"

// AddResponseHeaderNetHTTPHandlerConfig add_response_header@net/http.Handler
type AddResponseHeaderNetHTTPHandlerConfig struct {
	Key   string
	Value string
}

func init() {
	_ = defTypes.Register(
		kindAddResponseHeaderNetHTTPHandlerConfig,
		func(r *AddResponseHeaderNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (AddResponseHeaderNetHTTPHandlerConfig) isHTTPHandler() {}
func (AddResponseHeaderNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m AddResponseHeaderNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t AddResponseHeaderNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindAddResponseHeaderNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End add_response_header@net/http.Handler =========

// ========= Begin codec.Decoder =========
//

// CodecDecoder codec.Decoder
type CodecDecoder interface {
	isCodecDecoder()
	Component
}

// RawCodecDecoder is store raw bytes of CodecDecoder
type RawCodecDecoder []byte

func (RawCodecDecoder) isCodecDecoder() {}
func (RawCodecDecoder) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawCodecDecoder) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawCodecDecoder) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawCodecDecoder: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End codec.Decoder =========

// ========= Begin base32@codec.Decoder =========
//

const kindBase32CodecDecoderConfig = "base32@codec.Decoder"

// Base32CodecDecoderConfig base32@codec.Decoder
type Base32CodecDecoderConfig struct {
	Encoding string
}

func init() {
	_ = defTypes.Register(
		kindBase32CodecDecoderConfig,
		func(r *Base32CodecDecoderConfig) CodecDecoder {
			return r
		},
	)
}

func (Base32CodecDecoderConfig) isCodecDecoder() {}
func (Base32CodecDecoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m Base32CodecDecoderConfig) MarshalJSON() ([]byte, error) {
	type t Base32CodecDecoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindBase32CodecDecoderConfig, data)
	return data, nil
}

//
// ========= End base32@codec.Decoder =========

// ========= Begin codec.Encoder =========
//

// CodecEncoder codec.Encoder
type CodecEncoder interface {
	isCodecEncoder()
	Component
}

// RawCodecEncoder is store raw bytes of CodecEncoder
type RawCodecEncoder []byte

func (RawCodecEncoder) isCodecEncoder() {}
func (RawCodecEncoder) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawCodecEncoder) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawCodecEncoder) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawCodecEncoder: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End codec.Encoder =========

// ========= Begin base32@codec.Encoder =========
//

const kindBase32CodecEncoderConfig = "base32@codec.Encoder"

// Base32CodecEncoderConfig base32@codec.Encoder
type Base32CodecEncoderConfig struct {
	Encoding string
}

func init() {
	_ = defTypes.Register(
		kindBase32CodecEncoderConfig,
		func(r *Base32CodecEncoderConfig) CodecEncoder {
			return r
		},
	)
}

func (Base32CodecEncoderConfig) isCodecEncoder() {}
func (Base32CodecEncoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m Base32CodecEncoderConfig) MarshalJSON() ([]byte, error) {
	type t Base32CodecEncoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindBase32CodecEncoderConfig, data)
	return data, nil
}

//
// ========= End base32@codec.Encoder =========

// ========= Begin base64@codec.Decoder =========
//

const kindBase64CodecDecoderConfig = "base64@codec.Decoder"

// Base64CodecDecoderConfig base64@codec.Decoder
type Base64CodecDecoderConfig struct {
	Encoding string
}

func init() {
	_ = defTypes.Register(
		kindBase64CodecDecoderConfig,
		func(r *Base64CodecDecoderConfig) CodecDecoder {
			return r
		},
	)
}

func (Base64CodecDecoderConfig) isCodecDecoder() {}
func (Base64CodecDecoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m Base64CodecDecoderConfig) MarshalJSON() ([]byte, error) {
	type t Base64CodecDecoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindBase64CodecDecoderConfig, data)
	return data, nil
}

//
// ========= End base64@codec.Decoder =========

// ========= Begin base64@codec.Encoder =========
//

const kindBase64CodecEncoderConfig = "base64@codec.Encoder"

// Base64CodecEncoderConfig base64@codec.Encoder
type Base64CodecEncoderConfig struct {
	Encoding string
}

func init() {
	_ = defTypes.Register(
		kindBase64CodecEncoderConfig,
		func(r *Base64CodecEncoderConfig) CodecEncoder {
			return r
		},
	)
}

func (Base64CodecEncoderConfig) isCodecEncoder() {}
func (Base64CodecEncoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m Base64CodecEncoderConfig) MarshalJSON() ([]byte, error) {
	type t Base64CodecEncoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindBase64CodecEncoderConfig, data)
	return data, nil
}

//
// ========= End base64@codec.Encoder =========

// ========= Begin bzip2@codec.Decoder =========
//

const kindBzip2CodecDecoder = "bzip2@codec.Decoder"

// Bzip2CodecDecoder bzip2@codec.Decoder
type Bzip2CodecDecoder struct {
}

func init() {
	_ = defTypes.Register(
		kindBzip2CodecDecoder,
		func(r *Bzip2CodecDecoder) CodecDecoder {
			return r
		},
	)
}

func (Bzip2CodecDecoder) isCodecDecoder() {}
func (Bzip2CodecDecoder) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m Bzip2CodecDecoder) MarshalJSON() ([]byte, error) {
	type t Bzip2CodecDecoder
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindBzip2CodecDecoder, data)
	return data, nil
}

//
// ========= End bzip2@codec.Decoder =========

// ========= Begin compress@net/http.Handler =========
//

const kindCompressNetHTTPHandlerConfig = "compress@net/http.Handler"

// CompressNetHTTPHandlerConfig compress@net/http.Handler
type CompressNetHTTPHandlerConfig struct {
	Level   int
	Handler HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindCompressNetHTTPHandlerConfig,
		func(r *CompressNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (CompressNetHTTPHandlerConfig) isHTTPHandler() {}
func (CompressNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CompressNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t CompressNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindCompressNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End compress@net/http.Handler =========

// ========= Begin config_dump@net/http.Handler =========
//

const kindConfigDumpNetHTTPHandler = "config_dump@net/http.Handler"

// ConfigDumpNetHTTPHandler config_dump@net/http.Handler
type ConfigDumpNetHTTPHandler struct {
}

func init() {
	_ = defTypes.Register(
		kindConfigDumpNetHTTPHandler,
		func(r *ConfigDumpNetHTTPHandler) HTTPHandler {
			return r
		},
	)
}

func (ConfigDumpNetHTTPHandler) isHTTPHandler() {}
func (ConfigDumpNetHTTPHandler) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ConfigDumpNetHTTPHandler) MarshalJSON() ([]byte, error) {
	type t ConfigDumpNetHTTPHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindConfigDumpNetHTTPHandler, data)
	return data, nil
}

//
// ========= End config_dump@net/http.Handler =========

// ========= Begin def@codec.Decoder =========
//

const kindDefCodecDecoderConfig = "def@codec.Decoder"

// DefCodecDecoderConfig def@codec.Decoder
type DefCodecDecoderConfig struct {
	Name string
	Def  CodecDecoder
}

func init() {
	_ = defTypes.Register(
		kindDefCodecDecoderConfig,
		func(r *DefCodecDecoderConfig) CodecDecoder {
			return r
		},
	)
}

func (DefCodecDecoderConfig) isCodecDecoder() {}
func (DefCodecDecoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefCodecDecoderConfig) MarshalJSON() ([]byte, error) {
	type t DefCodecDecoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefCodecDecoderConfig, data)
	return data, nil
}

//
// ========= End def@codec.Decoder =========

// ========= Begin def@codec.Encoder =========
//

const kindDefCodecEncoderConfig = "def@codec.Encoder"

// DefCodecEncoderConfig def@codec.Encoder
type DefCodecEncoderConfig struct {
	Name string
	Def  CodecEncoder
}

func init() {
	_ = defTypes.Register(
		kindDefCodecEncoderConfig,
		func(r *DefCodecEncoderConfig) CodecEncoder {
			return r
		},
	)
}

func (DefCodecEncoderConfig) isCodecEncoder() {}
func (DefCodecEncoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefCodecEncoderConfig) MarshalJSON() ([]byte, error) {
	type t DefCodecEncoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefCodecEncoderConfig, data)
	return data, nil
}

//
// ========= End def@codec.Encoder =========

// ========= Begin codec.Marshaler =========
//

// CodecMarshaler codec.Marshaler
type CodecMarshaler interface {
	isCodecMarshaler()
	Component
}

// RawCodecMarshaler is store raw bytes of CodecMarshaler
type RawCodecMarshaler []byte

func (RawCodecMarshaler) isCodecMarshaler() {}
func (RawCodecMarshaler) isComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawCodecMarshaler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawCodecMarshaler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawCodecMarshaler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End codec.Marshaler =========

// ========= Begin def@codec.Marshaler =========
//

const kindDefCodecMarshalerConfig = "def@codec.Marshaler"

// DefCodecMarshalerConfig def@codec.Marshaler
type DefCodecMarshalerConfig struct {
	Name string
	Def  CodecMarshaler
}

func init() {
	_ = defTypes.Register(
		kindDefCodecMarshalerConfig,
		func(r *DefCodecMarshalerConfig) CodecMarshaler {
			return r
		},
	)
}

func (DefCodecMarshalerConfig) isCodecMarshaler() {}
func (DefCodecMarshalerConfig) isComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefCodecMarshalerConfig) MarshalJSON() ([]byte, error) {
	type t DefCodecMarshalerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefCodecMarshalerConfig, data)
	return data, nil
}

//
// ========= End def@codec.Marshaler =========

// ========= Begin codec.Unmarshaler =========
//

// CodecUnmarshaler codec.Unmarshaler
type CodecUnmarshaler interface {
	isCodecUnmarshaler()
	Component
}

// RawCodecUnmarshaler is store raw bytes of CodecUnmarshaler
type RawCodecUnmarshaler []byte

func (RawCodecUnmarshaler) isCodecUnmarshaler() {}
func (RawCodecUnmarshaler) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawCodecUnmarshaler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawCodecUnmarshaler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawCodecUnmarshaler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End codec.Unmarshaler =========

// ========= Begin def@codec.Unmarshaler =========
//

const kindDefCodecUnmarshalerConfig = "def@codec.Unmarshaler"

// DefCodecUnmarshalerConfig def@codec.Unmarshaler
type DefCodecUnmarshalerConfig struct {
	Name string
	Def  CodecUnmarshaler
}

func init() {
	_ = defTypes.Register(
		kindDefCodecUnmarshalerConfig,
		func(r *DefCodecUnmarshalerConfig) CodecUnmarshaler {
			return r
		},
	)
}

func (DefCodecUnmarshalerConfig) isCodecUnmarshaler() {}
func (DefCodecUnmarshalerConfig) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefCodecUnmarshalerConfig) MarshalJSON() ([]byte, error) {
	type t DefCodecUnmarshalerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefCodecUnmarshalerConfig, data)
	return data, nil
}

//
// ========= End def@codec.Unmarshaler =========

// ========= Begin io.Reader =========
//

// IoReader io.Reader
type IoReader interface {
	isIoReader()
	Component
}

// RawIoReader is store raw bytes of IoReader
type RawIoReader []byte

func (RawIoReader) isIoReader()  {}
func (RawIoReader) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawIoReader) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawIoReader) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawIoReader: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End io.Reader =========

// ========= Begin def@io.Reader =========
//

const kindDefIoReaderConfig = "def@io.Reader"

// DefIoReaderConfig def@io.Reader
type DefIoReaderConfig struct {
	Name string
	Def  IoReader
}

func init() {
	_ = defTypes.Register(
		kindDefIoReaderConfig,
		func(r *DefIoReaderConfig) IoReader {
			return r
		},
	)
}

func (DefIoReaderConfig) isIoReader()  {}
func (DefIoReaderConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefIoReaderConfig) MarshalJSON() ([]byte, error) {
	type t DefIoReaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefIoReaderConfig, data)
	return data, nil
}

//
// ========= End def@io.Reader =========

// ========= Begin io.Writer =========
//

// IoWriter io.Writer
type IoWriter interface {
	isIoWriter()
	Component
}

// RawIoWriter is store raw bytes of IoWriter
type RawIoWriter []byte

func (RawIoWriter) isIoWriter()  {}
func (RawIoWriter) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawIoWriter) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawIoWriter) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawIoWriter: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End io.Writer =========

// ========= Begin def@io.Writer =========
//

const kindDefIoWriterConfig = "def@io.Writer"

// DefIoWriterConfig def@io.Writer
type DefIoWriterConfig struct {
	Name string
	Def  IoWriter
}

func init() {
	_ = defTypes.Register(
		kindDefIoWriterConfig,
		func(r *DefIoWriterConfig) IoWriter {
			return r
		},
	)
}

func (DefIoWriterConfig) isIoWriter()  {}
func (DefIoWriterConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefIoWriterConfig) MarshalJSON() ([]byte, error) {
	type t DefIoWriterConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefIoWriterConfig, data)
	return data, nil
}

//
// ========= End def@io.Writer =========

// ========= Begin def@net/http.Handler =========
//

const kindDefNetHTTPHandlerConfig = "def@net/http.Handler"

// DefNetHTTPHandlerConfig def@net/http.Handler
type DefNetHTTPHandlerConfig struct {
	Name string
	Def  HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindDefNetHTTPHandlerConfig,
		func(r *DefNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (DefNetHTTPHandlerConfig) isHTTPHandler() {}
func (DefNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t DefNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End def@net/http.Handler =========

// ========= Begin http.RoundTripper =========
//

// HTTPRoundTripper http.RoundTripper
type HTTPRoundTripper interface {
	isHTTPRoundTripper()
	Component
}

// RawHTTPRoundTripper is store raw bytes of HTTPRoundTripper
type RawHTTPRoundTripper []byte

func (RawHTTPRoundTripper) isHTTPRoundTripper() {}
func (RawHTTPRoundTripper) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawHTTPRoundTripper) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawHTTPRoundTripper) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawHTTPRoundTripper: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End http.RoundTripper =========

// ========= Begin def@net/http.RoundTripper =========
//

const kindDefNetHTTPRoundTripperConfig = "def@net/http.RoundTripper"

// DefNetHTTPRoundTripperConfig def@net/http.RoundTripper
type DefNetHTTPRoundTripperConfig struct {
	Name string
	Def  HTTPRoundTripper
}

func init() {
	_ = defTypes.Register(
		kindDefNetHTTPRoundTripperConfig,
		func(r *DefNetHTTPRoundTripperConfig) HTTPRoundTripper {
			return r
		},
	)
}

func (DefNetHTTPRoundTripperConfig) isHTTPRoundTripper() {}
func (DefNetHTTPRoundTripperConfig) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefNetHTTPRoundTripperConfig) MarshalJSON() ([]byte, error) {
	type t DefNetHTTPRoundTripperConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefNetHTTPRoundTripperConfig, data)
	return data, nil
}

//
// ========= End def@net/http.RoundTripper =========

// ========= Begin once.Once =========
//

// Once once.Once
type Once interface {
	isOnce()
	Component
}

// RawOnce is store raw bytes of Once
type RawOnce []byte

func (RawOnce) isOnce()      {}
func (RawOnce) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawOnce) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawOnce) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawOnce: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End once.Once =========

// ========= Begin def@once.Once =========
//

const kindDefOnceConfig = "def@once.Once"

// DefOnceConfig def@once.Once
type DefOnceConfig struct {
	Name string
	Def  Once
}

func init() {
	_ = defTypes.Register(
		kindDefOnceConfig,
		func(r *DefOnceConfig) Once {
			return r
		},
	)
}

func (DefOnceConfig) isOnce()      {}
func (DefOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefOnceConfig) MarshalJSON() ([]byte, error) {
	type t DefOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefOnceConfig, data)
	return data, nil
}

//
// ========= End def@once.Once =========

// ========= Begin service.Service =========
//

// Service service.Service
type Service interface {
	isService()
	Component
}

// RawService is store raw bytes of Service
type RawService []byte

func (RawService) isService()   {}
func (RawService) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawService) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawService) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawService: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End service.Service =========

// ========= Begin def@service.Service =========
//

const kindDefServiceConfig = "def@service.Service"

// DefServiceConfig def@service.Service
type DefServiceConfig struct {
	Name string
	Def  Service
}

func init() {
	_ = defTypes.Register(
		kindDefServiceConfig,
		func(r *DefServiceConfig) Service {
			return r
		},
	)
}

func (DefServiceConfig) isService()   {}
func (DefServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefServiceConfig) MarshalJSON() ([]byte, error) {
	type t DefServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefServiceConfig, data)
	return data, nil
}

//
// ========= End def@service.Service =========

// ========= Begin stream.Handler =========
//

// StreamHandler stream.Handler
type StreamHandler interface {
	isStreamHandler()
	Component
}

// RawStreamHandler is store raw bytes of StreamHandler
type RawStreamHandler []byte

func (RawStreamHandler) isStreamHandler() {}
func (RawStreamHandler) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawStreamHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawStreamHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawStreamHandler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End stream.Handler =========

// ========= Begin def@stream.Handler =========
//

const kindDefStreamHandlerConfig = "def@stream.Handler"

// DefStreamHandlerConfig def@stream.Handler
type DefStreamHandlerConfig struct {
	Name string
	Def  StreamHandler
}

func init() {
	_ = defTypes.Register(
		kindDefStreamHandlerConfig,
		func(r *DefStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (DefStreamHandlerConfig) isStreamHandler() {}
func (DefStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t DefStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End def@stream.Handler =========

// ========= Begin dialer.Dialer =========
//

// Dialer dialer.Dialer
type Dialer interface {
	isDialer()
	Component
}

// RawDialer is store raw bytes of Dialer
type RawDialer []byte

func (RawDialer) isDialer()    {}
func (RawDialer) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawDialer) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawDialer) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawDialer: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End dialer.Dialer =========

// ========= Begin def@stream/dialer.Dialer =========
//

const kindDefStreamDialerConfig = "def@stream/dialer.Dialer"

// DefStreamDialerConfig def@stream/dialer.Dialer
type DefStreamDialerConfig struct {
	Name string
	Def  Dialer
}

func init() {
	_ = defTypes.Register(
		kindDefStreamDialerConfig,
		func(r *DefStreamDialerConfig) Dialer {
			return r
		},
	)
}

func (DefStreamDialerConfig) isDialer()    {}
func (DefStreamDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t DefStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefStreamDialerConfig, data)
	return data, nil
}

//
// ========= End def@stream/dialer.Dialer =========

// ========= Begin listener.ListenConfig =========
//

// ListenerListenConfig listener.ListenConfig
type ListenerListenConfig interface {
	isListenerListenConfig()
	Component
}

// RawListenerListenConfig is store raw bytes of ListenerListenConfig
type RawListenerListenConfig []byte

func (RawListenerListenConfig) isListenerListenConfig() {}
func (RawListenerListenConfig) isComponent()            {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawListenerListenConfig) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawListenerListenConfig) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawListenerListenConfig: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End listener.ListenConfig =========

// ========= Begin def@stream/listener.ListenConfig =========
//

const kindDefStreamListenerListenConfigConfig = "def@stream/listener.ListenConfig"

// DefStreamListenerListenConfigConfig def@stream/listener.ListenConfig
type DefStreamListenerListenConfigConfig struct {
	Name string
	Def  ListenerListenConfig
}

func init() {
	_ = defTypes.Register(
		kindDefStreamListenerListenConfigConfig,
		func(r *DefStreamListenerListenConfigConfig) ListenerListenConfig {
			return r
		},
	)
}

func (DefStreamListenerListenConfigConfig) isListenerListenConfig() {}
func (DefStreamListenerListenConfigConfig) isComponent()            {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefStreamListenerListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t DefStreamListenerListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefStreamListenerListenConfigConfig, data)
	return data, nil
}

//
// ========= End def@stream/listener.ListenConfig =========

// ========= Begin def@tls.TLS =========
//

const kindDefTLSConfig = "def@tls.TLS"

// DefTLSConfig def@tls.TLS
type DefTLSConfig struct {
	Name string
	Def  TLS
}

func init() {
	_ = defTypes.Register(
		kindDefTLSConfig,
		func(r *DefTLSConfig) TLS {
			return r
		},
	)
}

func (DefTLSConfig) isTLS()       {}
func (DefTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefTLSConfig) MarshalJSON() ([]byte, error) {
	type t DefTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefTLSConfig, data)
	return data, nil
}

//
// ========= End def@tls.TLS =========

// ========= Begin direct@net/http.Handler =========
//

const kindDirectNetHTTPHandlerConfig = "direct@net/http.Handler"

// DirectNetHTTPHandlerConfig direct@net/http.Handler
type DirectNetHTTPHandlerConfig struct {
	Code int
	Body IoReader
}

func init() {
	_ = defTypes.Register(
		kindDirectNetHTTPHandlerConfig,
		func(r *DirectNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (DirectNetHTTPHandlerConfig) isHTTPHandler() {}
func (DirectNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DirectNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t DirectNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDirectNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End direct@net/http.Handler =========

// ========= Begin expvar@net/http.Handler =========
//

const kindExpvarNetHTTPHandler = "expvar@net/http.Handler"

// ExpvarNetHTTPHandler expvar@net/http.Handler
type ExpvarNetHTTPHandler struct {
}

func init() {
	_ = defTypes.Register(
		kindExpvarNetHTTPHandler,
		func(r *ExpvarNetHTTPHandler) HTTPHandler {
			return r
		},
	)
}

func (ExpvarNetHTTPHandler) isHTTPHandler() {}
func (ExpvarNetHTTPHandler) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ExpvarNetHTTPHandler) MarshalJSON() ([]byte, error) {
	type t ExpvarNetHTTPHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindExpvarNetHTTPHandler, data)
	return data, nil
}

//
// ========= End expvar@net/http.Handler =========

// ========= Begin file@io.Reader =========
//

const kindFileIoReaderConfig = "file@io.Reader"

// FileIoReaderConfig file@io.Reader
type FileIoReaderConfig struct {
	Path string
}

func init() {
	_ = defTypes.Register(
		kindFileIoReaderConfig,
		func(r *FileIoReaderConfig) IoReader {
			return r
		},
	)
}

func (FileIoReaderConfig) isIoReader()  {}
func (FileIoReaderConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m FileIoReaderConfig) MarshalJSON() ([]byte, error) {
	type t FileIoReaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindFileIoReaderConfig, data)
	return data, nil
}

//
// ========= End file@io.Reader =========

// ========= Begin file@io.Writer =========
//

const kindFileIoWriterConfig = "file@io.Writer"

// FileIoWriterConfig file@io.Writer
type FileIoWriterConfig struct {
	Path string
}

func init() {
	_ = defTypes.Register(
		kindFileIoWriterConfig,
		func(r *FileIoWriterConfig) IoWriter {
			return r
		},
	)
}

func (FileIoWriterConfig) isIoWriter()  {}
func (FileIoWriterConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m FileIoWriterConfig) MarshalJSON() ([]byte, error) {
	type t FileIoWriterConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindFileIoWriterConfig, data)
	return data, nil
}

//
// ========= End file@io.Writer =========

// ========= Begin file@net/http.Handler =========
//

const kindFileNetHTTPHandlerConfig = "file@net/http.Handler"

// FileNetHTTPHandlerConfig file@net/http.Handler
type FileNetHTTPHandlerConfig struct {
	Root string
}

func init() {
	_ = defTypes.Register(
		kindFileNetHTTPHandlerConfig,
		func(r *FileNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (FileNetHTTPHandlerConfig) isHTTPHandler() {}
func (FileNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m FileNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t FileNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindFileNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End file@net/http.Handler =========

// ========= Begin forward@net/http.Handler =========
//

const kindForwardNetHTTPHandlerConfig = "forward@net/http.Handler"

// ForwardNetHTTPHandlerConfig forward@net/http.Handler
type ForwardNetHTTPHandlerConfig struct {
	RoundTripper HTTPRoundTripper
	URL          string
}

func init() {
	_ = defTypes.Register(
		kindForwardNetHTTPHandlerConfig,
		func(r *ForwardNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (ForwardNetHTTPHandlerConfig) isHTTPHandler() {}
func (ForwardNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ForwardNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t ForwardNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindForwardNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End forward@net/http.Handler =========

// ========= Begin forward@stream.Handler =========
//

const kindForwardStreamHandlerConfig = "forward@stream.Handler"

// ForwardStreamHandlerConfig forward@stream.Handler
type ForwardStreamHandlerConfig struct {
	Dialer Dialer
}

func init() {
	_ = defTypes.Register(
		kindForwardStreamHandlerConfig,
		func(r *ForwardStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (ForwardStreamHandlerConfig) isStreamHandler() {}
func (ForwardStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ForwardStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t ForwardStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindForwardStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End forward@stream.Handler =========

// ========= Begin from@tls.TLS =========
//

const kindFromTLSConfig = "from@tls.TLS"

// FromTLSConfig from@tls.TLS
type FromTLSConfig struct {
	Domain string
	Cert   IoReader
	Key    IoReader
}

func init() {
	_ = defTypes.Register(
		kindFromTLSConfig,
		func(r *FromTLSConfig) TLS {
			return r
		},
	)
}

func (FromTLSConfig) isTLS()       {}
func (FromTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m FromTLSConfig) MarshalJSON() ([]byte, error) {
	type t FromTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindFromTLSConfig, data)
	return data, nil
}

//
// ========= End from@tls.TLS =========

// ========= Begin gzip@codec.Decoder =========
//

const kindGzipCodecDecoder = "gzip@codec.Decoder"

// GzipCodecDecoder gzip@codec.Decoder
type GzipCodecDecoder struct {
}

func init() {
	_ = defTypes.Register(
		kindGzipCodecDecoder,
		func(r *GzipCodecDecoder) CodecDecoder {
			return r
		},
	)
}

func (GzipCodecDecoder) isCodecDecoder() {}
func (GzipCodecDecoder) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m GzipCodecDecoder) MarshalJSON() ([]byte, error) {
	type t GzipCodecDecoder
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindGzipCodecDecoder, data)
	return data, nil
}

//
// ========= End gzip@codec.Decoder =========

// ========= Begin gzip@codec.Encoder =========
//

const kindGzipCodecEncoder = "gzip@codec.Encoder"

// GzipCodecEncoder gzip@codec.Encoder
type GzipCodecEncoder struct {
}

func init() {
	_ = defTypes.Register(
		kindGzipCodecEncoder,
		func(r *GzipCodecEncoder) CodecEncoder {
			return r
		},
	)
}

func (GzipCodecEncoder) isCodecEncoder() {}
func (GzipCodecEncoder) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m GzipCodecEncoder) MarshalJSON() ([]byte, error) {
	type t GzipCodecEncoder
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindGzipCodecEncoder, data)
	return data, nil
}

//
// ========= End gzip@codec.Encoder =========

// ========= Begin h2c@net/http.Handler =========
//

const kindH2CNetHTTPHandlerConfig = "h2c@net/http.Handler"

// H2CNetHTTPHandlerConfig h2c@net/http.Handler
type H2CNetHTTPHandlerConfig struct {
	Handler HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindH2CNetHTTPHandlerConfig,
		func(r *H2CNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (H2CNetHTTPHandlerConfig) isHTTPHandler() {}
func (H2CNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m H2CNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t H2CNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindH2CNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End h2c@net/http.Handler =========

// ========= Begin hex@codec.Decoder =========
//

const kindHexCodecDecoder = "hex@codec.Decoder"

// HexCodecDecoder hex@codec.Decoder
type HexCodecDecoder struct {
}

func init() {
	_ = defTypes.Register(
		kindHexCodecDecoder,
		func(r *HexCodecDecoder) CodecDecoder {
			return r
		},
	)
}

func (HexCodecDecoder) isCodecDecoder() {}
func (HexCodecDecoder) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HexCodecDecoder) MarshalJSON() ([]byte, error) {
	type t HexCodecDecoder
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindHexCodecDecoder, data)
	return data, nil
}

//
// ========= End hex@codec.Decoder =========

// ========= Begin hex@codec.Encoder =========
//

const kindHexCodecEncoder = "hex@codec.Encoder"

// HexCodecEncoder hex@codec.Encoder
type HexCodecEncoder struct {
}

func init() {
	_ = defTypes.Register(
		kindHexCodecEncoder,
		func(r *HexCodecEncoder) CodecEncoder {
			return r
		},
	)
}

func (HexCodecEncoder) isCodecEncoder() {}
func (HexCodecEncoder) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HexCodecEncoder) MarshalJSON() ([]byte, error) {
	type t HexCodecEncoder
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindHexCodecEncoder, data)
	return data, nil
}

//
// ========= End hex@codec.Encoder =========

// ========= Begin host@net/http.Handler =========
//

const kindHostNetHTTPHandlerConfig = "host@net/http.Handler"

// HostNetHTTPHandlerConfig host@net/http.Handler
type HostNetHTTPHandlerConfig struct {
	Hosts    []HostNetHTTPHandlerRoute
	NotFound HTTPHandler
}

type HostNetHTTPHandlerRoute struct {
	Domain  string
	Handler HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindHostNetHTTPHandlerConfig,
		func(r *HostNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (HostNetHTTPHandlerConfig) isHTTPHandler() {}
func (HostNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HostNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t HostNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindHostNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End host@net/http.Handler =========

// ========= Begin http@stream.Handler =========
//

const kindHTTPStreamHandlerConfig = "http@stream.Handler"

// HTTPStreamHandlerConfig http@stream.Handler
type HTTPStreamHandlerConfig struct {
	Handler HTTPHandler
	TLS     TLS
}

func init() {
	_ = defTypes.Register(
		kindHTTPStreamHandlerConfig,
		func(r *HTTPStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (HTTPStreamHandlerConfig) isStreamHandler() {}
func (HTTPStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t HTTPStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindHTTPStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End http@stream.Handler =========

// ========= Begin inline@io.Reader =========
//

const kindInlineIoReaderConfig = "inline@io.Reader"

// InlineIoReaderConfig inline@io.Reader
type InlineIoReaderConfig struct {
	Data string
}

func init() {
	_ = defTypes.Register(
		kindInlineIoReaderConfig,
		func(r *InlineIoReaderConfig) IoReader {
			return r
		},
	)
}

func (InlineIoReaderConfig) isIoReader()  {}
func (InlineIoReaderConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m InlineIoReaderConfig) MarshalJSON() ([]byte, error) {
	type t InlineIoReaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindInlineIoReaderConfig, data)
	return data, nil
}

//
// ========= End inline@io.Reader =========

// ========= Begin json@codec.Marshaler =========
//

const kindJSONCodecMarshaler = "json@codec.Marshaler"

// JSONCodecMarshaler json@codec.Marshaler
type JSONCodecMarshaler struct {
}

func init() {
	_ = defTypes.Register(
		kindJSONCodecMarshaler,
		func(r *JSONCodecMarshaler) CodecMarshaler {
			return r
		},
	)
}

func (JSONCodecMarshaler) isCodecMarshaler() {}
func (JSONCodecMarshaler) isComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m JSONCodecMarshaler) MarshalJSON() ([]byte, error) {
	type t JSONCodecMarshaler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindJSONCodecMarshaler, data)
	return data, nil
}

//
// ========= End json@codec.Marshaler =========

// ========= Begin json@codec.Unmarshaler =========
//

const kindJSONCodecUnmarshaler = "json@codec.Unmarshaler"

// JSONCodecUnmarshaler json@codec.Unmarshaler
type JSONCodecUnmarshaler struct {
}

func init() {
	_ = defTypes.Register(
		kindJSONCodecUnmarshaler,
		func(r *JSONCodecUnmarshaler) CodecUnmarshaler {
			return r
		},
	)
}

func (JSONCodecUnmarshaler) isCodecUnmarshaler() {}
func (JSONCodecUnmarshaler) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m JSONCodecUnmarshaler) MarshalJSON() ([]byte, error) {
	type t JSONCodecUnmarshaler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindJSONCodecUnmarshaler, data)
	return data, nil
}

//
// ========= End json@codec.Unmarshaler =========

// ========= Begin load@codec.Decoder =========
//

const kindLoadCodecDecoderConfig = "load@codec.Decoder"

// LoadCodecDecoderConfig load@codec.Decoder
type LoadCodecDecoderConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadCodecDecoderConfig,
		func(r *LoadCodecDecoderConfig) CodecDecoder {
			return r
		},
	)
}

func (LoadCodecDecoderConfig) isCodecDecoder() {}
func (LoadCodecDecoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadCodecDecoderConfig) MarshalJSON() ([]byte, error) {
	type t LoadCodecDecoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadCodecDecoderConfig, data)
	return data, nil
}

//
// ========= End load@codec.Decoder =========

// ========= Begin load@codec.Encoder =========
//

const kindLoadCodecEncoderConfig = "load@codec.Encoder"

// LoadCodecEncoderConfig load@codec.Encoder
type LoadCodecEncoderConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadCodecEncoderConfig,
		func(r *LoadCodecEncoderConfig) CodecEncoder {
			return r
		},
	)
}

func (LoadCodecEncoderConfig) isCodecEncoder() {}
func (LoadCodecEncoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadCodecEncoderConfig) MarshalJSON() ([]byte, error) {
	type t LoadCodecEncoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadCodecEncoderConfig, data)
	return data, nil
}

//
// ========= End load@codec.Encoder =========

// ========= Begin load@codec.Marshaler =========
//

const kindLoadCodecMarshalerConfig = "load@codec.Marshaler"

// LoadCodecMarshalerConfig load@codec.Marshaler
type LoadCodecMarshalerConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadCodecMarshalerConfig,
		func(r *LoadCodecMarshalerConfig) CodecMarshaler {
			return r
		},
	)
}

func (LoadCodecMarshalerConfig) isCodecMarshaler() {}
func (LoadCodecMarshalerConfig) isComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadCodecMarshalerConfig) MarshalJSON() ([]byte, error) {
	type t LoadCodecMarshalerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadCodecMarshalerConfig, data)
	return data, nil
}

//
// ========= End load@codec.Marshaler =========

// ========= Begin load@codec.Unmarshaler =========
//

const kindLoadCodecUnmarshalerConfig = "load@codec.Unmarshaler"

// LoadCodecUnmarshalerConfig load@codec.Unmarshaler
type LoadCodecUnmarshalerConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadCodecUnmarshalerConfig,
		func(r *LoadCodecUnmarshalerConfig) CodecUnmarshaler {
			return r
		},
	)
}

func (LoadCodecUnmarshalerConfig) isCodecUnmarshaler() {}
func (LoadCodecUnmarshalerConfig) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadCodecUnmarshalerConfig) MarshalJSON() ([]byte, error) {
	type t LoadCodecUnmarshalerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadCodecUnmarshalerConfig, data)
	return data, nil
}

//
// ========= End load@codec.Unmarshaler =========

// ========= Begin load@io.Reader =========
//

const kindLoadIoReaderConfig = "load@io.Reader"

// LoadIoReaderConfig load@io.Reader
type LoadIoReaderConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadIoReaderConfig,
		func(r *LoadIoReaderConfig) IoReader {
			return r
		},
	)
}

func (LoadIoReaderConfig) isIoReader()  {}
func (LoadIoReaderConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadIoReaderConfig) MarshalJSON() ([]byte, error) {
	type t LoadIoReaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadIoReaderConfig, data)
	return data, nil
}

//
// ========= End load@io.Reader =========

// ========= Begin load@io.Writer =========
//

const kindLoadIoWriterConfig = "load@io.Writer"

// LoadIoWriterConfig load@io.Writer
type LoadIoWriterConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadIoWriterConfig,
		func(r *LoadIoWriterConfig) IoWriter {
			return r
		},
	)
}

func (LoadIoWriterConfig) isIoWriter()  {}
func (LoadIoWriterConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadIoWriterConfig) MarshalJSON() ([]byte, error) {
	type t LoadIoWriterConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadIoWriterConfig, data)
	return data, nil
}

//
// ========= End load@io.Writer =========

// ========= Begin load@net/http.Handler =========
//

const kindLoadNetHTTPHandlerConfig = "load@net/http.Handler"

// LoadNetHTTPHandlerConfig load@net/http.Handler
type LoadNetHTTPHandlerConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadNetHTTPHandlerConfig,
		func(r *LoadNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (LoadNetHTTPHandlerConfig) isHTTPHandler() {}
func (LoadNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t LoadNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End load@net/http.Handler =========

// ========= Begin load@net/http.RoundTripper =========
//

const kindLoadNetHTTPRoundTripperConfig = "load@net/http.RoundTripper"

// LoadNetHTTPRoundTripperConfig load@net/http.RoundTripper
type LoadNetHTTPRoundTripperConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadNetHTTPRoundTripperConfig,
		func(r *LoadNetHTTPRoundTripperConfig) HTTPRoundTripper {
			return r
		},
	)
}

func (LoadNetHTTPRoundTripperConfig) isHTTPRoundTripper() {}
func (LoadNetHTTPRoundTripperConfig) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadNetHTTPRoundTripperConfig) MarshalJSON() ([]byte, error) {
	type t LoadNetHTTPRoundTripperConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadNetHTTPRoundTripperConfig, data)
	return data, nil
}

//
// ========= End load@net/http.RoundTripper =========

// ========= Begin load@once.Once =========
//

const kindLoadOnceConfig = "load@once.Once"

// LoadOnceConfig load@once.Once
type LoadOnceConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadOnceConfig,
		func(r *LoadOnceConfig) Once {
			return r
		},
	)
}

func (LoadOnceConfig) isOnce()      {}
func (LoadOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadOnceConfig) MarshalJSON() ([]byte, error) {
	type t LoadOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadOnceConfig, data)
	return data, nil
}

//
// ========= End load@once.Once =========

// ========= Begin load@service.Service =========
//

const kindLoadServiceConfig = "load@service.Service"

// LoadServiceConfig load@service.Service
type LoadServiceConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadServiceConfig,
		func(r *LoadServiceConfig) Service {
			return r
		},
	)
}

func (LoadServiceConfig) isService()   {}
func (LoadServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadServiceConfig) MarshalJSON() ([]byte, error) {
	type t LoadServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadServiceConfig, data)
	return data, nil
}

//
// ========= End load@service.Service =========

// ========= Begin load@stream.Handler =========
//

const kindLoadStreamHandlerConfig = "load@stream.Handler"

// LoadStreamHandlerConfig load@stream.Handler
type LoadStreamHandlerConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadStreamHandlerConfig,
		func(r *LoadStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (LoadStreamHandlerConfig) isStreamHandler() {}
func (LoadStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t LoadStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End load@stream.Handler =========

// ========= Begin load@stream/dialer.Dialer =========
//

const kindLoadStreamDialerConfig = "load@stream/dialer.Dialer"

// LoadStreamDialerConfig load@stream/dialer.Dialer
type LoadStreamDialerConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadStreamDialerConfig,
		func(r *LoadStreamDialerConfig) Dialer {
			return r
		},
	)
}

func (LoadStreamDialerConfig) isDialer()    {}
func (LoadStreamDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t LoadStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadStreamDialerConfig, data)
	return data, nil
}

//
// ========= End load@stream/dialer.Dialer =========

// ========= Begin load@stream/listener.ListenConfig =========
//

const kindLoadStreamListenerListenConfigConfig = "load@stream/listener.ListenConfig"

// LoadStreamListenerListenConfigConfig load@stream/listener.ListenConfig
type LoadStreamListenerListenConfigConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadStreamListenerListenConfigConfig,
		func(r *LoadStreamListenerListenConfigConfig) ListenerListenConfig {
			return r
		},
	)
}

func (LoadStreamListenerListenConfigConfig) isListenerListenConfig() {}
func (LoadStreamListenerListenConfigConfig) isComponent()            {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadStreamListenerListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t LoadStreamListenerListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadStreamListenerListenConfigConfig, data)
	return data, nil
}

//
// ========= End load@stream/listener.ListenConfig =========

// ========= Begin load@tls.TLS =========
//

const kindLoadTLSConfig = "load@tls.TLS"

// LoadTLSConfig load@tls.TLS
type LoadTLSConfig struct {
	Load IoReader
}

func init() {
	_ = defTypes.Register(
		kindLoadTLSConfig,
		func(r *LoadTLSConfig) TLS {
			return r
		},
	)
}

func (LoadTLSConfig) isTLS()       {}
func (LoadTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadTLSConfig) MarshalJSON() ([]byte, error) {
	type t LoadTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadTLSConfig, data)
	return data, nil
}

//
// ========= End load@tls.TLS =========

// ========= Begin log@net/http.Handler =========
//

const kindLogNetHTTPHandlerConfig = "log@net/http.Handler"

// LogNetHTTPHandlerConfig log@net/http.Handler
type LogNetHTTPHandlerConfig struct {
	Output  IoWriter
	Handler HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindLogNetHTTPHandlerConfig,
		func(r *LogNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (LogNetHTTPHandlerConfig) isHTTPHandler() {}
func (LogNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LogNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t LogNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLogNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End log@net/http.Handler =========

// ========= Begin message@once.Once =========
//

const kindMessageOnceConfig = "message@once.Once"

// MessageOnceConfig message@once.Once
type MessageOnceConfig struct {
	Message string
}

func init() {
	_ = defTypes.Register(
		kindMessageOnceConfig,
		func(r *MessageOnceConfig) Once {
			return r
		},
	)
}

func (MessageOnceConfig) isOnce()      {}
func (MessageOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MessageOnceConfig) MarshalJSON() ([]byte, error) {
	type t MessageOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMessageOnceConfig, data)
	return data, nil
}

//
// ========= End message@once.Once =========

// ========= Begin multi@net/http.Handler =========
//

const kindMultiNetHTTPHandlerConfig = "multi@net/http.Handler"

// MultiNetHTTPHandlerConfig multi@net/http.Handler
type MultiNetHTTPHandlerConfig struct {
	Multi []HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindMultiNetHTTPHandlerConfig,
		func(r *MultiNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (MultiNetHTTPHandlerConfig) isHTTPHandler() {}
func (MultiNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MultiNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t MultiNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMultiNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End multi@net/http.Handler =========

// ========= Begin multi@once.Once =========
//

const kindMultiOnceConfig = "multi@once.Once"

// MultiOnceConfig multi@once.Once
type MultiOnceConfig struct {
	Multi []Once
}

func init() {
	_ = defTypes.Register(
		kindMultiOnceConfig,
		func(r *MultiOnceConfig) Once {
			return r
		},
	)
}

func (MultiOnceConfig) isOnce()      {}
func (MultiOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MultiOnceConfig) MarshalJSON() ([]byte, error) {
	type t MultiOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMultiOnceConfig, data)
	return data, nil
}

//
// ========= End multi@once.Once =========

// ========= Begin multi@service.Service =========
//

const kindMultiServiceConfig = "multi@service.Service"

// MultiServiceConfig multi@service.Service
type MultiServiceConfig struct {
	Multi []Service
}

func init() {
	_ = defTypes.Register(
		kindMultiServiceConfig,
		func(r *MultiServiceConfig) Service {
			return r
		},
	)
}

func (MultiServiceConfig) isService()   {}
func (MultiServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MultiServiceConfig) MarshalJSON() ([]byte, error) {
	type t MultiServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMultiServiceConfig, data)
	return data, nil
}

//
// ========= End multi@service.Service =========

// ========= Begin multi@stream.Handler =========
//

const kindMultiStreamHandlerConfig = "multi@stream.Handler"

// MultiStreamHandlerConfig multi@stream.Handler
type MultiStreamHandlerConfig struct {
	Multi []StreamHandler
}

func init() {
	_ = defTypes.Register(
		kindMultiStreamHandlerConfig,
		func(r *MultiStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (MultiStreamHandlerConfig) isStreamHandler() {}
func (MultiStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MultiStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t MultiStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMultiStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End multi@stream.Handler =========

// ========= Begin mux@net/http.Handler =========
//

const kindMuxNetHTTPHandlerConfig = "mux@net/http.Handler"

// MuxNetHTTPHandlerConfig mux@net/http.Handler
type MuxNetHTTPHandlerConfig struct {
	Routes   []MuxNetHTTPHandlerRoute
	NotFound HTTPHandler
}

type MuxNetHTTPHandlerRoute struct {
	Prefix  string
	Path    string
	Regexp  string
	Handler HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindMuxNetHTTPHandlerConfig,
		func(r *MuxNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (MuxNetHTTPHandlerConfig) isHTTPHandler() {}
func (MuxNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MuxNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t MuxNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMuxNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End mux@net/http.Handler =========

// ========= Begin mux@stream.Handler =========
//

const kindMuxStreamHandlerConfig = "mux@stream.Handler"

// MuxStreamHandlerConfig mux@stream.Handler
type MuxStreamHandlerConfig struct {
	Routes   []MuxStreamHandlerRoute
	NotFound StreamHandler
}

type MuxStreamHandlerRoute struct {
	Pattern string
	Regexp  string
	Prefix  string
	Handler StreamHandler
}

func init() {
	_ = defTypes.Register(
		kindMuxStreamHandlerConfig,
		func(r *MuxStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (MuxStreamHandlerConfig) isStreamHandler() {}
func (MuxStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MuxStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t MuxStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMuxStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End mux@stream.Handler =========

// ========= Begin network@stream/dialer.Dialer =========
//

const kindNetworkStreamDialerConfig = "network@stream/dialer.Dialer"

// NetworkStreamDialerConfig network@stream/dialer.Dialer
type NetworkStreamDialerConfig struct {
	Network NetworkStreamDialerNetworkEnum
	Address string
}

type NetworkStreamDialerNetworkEnum string

const (
	NetworkStreamDialerNetworkEnumEnumUnix NetworkStreamDialerNetworkEnum = "unix"
	NetworkStreamDialerNetworkEnumEnumTCP6 NetworkStreamDialerNetworkEnum = "tcp6"
	NetworkStreamDialerNetworkEnumEnumTCP4 NetworkStreamDialerNetworkEnum = "tcp4"
	NetworkStreamDialerNetworkEnumEnumTCP  NetworkStreamDialerNetworkEnum = "tcp"
)

func init() {
	_ = defTypes.Register(
		kindNetworkStreamDialerConfig,
		func(r *NetworkStreamDialerConfig) Dialer {
			return r
		},
	)
}

func (NetworkStreamDialerConfig) isDialer()    {}
func (NetworkStreamDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NetworkStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t NetworkStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindNetworkStreamDialerConfig, data)
	return data, nil
}

//
// ========= End network@stream/dialer.Dialer =========

// ========= Begin network@stream/listener.ListenConfig =========
//

const kindNetworkStreamListenerListenConfigConfig = "network@stream/listener.ListenConfig"

// NetworkStreamListenerListenConfigConfig network@stream/listener.ListenConfig
type NetworkStreamListenerListenConfigConfig struct {
	Network NetworkStreamListenerListenConfigNetworkEnum
	Address string
}

type NetworkStreamListenerListenConfigNetworkEnum string

const (
	NetworkStreamListenerListenConfigNetworkEnumEnumUnix NetworkStreamListenerListenConfigNetworkEnum = "unix"
	NetworkStreamListenerListenConfigNetworkEnumEnumTCP6 NetworkStreamListenerListenConfigNetworkEnum = "tcp6"
	NetworkStreamListenerListenConfigNetworkEnumEnumTCP4 NetworkStreamListenerListenConfigNetworkEnum = "tcp4"
	NetworkStreamListenerListenConfigNetworkEnumEnumTCP  NetworkStreamListenerListenConfigNetworkEnum = "tcp"
)

func init() {
	_ = defTypes.Register(
		kindNetworkStreamListenerListenConfigConfig,
		func(r *NetworkStreamListenerListenConfigConfig) ListenerListenConfig {
			return r
		},
	)
}

func (NetworkStreamListenerListenConfigConfig) isListenerListenConfig() {}
func (NetworkStreamListenerListenConfigConfig) isComponent()            {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NetworkStreamListenerListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t NetworkStreamListenerListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindNetworkStreamListenerListenConfigConfig, data)
	return data, nil
}

//
// ========= End network@stream/listener.ListenConfig =========

// ========= Begin none@once.Once =========
//

const kindNoneOnceConfig = "none@once.Once"

// NoneOnceConfig none@once.Once
type NoneOnceConfig struct {
	Any Component
}

func init() {
	_ = defTypes.Register(
		kindNoneOnceConfig,
		func(r *NoneOnceConfig) Once {
			return r
		},
	)
}

func (NoneOnceConfig) isOnce()      {}
func (NoneOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneOnceConfig) MarshalJSON() ([]byte, error) {
	type t NoneOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindNoneOnceConfig, data)
	return data, nil
}

//
// ========= End none@once.Once =========

// ========= Begin poller@net/http.Handler =========
//

const kindPollerNetHTTPHandlerConfig = "poller@net/http.Handler"

// PollerNetHTTPHandlerConfig poller@net/http.Handler
type PollerNetHTTPHandlerConfig struct {
	Poller   PollerNetHTTPHandlerPollerEnum
	Handlers []HTTPHandler
}

type PollerNetHTTPHandlerPollerEnum string

const (
	PollerNetHTTPHandlerPollerEnumEnumRoundRobin PollerNetHTTPHandlerPollerEnum = "round_robin"
	PollerNetHTTPHandlerPollerEnumEnumRandom     PollerNetHTTPHandlerPollerEnum = "random"
)

func init() {
	_ = defTypes.Register(
		kindPollerNetHTTPHandlerConfig,
		func(r *PollerNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (PollerNetHTTPHandlerConfig) isHTTPHandler() {}
func (PollerNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m PollerNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t PollerNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindPollerNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End poller@net/http.Handler =========

// ========= Begin poller@stream.Handler =========
//

const kindPollerStreamHandlerConfig = "poller@stream.Handler"

// PollerStreamHandlerConfig poller@stream.Handler
type PollerStreamHandlerConfig struct {
	Poller   PollerStreamHandlerPollerEnum
	Handlers []StreamHandler
}

type PollerStreamHandlerPollerEnum string

const (
	PollerStreamHandlerPollerEnumEnumRoundRobin PollerStreamHandlerPollerEnum = "round_robin"
	PollerStreamHandlerPollerEnumEnumRandom     PollerStreamHandlerPollerEnum = "random"
)

func init() {
	_ = defTypes.Register(
		kindPollerStreamHandlerConfig,
		func(r *PollerStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (PollerStreamHandlerConfig) isStreamHandler() {}
func (PollerStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m PollerStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t PollerStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindPollerStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End poller@stream.Handler =========

// ========= Begin poller@stream/dialer.Dialer =========
//

const kindPollerStreamDialerConfig = "poller@stream/dialer.Dialer"

// PollerStreamDialerConfig poller@stream/dialer.Dialer
type PollerStreamDialerConfig struct {
	Poller  PollerStreamDialerPollerEnum
	Dialers []Dialer
}

type PollerStreamDialerPollerEnum string

const (
	PollerStreamDialerPollerEnumEnumRoundRobin PollerStreamDialerPollerEnum = "round_robin"
	PollerStreamDialerPollerEnumEnumRandom     PollerStreamDialerPollerEnum = "random"
)

func init() {
	_ = defTypes.Register(
		kindPollerStreamDialerConfig,
		func(r *PollerStreamDialerConfig) Dialer {
			return r
		},
	)
}

func (PollerStreamDialerConfig) isDialer()    {}
func (PollerStreamDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m PollerStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t PollerStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindPollerStreamDialerConfig, data)
	return data, nil
}

//
// ========= End poller@stream/dialer.Dialer =========

// ========= Begin pprof@net/http.Handler =========
//

const kindPprofNetHTTPHandler = "pprof@net/http.Handler"

// PprofNetHTTPHandler pprof@net/http.Handler
type PprofNetHTTPHandler struct {
}

func init() {
	_ = defTypes.Register(
		kindPprofNetHTTPHandler,
		func(r *PprofNetHTTPHandler) HTTPHandler {
			return r
		},
	)
}

func (PprofNetHTTPHandler) isHTTPHandler() {}
func (PprofNetHTTPHandler) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m PprofNetHTTPHandler) MarshalJSON() ([]byte, error) {
	type t PprofNetHTTPHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindPprofNetHTTPHandler, data)
	return data, nil
}

//
// ========= End pprof@net/http.Handler =========

// ========= Begin redirect@net/http.Handler =========
//

const kindRedirectNetHTTPHandlerConfig = "redirect@net/http.Handler"

// RedirectNetHTTPHandlerConfig redirect@net/http.Handler
type RedirectNetHTTPHandlerConfig struct {
	Code     int
	Location string
}

func init() {
	_ = defTypes.Register(
		kindRedirectNetHTTPHandlerConfig,
		func(r *RedirectNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (RedirectNetHTTPHandlerConfig) isHTTPHandler() {}
func (RedirectNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RedirectNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t RedirectNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRedirectNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End redirect@net/http.Handler =========

// ========= Begin ref@codec.Decoder =========
//

const kindRefCodecDecoderConfig = "ref@codec.Decoder"

// RefCodecDecoderConfig ref@codec.Decoder
type RefCodecDecoderConfig struct {
	Name string
	Def  CodecDecoder
}

func init() {
	_ = defTypes.Register(
		kindRefCodecDecoderConfig,
		func(r *RefCodecDecoderConfig) CodecDecoder {
			return r
		},
	)
}

func (RefCodecDecoderConfig) isCodecDecoder() {}
func (RefCodecDecoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefCodecDecoderConfig) MarshalJSON() ([]byte, error) {
	type t RefCodecDecoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefCodecDecoderConfig, data)
	return data, nil
}

//
// ========= End ref@codec.Decoder =========

// ========= Begin ref@codec.Encoder =========
//

const kindRefCodecEncoderConfig = "ref@codec.Encoder"

// RefCodecEncoderConfig ref@codec.Encoder
type RefCodecEncoderConfig struct {
	Name string
	Def  CodecEncoder
}

func init() {
	_ = defTypes.Register(
		kindRefCodecEncoderConfig,
		func(r *RefCodecEncoderConfig) CodecEncoder {
			return r
		},
	)
}

func (RefCodecEncoderConfig) isCodecEncoder() {}
func (RefCodecEncoderConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefCodecEncoderConfig) MarshalJSON() ([]byte, error) {
	type t RefCodecEncoderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefCodecEncoderConfig, data)
	return data, nil
}

//
// ========= End ref@codec.Encoder =========

// ========= Begin ref@codec.Marshaler =========
//

const kindRefCodecMarshalerConfig = "ref@codec.Marshaler"

// RefCodecMarshalerConfig ref@codec.Marshaler
type RefCodecMarshalerConfig struct {
	Name string
	Def  CodecMarshaler
}

func init() {
	_ = defTypes.Register(
		kindRefCodecMarshalerConfig,
		func(r *RefCodecMarshalerConfig) CodecMarshaler {
			return r
		},
	)
}

func (RefCodecMarshalerConfig) isCodecMarshaler() {}
func (RefCodecMarshalerConfig) isComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefCodecMarshalerConfig) MarshalJSON() ([]byte, error) {
	type t RefCodecMarshalerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefCodecMarshalerConfig, data)
	return data, nil
}

//
// ========= End ref@codec.Marshaler =========

// ========= Begin ref@codec.Unmarshaler =========
//

const kindRefCodecUnmarshalerConfig = "ref@codec.Unmarshaler"

// RefCodecUnmarshalerConfig ref@codec.Unmarshaler
type RefCodecUnmarshalerConfig struct {
	Name string
	Def  CodecUnmarshaler
}

func init() {
	_ = defTypes.Register(
		kindRefCodecUnmarshalerConfig,
		func(r *RefCodecUnmarshalerConfig) CodecUnmarshaler {
			return r
		},
	)
}

func (RefCodecUnmarshalerConfig) isCodecUnmarshaler() {}
func (RefCodecUnmarshalerConfig) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefCodecUnmarshalerConfig) MarshalJSON() ([]byte, error) {
	type t RefCodecUnmarshalerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefCodecUnmarshalerConfig, data)
	return data, nil
}

//
// ========= End ref@codec.Unmarshaler =========

// ========= Begin ref@io.Reader =========
//

const kindRefIoReaderConfig = "ref@io.Reader"

// RefIoReaderConfig ref@io.Reader
type RefIoReaderConfig struct {
	Name string
	Def  IoReader
}

func init() {
	_ = defTypes.Register(
		kindRefIoReaderConfig,
		func(r *RefIoReaderConfig) IoReader {
			return r
		},
	)
}

func (RefIoReaderConfig) isIoReader()  {}
func (RefIoReaderConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefIoReaderConfig) MarshalJSON() ([]byte, error) {
	type t RefIoReaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefIoReaderConfig, data)
	return data, nil
}

//
// ========= End ref@io.Reader =========

// ========= Begin ref@io.Writer =========
//

const kindRefIoWriterConfig = "ref@io.Writer"

// RefIoWriterConfig ref@io.Writer
type RefIoWriterConfig struct {
	Name string
	Def  IoWriter
}

func init() {
	_ = defTypes.Register(
		kindRefIoWriterConfig,
		func(r *RefIoWriterConfig) IoWriter {
			return r
		},
	)
}

func (RefIoWriterConfig) isIoWriter()  {}
func (RefIoWriterConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefIoWriterConfig) MarshalJSON() ([]byte, error) {
	type t RefIoWriterConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefIoWriterConfig, data)
	return data, nil
}

//
// ========= End ref@io.Writer =========

// ========= Begin ref@net/http.Handler =========
//

const kindRefNetHTTPHandlerConfig = "ref@net/http.Handler"

// RefNetHTTPHandlerConfig ref@net/http.Handler
type RefNetHTTPHandlerConfig struct {
	Name string
	Def  HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindRefNetHTTPHandlerConfig,
		func(r *RefNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (RefNetHTTPHandlerConfig) isHTTPHandler() {}
func (RefNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t RefNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End ref@net/http.Handler =========

// ========= Begin ref@net/http.RoundTripper =========
//

const kindRefNetHTTPRoundTripperConfig = "ref@net/http.RoundTripper"

// RefNetHTTPRoundTripperConfig ref@net/http.RoundTripper
type RefNetHTTPRoundTripperConfig struct {
	Name string
	Def  HTTPRoundTripper
}

func init() {
	_ = defTypes.Register(
		kindRefNetHTTPRoundTripperConfig,
		func(r *RefNetHTTPRoundTripperConfig) HTTPRoundTripper {
			return r
		},
	)
}

func (RefNetHTTPRoundTripperConfig) isHTTPRoundTripper() {}
func (RefNetHTTPRoundTripperConfig) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefNetHTTPRoundTripperConfig) MarshalJSON() ([]byte, error) {
	type t RefNetHTTPRoundTripperConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefNetHTTPRoundTripperConfig, data)
	return data, nil
}

//
// ========= End ref@net/http.RoundTripper =========

// ========= Begin ref@once.Once =========
//

const kindRefOnceConfig = "ref@once.Once"

// RefOnceConfig ref@once.Once
type RefOnceConfig struct {
	Name string
	Def  Once
}

func init() {
	_ = defTypes.Register(
		kindRefOnceConfig,
		func(r *RefOnceConfig) Once {
			return r
		},
	)
}

func (RefOnceConfig) isOnce()      {}
func (RefOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefOnceConfig) MarshalJSON() ([]byte, error) {
	type t RefOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefOnceConfig, data)
	return data, nil
}

//
// ========= End ref@once.Once =========

// ========= Begin ref@service.Service =========
//

const kindRefServiceConfig = "ref@service.Service"

// RefServiceConfig ref@service.Service
type RefServiceConfig struct {
	Name string
	Def  Service
}

func init() {
	_ = defTypes.Register(
		kindRefServiceConfig,
		func(r *RefServiceConfig) Service {
			return r
		},
	)
}

func (RefServiceConfig) isService()   {}
func (RefServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefServiceConfig) MarshalJSON() ([]byte, error) {
	type t RefServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefServiceConfig, data)
	return data, nil
}

//
// ========= End ref@service.Service =========

// ========= Begin ref@stream.Handler =========
//

const kindRefStreamHandlerConfig = "ref@stream.Handler"

// RefStreamHandlerConfig ref@stream.Handler
type RefStreamHandlerConfig struct {
	Name string
	Def  StreamHandler
}

func init() {
	_ = defTypes.Register(
		kindRefStreamHandlerConfig,
		func(r *RefStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (RefStreamHandlerConfig) isStreamHandler() {}
func (RefStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t RefStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End ref@stream.Handler =========

// ========= Begin ref@stream/dialer.Dialer =========
//

const kindRefStreamDialerConfig = "ref@stream/dialer.Dialer"

// RefStreamDialerConfig ref@stream/dialer.Dialer
type RefStreamDialerConfig struct {
	Name string
	Def  Dialer
}

func init() {
	_ = defTypes.Register(
		kindRefStreamDialerConfig,
		func(r *RefStreamDialerConfig) Dialer {
			return r
		},
	)
}

func (RefStreamDialerConfig) isDialer()    {}
func (RefStreamDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t RefStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefStreamDialerConfig, data)
	return data, nil
}

//
// ========= End ref@stream/dialer.Dialer =========

// ========= Begin ref@stream/listener.ListenConfig =========
//

const kindRefStreamListenerListenConfigConfig = "ref@stream/listener.ListenConfig"

// RefStreamListenerListenConfigConfig ref@stream/listener.ListenConfig
type RefStreamListenerListenConfigConfig struct {
	Name string
	Def  ListenerListenConfig
}

func init() {
	_ = defTypes.Register(
		kindRefStreamListenerListenConfigConfig,
		func(r *RefStreamListenerListenConfigConfig) ListenerListenConfig {
			return r
		},
	)
}

func (RefStreamListenerListenConfigConfig) isListenerListenConfig() {}
func (RefStreamListenerListenConfigConfig) isComponent()            {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefStreamListenerListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t RefStreamListenerListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefStreamListenerListenConfigConfig, data)
	return data, nil
}

//
// ========= End ref@stream/listener.ListenConfig =========

// ========= Begin ref@tls.TLS =========
//

const kindRefTLSConfig = "ref@tls.TLS"

// RefTLSConfig ref@tls.TLS
type RefTLSConfig struct {
	Name string
	Def  TLS
}

func init() {
	_ = defTypes.Register(
		kindRefTLSConfig,
		func(r *RefTLSConfig) TLS {
			return r
		},
	)
}

func (RefTLSConfig) isTLS()       {}
func (RefTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefTLSConfig) MarshalJSON() ([]byte, error) {
	type t RefTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefTLSConfig, data)
	return data, nil
}

//
// ========= End ref@tls.TLS =========

// ========= Begin remove_request_header@net/http.Handler =========
//

const kindRemoveRequestHeaderNetHTTPHandlerConfig = "remove_request_header@net/http.Handler"

// RemoveRequestHeaderNetHTTPHandlerConfig remove_request_header@net/http.Handler
type RemoveRequestHeaderNetHTTPHandlerConfig struct {
	Key string
}

func init() {
	_ = defTypes.Register(
		kindRemoveRequestHeaderNetHTTPHandlerConfig,
		func(r *RemoveRequestHeaderNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (RemoveRequestHeaderNetHTTPHandlerConfig) isHTTPHandler() {}
func (RemoveRequestHeaderNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RemoveRequestHeaderNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t RemoveRequestHeaderNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRemoveRequestHeaderNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End remove_request_header@net/http.Handler =========

// ========= Begin remove_response_header@net/http.Handler =========
//

const kindRemoveResponseHeaderNetHTTPHandlerConfig = "remove_response_header@net/http.Handler"

// RemoveResponseHeaderNetHTTPHandlerConfig remove_response_header@net/http.Handler
type RemoveResponseHeaderNetHTTPHandlerConfig struct {
	Key string
}

func init() {
	_ = defTypes.Register(
		kindRemoveResponseHeaderNetHTTPHandlerConfig,
		func(r *RemoveResponseHeaderNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (RemoveResponseHeaderNetHTTPHandlerConfig) isHTTPHandler() {}
func (RemoveResponseHeaderNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RemoveResponseHeaderNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t RemoveResponseHeaderNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRemoveResponseHeaderNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End remove_response_header@net/http.Handler =========

// ========= Begin sample@once.Once =========
//

const kindSampleOnceConfig = "sample@once.Once"

// SampleOnceConfig sample@once.Once
type SampleOnceConfig struct {
	Components []Component
	Pipe       Service
	Init       []Once
}

func init() {
	_ = defTypes.Register(
		kindSampleOnceConfig,
		func(r *SampleOnceConfig) Once {
			return r
		},
	)
}

func (SampleOnceConfig) isOnce()      {}
func (SampleOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m SampleOnceConfig) MarshalJSON() ([]byte, error) {
	type t SampleOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindSampleOnceConfig, data)
	return data, nil
}

//
// ========= End sample@once.Once =========

// ========= Begin self_signed@tls.TLS =========
//

const kindSelfSignedTLS = "self_signed@tls.TLS"

// SelfSignedTLS self_signed@tls.TLS
type SelfSignedTLS struct {
}

func init() {
	_ = defTypes.Register(
		kindSelfSignedTLS,
		func(r *SelfSignedTLS) TLS {
			return r
		},
	)
}

func (SelfSignedTLS) isTLS()       {}
func (SelfSignedTLS) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m SelfSignedTLS) MarshalJSON() ([]byte, error) {
	type t SelfSignedTLS
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindSelfSignedTLS, data)
	return data, nil
}

//
// ========= End self_signed@tls.TLS =========

// ========= Begin service@once.Once =========
//

const kindServiceOnceConfig = "service@once.Once"

// ServiceOnceConfig service@once.Once
type ServiceOnceConfig struct {
	Service Service
}

func init() {
	_ = defTypes.Register(
		kindServiceOnceConfig,
		func(r *ServiceOnceConfig) Once {
			return r
		},
	)
}

func (ServiceOnceConfig) isOnce()      {}
func (ServiceOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceOnceConfig) MarshalJSON() ([]byte, error) {
	type t ServiceOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindServiceOnceConfig, data)
	return data, nil
}

//
// ========= End service@once.Once =========

// ========= Begin stream@service.Service =========
//

const kindStreamServiceConfig = "stream@service.Service"

// StreamServiceConfig stream@service.Service
type StreamServiceConfig struct {
	Listener ListenerListenConfig
	Handler  StreamHandler
}

func init() {
	_ = defTypes.Register(
		kindStreamServiceConfig,
		func(r *StreamServiceConfig) Service {
			return r
		},
	)
}

func (StreamServiceConfig) isService()   {}
func (StreamServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamServiceConfig) MarshalJSON() ([]byte, error) {
	type t StreamServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindStreamServiceConfig, data)
	return data, nil
}

//
// ========= End stream@service.Service =========

// ========= Begin tls@stream/dialer.Dialer =========
//

const kindTLSStreamDialerConfig = "tls@stream/dialer.Dialer"

// TLSStreamDialerConfig tls@stream/dialer.Dialer
type TLSStreamDialerConfig struct {
	Dialer Dialer
	TLS    TLS
}

func init() {
	_ = defTypes.Register(
		kindTLSStreamDialerConfig,
		func(r *TLSStreamDialerConfig) Dialer {
			return r
		},
	)
}

func (TLSStreamDialerConfig) isDialer()    {}
func (TLSStreamDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t TLSStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindTLSStreamDialerConfig, data)
	return data, nil
}

//
// ========= End tls@stream/dialer.Dialer =========

// ========= Begin tls@stream/listener.ListenConfig =========
//

const kindTLSStreamListenerListenConfigConfig = "tls@stream/listener.ListenConfig"

// TLSStreamListenerListenConfigConfig tls@stream/listener.ListenConfig
type TLSStreamListenerListenConfigConfig struct {
	ListenConfig ListenerListenConfig
	TLS          TLS
}

func init() {
	_ = defTypes.Register(
		kindTLSStreamListenerListenConfigConfig,
		func(r *TLSStreamListenerListenConfigConfig) ListenerListenConfig {
			return r
		},
	)
}

func (TLSStreamListenerListenConfigConfig) isListenerListenConfig() {}
func (TLSStreamListenerListenConfigConfig) isComponent()            {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSStreamListenerListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t TLSStreamListenerListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindTLSStreamListenerListenConfigConfig, data)
	return data, nil
}

//
// ========= End tls@stream/listener.ListenConfig =========

// ========= Begin tls_down@stream.Handler =========
//

const kindTLSDownStreamHandlerConfig = "tls_down@stream.Handler"

// TLSDownStreamHandlerConfig tls_down@stream.Handler
type TLSDownStreamHandlerConfig struct {
	Handler StreamHandler
	TLS     TLS
}

func init() {
	_ = defTypes.Register(
		kindTLSDownStreamHandlerConfig,
		func(r *TLSDownStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (TLSDownStreamHandlerConfig) isStreamHandler() {}
func (TLSDownStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSDownStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t TLSDownStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindTLSDownStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End tls_down@stream.Handler =========

// ========= Begin tls_up@stream.Handler =========
//

const kindTLSUpStreamHandlerConfig = "tls_up@stream.Handler"

// TLSUpStreamHandlerConfig tls_up@stream.Handler
type TLSUpStreamHandlerConfig struct {
	Handler StreamHandler
	TLS     TLS
}

func init() {
	_ = defTypes.Register(
		kindTLSUpStreamHandlerConfig,
		func(r *TLSUpStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (TLSUpStreamHandlerConfig) isStreamHandler() {}
func (TLSUpStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSUpStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t TLSUpStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindTLSUpStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End tls_up@stream.Handler =========

// ========= Begin transport@net/http.RoundTripper =========
//

const kindTransportNetHTTPRoundTripperConfig = "transport@net/http.RoundTripper"

// TransportNetHTTPRoundTripperConfig transport@net/http.RoundTripper
type TransportNetHTTPRoundTripperConfig struct {
	TLS    TLS
	Dialer Dialer
}

func init() {
	_ = defTypes.Register(
		kindTransportNetHTTPRoundTripperConfig,
		func(r *TransportNetHTTPRoundTripperConfig) HTTPRoundTripper {
			return r
		},
	)
}

func (TransportNetHTTPRoundTripperConfig) isHTTPRoundTripper() {}
func (TransportNetHTTPRoundTripperConfig) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TransportNetHTTPRoundTripperConfig) MarshalJSON() ([]byte, error) {
	type t TransportNetHTTPRoundTripperConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindTransportNetHTTPRoundTripperConfig, data)
	return data, nil
}

//
// ========= End transport@net/http.RoundTripper =========

// ========= Begin weighted@net/http.Handler =========
//

const kindWeightedNetHTTPHandlerConfig = "weighted@net/http.Handler"

// WeightedNetHTTPHandlerConfig weighted@net/http.Handler
type WeightedNetHTTPHandlerConfig struct {
	Weighted []WeightedNetHTTPHandlerWeighted
}

type WeightedNetHTTPHandlerWeighted struct {
	Weight  uint
	Handler HTTPHandler
}

func init() {
	_ = defTypes.Register(
		kindWeightedNetHTTPHandlerConfig,
		func(r *WeightedNetHTTPHandlerConfig) HTTPHandler {
			return r
		},
	)
}

func (WeightedNetHTTPHandlerConfig) isHTTPHandler() {}
func (WeightedNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m WeightedNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t WeightedNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindWeightedNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End weighted@net/http.Handler =========

// ========= Begin weighted@stream.Handler =========
//

const kindWeightedStreamHandlerConfig = "weighted@stream.Handler"

// WeightedStreamHandlerConfig weighted@stream.Handler
type WeightedStreamHandlerConfig struct {
	Weighted []WeightedStreamHandlerWeighted
}

type WeightedStreamHandlerWeighted struct {
	Weight  uint
	Handler StreamHandler
}

func init() {
	_ = defTypes.Register(
		kindWeightedStreamHandlerConfig,
		func(r *WeightedStreamHandlerConfig) StreamHandler {
			return r
		},
	)
}

func (WeightedStreamHandlerConfig) isStreamHandler() {}
func (WeightedStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m WeightedStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t WeightedStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindWeightedStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End weighted@stream.Handler =========

// ========= Begin yaml@codec.Marshaler =========
//

const kindYamlCodecMarshaler = "yaml@codec.Marshaler"

// YamlCodecMarshaler yaml@codec.Marshaler
type YamlCodecMarshaler struct {
}

func init() {
	_ = defTypes.Register(
		kindYamlCodecMarshaler,
		func(r *YamlCodecMarshaler) CodecMarshaler {
			return r
		},
	)
}

func (YamlCodecMarshaler) isCodecMarshaler() {}
func (YamlCodecMarshaler) isComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m YamlCodecMarshaler) MarshalJSON() ([]byte, error) {
	type t YamlCodecMarshaler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindYamlCodecMarshaler, data)
	return data, nil
}

//
// ========= End yaml@codec.Marshaler =========

// ========= Begin yaml@codec.Unmarshaler =========
//

const kindYamlCodecUnmarshaler = "yaml@codec.Unmarshaler"

// YamlCodecUnmarshaler yaml@codec.Unmarshaler
type YamlCodecUnmarshaler struct {
}

func init() {
	_ = defTypes.Register(
		kindYamlCodecUnmarshaler,
		func(r *YamlCodecUnmarshaler) CodecUnmarshaler {
			return r
		},
	)
}

func (YamlCodecUnmarshaler) isCodecUnmarshaler() {}
func (YamlCodecUnmarshaler) isComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m YamlCodecUnmarshaler) MarshalJSON() ([]byte, error) {
	type t YamlCodecUnmarshaler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindYamlCodecUnmarshaler, data)
	return data, nil
}

//
// ========= End yaml@codec.Unmarshaler =========
