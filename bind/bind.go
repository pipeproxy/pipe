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

const kindAcmeTLSTLSConfig = "acme@tls.TLS"

// AcmeTLSTLSConfig acme@tls.TLS
type AcmeTLSTLSConfig struct {
	Domains  []string
	CacheDir string
}

func init() {
	_ = defTypes.Register(
		kindAcmeTLSTLSConfig,
		func(r *AcmeTLSTLSConfig) TLS {
			return r
		},
	)
}

func (AcmeTLSTLSConfig) isTLS()       {}
func (AcmeTLSTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m AcmeTLSTLSConfig) MarshalJSON() ([]byte, error) {
	type t AcmeTLSTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindAcmeTLSTLSConfig, data)
	return data, nil
}

//
// ========= End acme@tls.TLS =========

// ========= Begin http.Handler =========
//

// Handler http.Handler
type Handler interface {
	isHandler()
	Component
}

// RawHandler is store raw bytes of Handler
type RawHandler []byte

func (RawHandler) isHandler()   {}
func (RawHandler) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawHandler: UnmarshalJSON on nil pointer")
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
		func(r *AddRequestHeaderNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (AddRequestHeaderNetHTTPHandlerConfig) isHandler()   {}
func (AddRequestHeaderNetHTTPHandlerConfig) isComponent() {}

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
		func(r *AddResponseHeaderNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (AddResponseHeaderNetHTTPHandlerConfig) isHandler()   {}
func (AddResponseHeaderNetHTTPHandlerConfig) isComponent() {}

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

// Decoder codec.Decoder
type Decoder interface {
	isDecoder()
	Component
}

// RawDecoder is store raw bytes of Decoder
type RawDecoder []byte

func (RawDecoder) isDecoder()   {}
func (RawDecoder) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawDecoder) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawDecoder) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawDecoder: UnmarshalJSON on nil pointer")
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
		func(r *Base32CodecDecoderConfig) Decoder {
			return r
		},
	)
}

func (Base32CodecDecoderConfig) isDecoder()   {}
func (Base32CodecDecoderConfig) isComponent() {}

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

// Encoder codec.Encoder
type Encoder interface {
	isEncoder()
	Component
}

// RawEncoder is store raw bytes of Encoder
type RawEncoder []byte

func (RawEncoder) isEncoder()   {}
func (RawEncoder) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawEncoder) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawEncoder) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawEncoder: UnmarshalJSON on nil pointer")
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
		func(r *Base32CodecEncoderConfig) Encoder {
			return r
		},
	)
}

func (Base32CodecEncoderConfig) isEncoder()   {}
func (Base32CodecEncoderConfig) isComponent() {}

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
		func(r *Base64CodecDecoderConfig) Decoder {
			return r
		},
	)
}

func (Base64CodecDecoderConfig) isDecoder()   {}
func (Base64CodecDecoderConfig) isComponent() {}

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
		func(r *Base64CodecEncoderConfig) Encoder {
			return r
		},
	)
}

func (Base64CodecEncoderConfig) isEncoder()   {}
func (Base64CodecEncoderConfig) isComponent() {}

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
		func(r *Bzip2CodecDecoder) Decoder {
			return r
		},
	)
}

func (Bzip2CodecDecoder) isDecoder()   {}
func (Bzip2CodecDecoder) isComponent() {}

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
	Handler Handler
}

func init() {
	_ = defTypes.Register(
		kindCompressNetHTTPHandlerConfig,
		func(r *CompressNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (CompressNetHTTPHandlerConfig) isHandler()   {}
func (CompressNetHTTPHandlerConfig) isComponent() {}

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
		func(r *ConfigDumpNetHTTPHandler) Handler {
			return r
		},
	)
}

func (ConfigDumpNetHTTPHandler) isHandler()   {}
func (ConfigDumpNetHTTPHandler) isComponent() {}

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
	Def  Decoder
}

func init() {
	_ = defTypes.Register(
		kindDefCodecDecoderConfig,
		func(r *DefCodecDecoderConfig) Decoder {
			return r
		},
	)
}

func (DefCodecDecoderConfig) isDecoder()   {}
func (DefCodecDecoderConfig) isComponent() {}

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
	Def  Encoder
}

func init() {
	_ = defTypes.Register(
		kindDefCodecEncoderConfig,
		func(r *DefCodecEncoderConfig) Encoder {
			return r
		},
	)
}

func (DefCodecEncoderConfig) isEncoder()   {}
func (DefCodecEncoderConfig) isComponent() {}

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

// Marshaler codec.Marshaler
type Marshaler interface {
	isMarshaler()
	Component
}

// RawMarshaler is store raw bytes of Marshaler
type RawMarshaler []byte

func (RawMarshaler) isMarshaler() {}
func (RawMarshaler) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawMarshaler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMarshaler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawMarshaler: UnmarshalJSON on nil pointer")
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
	Def  Marshaler
}

func init() {
	_ = defTypes.Register(
		kindDefCodecMarshalerConfig,
		func(r *DefCodecMarshalerConfig) Marshaler {
			return r
		},
	)
}

func (DefCodecMarshalerConfig) isMarshaler() {}
func (DefCodecMarshalerConfig) isComponent() {}

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

// Unmarshaler codec.Unmarshaler
type Unmarshaler interface {
	isUnmarshaler()
	Component
}

// RawUnmarshaler is store raw bytes of Unmarshaler
type RawUnmarshaler []byte

func (RawUnmarshaler) isUnmarshaler() {}
func (RawUnmarshaler) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawUnmarshaler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawUnmarshaler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawUnmarshaler: UnmarshalJSON on nil pointer")
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
	Def  Unmarshaler
}

func init() {
	_ = defTypes.Register(
		kindDefCodecUnmarshalerConfig,
		func(r *DefCodecUnmarshalerConfig) Unmarshaler {
			return r
		},
	)
}

func (DefCodecUnmarshalerConfig) isUnmarshaler() {}
func (DefCodecUnmarshalerConfig) isComponent()   {}

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

// Reader io.Reader
type Reader interface {
	isReader()
	Component
}

// RawReader is store raw bytes of Reader
type RawReader []byte

func (RawReader) isReader()    {}
func (RawReader) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawReader) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawReader) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawReader: UnmarshalJSON on nil pointer")
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
	Def  Reader
}

func init() {
	_ = defTypes.Register(
		kindDefIoReaderConfig,
		func(r *DefIoReaderConfig) Reader {
			return r
		},
	)
}

func (DefIoReaderConfig) isReader()    {}
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

// Writer io.Writer
type Writer interface {
	isWriter()
	Component
}

// RawWriter is store raw bytes of Writer
type RawWriter []byte

func (RawWriter) isWriter()    {}
func (RawWriter) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawWriter) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawWriter) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawWriter: UnmarshalJSON on nil pointer")
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
	Def  Writer
}

func init() {
	_ = defTypes.Register(
		kindDefIoWriterConfig,
		func(r *DefIoWriterConfig) Writer {
			return r
		},
	)
}

func (DefIoWriterConfig) isWriter()    {}
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

// ========= Begin http.RoundTripper =========
//

// RoundTripper http.RoundTripper
type RoundTripper interface {
	isRoundTripper()
	Component
}

// RawRoundTripper is store raw bytes of RoundTripper
type RawRoundTripper []byte

func (RawRoundTripper) isRoundTripper() {}
func (RawRoundTripper) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawRoundTripper) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawRoundTripper) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawRoundTripper: UnmarshalJSON on nil pointer")
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
	Def  RoundTripper
}

func init() {
	_ = defTypes.Register(
		kindDefNetHTTPRoundTripperConfig,
		func(r *DefNetHTTPRoundTripperConfig) RoundTripper {
			return r
		},
	)
}

func (DefNetHTTPRoundTripperConfig) isRoundTripper() {}
func (DefNetHTTPRoundTripperConfig) isComponent()    {}

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

const kindDefOnceOnceConfig = "def@once.Once"

// DefOnceOnceConfig def@once.Once
type DefOnceOnceConfig struct {
	Name string
	Def  Once
}

func init() {
	_ = defTypes.Register(
		kindDefOnceOnceConfig,
		func(r *DefOnceOnceConfig) Once {
			return r
		},
	)
}

func (DefOnceOnceConfig) isOnce()      {}
func (DefOnceOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefOnceOnceConfig) MarshalJSON() ([]byte, error) {
	type t DefOnceOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefOnceOnceConfig, data)
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

const kindDefServiceServiceConfig = "def@service.Service"

// DefServiceServiceConfig def@service.Service
type DefServiceServiceConfig struct {
	Name string
	Def  Service
}

func init() {
	_ = defTypes.Register(
		kindDefServiceServiceConfig,
		func(r *DefServiceServiceConfig) Service {
			return r
		},
	)
}

func (DefServiceServiceConfig) isService()   {}
func (DefServiceServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefServiceServiceConfig) MarshalJSON() ([]byte, error) {
	type t DefServiceServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefServiceServiceConfig, data)
	return data, nil
}

//
// ========= End def@service.Service =========

// ========= Begin def@stream.Handler =========
//

const kindDefStreamHandlerConfig = "def@stream.Handler"

// DefStreamHandlerConfig def@stream.Handler
type DefStreamHandlerConfig struct {
	Name string
	Def  Handler
}

func init() {
	_ = defTypes.Register(
		kindDefStreamHandlerConfig,
		func(r *DefStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (DefStreamHandlerConfig) isHandler()   {}
func (DefStreamHandlerConfig) isComponent() {}

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

const kindDefStreamDialerDialerConfig = "def@stream/dialer.Dialer"

// DefStreamDialerDialerConfig def@stream/dialer.Dialer
type DefStreamDialerDialerConfig struct {
	Name string
	Def  Dialer
}

func init() {
	_ = defTypes.Register(
		kindDefStreamDialerDialerConfig,
		func(r *DefStreamDialerDialerConfig) Dialer {
			return r
		},
	)
}

func (DefStreamDialerDialerConfig) isDialer()    {}
func (DefStreamDialerDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefStreamDialerDialerConfig) MarshalJSON() ([]byte, error) {
	type t DefStreamDialerDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefStreamDialerDialerConfig, data)
	return data, nil
}

//
// ========= End def@stream/dialer.Dialer =========

// ========= Begin listener.ListenConfig =========
//

// ListenConfig listener.ListenConfig
type ListenConfig interface {
	isListenConfig()
	Component
}

// RawListenConfig is store raw bytes of ListenConfig
type RawListenConfig []byte

func (RawListenConfig) isListenConfig() {}
func (RawListenConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawListenConfig) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawListenConfig) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawListenConfig: UnmarshalJSON on nil pointer")
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
	Def  ListenConfig
}

func init() {
	_ = defTypes.Register(
		kindDefStreamListenerListenConfigConfig,
		func(r *DefStreamListenerListenConfigConfig) ListenConfig {
			return r
		},
	)
}

func (DefStreamListenerListenConfigConfig) isListenConfig() {}
func (DefStreamListenerListenConfigConfig) isComponent()    {}

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

const kindDefTLSTLSConfig = "def@tls.TLS"

// DefTLSTLSConfig def@tls.TLS
type DefTLSTLSConfig struct {
	Name string
	Def  TLS
}

func init() {
	_ = defTypes.Register(
		kindDefTLSTLSConfig,
		func(r *DefTLSTLSConfig) TLS {
			return r
		},
	)
}

func (DefTLSTLSConfig) isTLS()       {}
func (DefTLSTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefTLSTLSConfig) MarshalJSON() ([]byte, error) {
	type t DefTLSTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindDefTLSTLSConfig, data)
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
	Body Reader
}

func init() {
	_ = defTypes.Register(
		kindDirectNetHTTPHandlerConfig,
		func(r *DirectNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (DirectNetHTTPHandlerConfig) isHandler()   {}
func (DirectNetHTTPHandlerConfig) isComponent() {}

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
		func(r *ExpvarNetHTTPHandler) Handler {
			return r
		},
	)
}

func (ExpvarNetHTTPHandler) isHandler()   {}
func (ExpvarNetHTTPHandler) isComponent() {}

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
		func(r *FileIoReaderConfig) Reader {
			return r
		},
	)
}

func (FileIoReaderConfig) isReader()    {}
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
		func(r *FileIoWriterConfig) Writer {
			return r
		},
	)
}

func (FileIoWriterConfig) isWriter()    {}
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
		func(r *FileNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (FileNetHTTPHandlerConfig) isHandler()   {}
func (FileNetHTTPHandlerConfig) isComponent() {}

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
	RoundTripper RoundTripper
	URL          string
}

func init() {
	_ = defTypes.Register(
		kindForwardNetHTTPHandlerConfig,
		func(r *ForwardNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (ForwardNetHTTPHandlerConfig) isHandler()   {}
func (ForwardNetHTTPHandlerConfig) isComponent() {}

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
		func(r *ForwardStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (ForwardStreamHandlerConfig) isHandler()   {}
func (ForwardStreamHandlerConfig) isComponent() {}

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

const kindFromTLSTLSConfig = "from@tls.TLS"

// FromTLSTLSConfig from@tls.TLS
type FromTLSTLSConfig struct {
	Domain string
	Cert   Reader
	Key    Reader
}

func init() {
	_ = defTypes.Register(
		kindFromTLSTLSConfig,
		func(r *FromTLSTLSConfig) TLS {
			return r
		},
	)
}

func (FromTLSTLSConfig) isTLS()       {}
func (FromTLSTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m FromTLSTLSConfig) MarshalJSON() ([]byte, error) {
	type t FromTLSTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindFromTLSTLSConfig, data)
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
		func(r *GzipCodecDecoder) Decoder {
			return r
		},
	)
}

func (GzipCodecDecoder) isDecoder()   {}
func (GzipCodecDecoder) isComponent() {}

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
		func(r *GzipCodecEncoder) Encoder {
			return r
		},
	)
}

func (GzipCodecEncoder) isEncoder()   {}
func (GzipCodecEncoder) isComponent() {}

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
	Handler Handler
}

func init() {
	_ = defTypes.Register(
		kindH2CNetHTTPHandlerConfig,
		func(r *H2CNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (H2CNetHTTPHandlerConfig) isHandler()   {}
func (H2CNetHTTPHandlerConfig) isComponent() {}

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
		func(r *HexCodecDecoder) Decoder {
			return r
		},
	)
}

func (HexCodecDecoder) isDecoder()   {}
func (HexCodecDecoder) isComponent() {}

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
		func(r *HexCodecEncoder) Encoder {
			return r
		},
	)
}

func (HexCodecEncoder) isEncoder()   {}
func (HexCodecEncoder) isComponent() {}

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
	NotFound Handler
}

type HostNetHTTPHandlerRoute struct {
	Domain  string
	Handler Handler
}

func init() {
	_ = defTypes.Register(
		kindHostNetHTTPHandlerConfig,
		func(r *HostNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (HostNetHTTPHandlerConfig) isHandler()   {}
func (HostNetHTTPHandlerConfig) isComponent() {}

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
	Handler Handler
	TLS     TLS
}

func init() {
	_ = defTypes.Register(
		kindHTTPStreamHandlerConfig,
		func(r *HTTPStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (HTTPStreamHandlerConfig) isHandler()   {}
func (HTTPStreamHandlerConfig) isComponent() {}

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
		func(r *InlineIoReaderConfig) Reader {
			return r
		},
	)
}

func (InlineIoReaderConfig) isReader()    {}
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
		func(r *JSONCodecMarshaler) Marshaler {
			return r
		},
	)
}

func (JSONCodecMarshaler) isMarshaler() {}
func (JSONCodecMarshaler) isComponent() {}

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
		func(r *JSONCodecUnmarshaler) Unmarshaler {
			return r
		},
	)
}

func (JSONCodecUnmarshaler) isUnmarshaler() {}
func (JSONCodecUnmarshaler) isComponent()   {}

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
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadCodecDecoderConfig,
		func(r *LoadCodecDecoderConfig) Decoder {
			return r
		},
	)
}

func (LoadCodecDecoderConfig) isDecoder()   {}
func (LoadCodecDecoderConfig) isComponent() {}

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
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadCodecEncoderConfig,
		func(r *LoadCodecEncoderConfig) Encoder {
			return r
		},
	)
}

func (LoadCodecEncoderConfig) isEncoder()   {}
func (LoadCodecEncoderConfig) isComponent() {}

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
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadCodecMarshalerConfig,
		func(r *LoadCodecMarshalerConfig) Marshaler {
			return r
		},
	)
}

func (LoadCodecMarshalerConfig) isMarshaler() {}
func (LoadCodecMarshalerConfig) isComponent() {}

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
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadCodecUnmarshalerConfig,
		func(r *LoadCodecUnmarshalerConfig) Unmarshaler {
			return r
		},
	)
}

func (LoadCodecUnmarshalerConfig) isUnmarshaler() {}
func (LoadCodecUnmarshalerConfig) isComponent()   {}

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
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadIoReaderConfig,
		func(r *LoadIoReaderConfig) Reader {
			return r
		},
	)
}

func (LoadIoReaderConfig) isReader()    {}
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
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadIoWriterConfig,
		func(r *LoadIoWriterConfig) Writer {
			return r
		},
	)
}

func (LoadIoWriterConfig) isWriter()    {}
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

// ========= Begin load@net/http.RoundTripper =========
//

const kindLoadNetHTTPRoundTripperConfig = "load@net/http.RoundTripper"

// LoadNetHTTPRoundTripperConfig load@net/http.RoundTripper
type LoadNetHTTPRoundTripperConfig struct {
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadNetHTTPRoundTripperConfig,
		func(r *LoadNetHTTPRoundTripperConfig) RoundTripper {
			return r
		},
	)
}

func (LoadNetHTTPRoundTripperConfig) isRoundTripper() {}
func (LoadNetHTTPRoundTripperConfig) isComponent()    {}

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

const kindLoadOnceOnceConfig = "load@once.Once"

// LoadOnceOnceConfig load@once.Once
type LoadOnceOnceConfig struct {
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadOnceOnceConfig,
		func(r *LoadOnceOnceConfig) Once {
			return r
		},
	)
}

func (LoadOnceOnceConfig) isOnce()      {}
func (LoadOnceOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadOnceOnceConfig) MarshalJSON() ([]byte, error) {
	type t LoadOnceOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadOnceOnceConfig, data)
	return data, nil
}

//
// ========= End load@once.Once =========

// ========= Begin load@service.Service =========
//

const kindLoadServiceServiceConfig = "load@service.Service"

// LoadServiceServiceConfig load@service.Service
type LoadServiceServiceConfig struct {
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadServiceServiceConfig,
		func(r *LoadServiceServiceConfig) Service {
			return r
		},
	)
}

func (LoadServiceServiceConfig) isService()   {}
func (LoadServiceServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadServiceServiceConfig) MarshalJSON() ([]byte, error) {
	type t LoadServiceServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadServiceServiceConfig, data)
	return data, nil
}

//
// ========= End load@service.Service =========

// ========= Begin load@stream.Handler =========
//

const kindLoadStreamHandlerConfig = "load@stream.Handler"

// LoadStreamHandlerConfig load@stream.Handler
type LoadStreamHandlerConfig struct {
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadStreamHandlerConfig,
		func(r *LoadStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (LoadStreamHandlerConfig) isHandler()   {}
func (LoadStreamHandlerConfig) isComponent() {}

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

const kindLoadStreamDialerDialerConfig = "load@stream/dialer.Dialer"

// LoadStreamDialerDialerConfig load@stream/dialer.Dialer
type LoadStreamDialerDialerConfig struct {
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadStreamDialerDialerConfig,
		func(r *LoadStreamDialerDialerConfig) Dialer {
			return r
		},
	)
}

func (LoadStreamDialerDialerConfig) isDialer()    {}
func (LoadStreamDialerDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadStreamDialerDialerConfig) MarshalJSON() ([]byte, error) {
	type t LoadStreamDialerDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadStreamDialerDialerConfig, data)
	return data, nil
}

//
// ========= End load@stream/dialer.Dialer =========

// ========= Begin load@stream/listener.ListenConfig =========
//

const kindLoadStreamListenerListenConfigConfig = "load@stream/listener.ListenConfig"

// LoadStreamListenerListenConfigConfig load@stream/listener.ListenConfig
type LoadStreamListenerListenConfigConfig struct {
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadStreamListenerListenConfigConfig,
		func(r *LoadStreamListenerListenConfigConfig) ListenConfig {
			return r
		},
	)
}

func (LoadStreamListenerListenConfigConfig) isListenConfig() {}
func (LoadStreamListenerListenConfigConfig) isComponent()    {}

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

const kindLoadTLSTLSConfig = "load@tls.TLS"

// LoadTLSTLSConfig load@tls.TLS
type LoadTLSTLSConfig struct {
	Load Reader
}

func init() {
	_ = defTypes.Register(
		kindLoadTLSTLSConfig,
		func(r *LoadTLSTLSConfig) TLS {
			return r
		},
	)
}

func (LoadTLSTLSConfig) isTLS()       {}
func (LoadTLSTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadTLSTLSConfig) MarshalJSON() ([]byte, error) {
	type t LoadTLSTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindLoadTLSTLSConfig, data)
	return data, nil
}

//
// ========= End load@tls.TLS =========

// ========= Begin log@net/http.Handler =========
//

const kindLogNetHTTPHandlerConfig = "log@net/http.Handler"

// LogNetHTTPHandlerConfig log@net/http.Handler
type LogNetHTTPHandlerConfig struct {
	Output  Writer
	Handler Handler
}

func init() {
	_ = defTypes.Register(
		kindLogNetHTTPHandlerConfig,
		func(r *LogNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (LogNetHTTPHandlerConfig) isHandler()   {}
func (LogNetHTTPHandlerConfig) isComponent() {}

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

const kindMessageOnceOnceConfig = "message@once.Once"

// MessageOnceOnceConfig message@once.Once
type MessageOnceOnceConfig struct {
	Message string
}

func init() {
	_ = defTypes.Register(
		kindMessageOnceOnceConfig,
		func(r *MessageOnceOnceConfig) Once {
			return r
		},
	)
}

func (MessageOnceOnceConfig) isOnce()      {}
func (MessageOnceOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MessageOnceOnceConfig) MarshalJSON() ([]byte, error) {
	type t MessageOnceOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMessageOnceOnceConfig, data)
	return data, nil
}

//
// ========= End message@once.Once =========

// ========= Begin multi@net/http.Handler =========
//

const kindMultiNetHTTPHandlerConfig = "multi@net/http.Handler"

// MultiNetHTTPHandlerConfig multi@net/http.Handler
type MultiNetHTTPHandlerConfig struct {
	Multi []Handler
}

func init() {
	_ = defTypes.Register(
		kindMultiNetHTTPHandlerConfig,
		func(r *MultiNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (MultiNetHTTPHandlerConfig) isHandler()   {}
func (MultiNetHTTPHandlerConfig) isComponent() {}

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

const kindMultiOnceOnceConfig = "multi@once.Once"

// MultiOnceOnceConfig multi@once.Once
type MultiOnceOnceConfig struct {
	Multi []Once
}

func init() {
	_ = defTypes.Register(
		kindMultiOnceOnceConfig,
		func(r *MultiOnceOnceConfig) Once {
			return r
		},
	)
}

func (MultiOnceOnceConfig) isOnce()      {}
func (MultiOnceOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MultiOnceOnceConfig) MarshalJSON() ([]byte, error) {
	type t MultiOnceOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMultiOnceOnceConfig, data)
	return data, nil
}

//
// ========= End multi@once.Once =========

// ========= Begin multi@service.Service =========
//

const kindMultiServiceServiceConfig = "multi@service.Service"

// MultiServiceServiceConfig multi@service.Service
type MultiServiceServiceConfig struct {
	Multi []Service
}

func init() {
	_ = defTypes.Register(
		kindMultiServiceServiceConfig,
		func(r *MultiServiceServiceConfig) Service {
			return r
		},
	)
}

func (MultiServiceServiceConfig) isService()   {}
func (MultiServiceServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MultiServiceServiceConfig) MarshalJSON() ([]byte, error) {
	type t MultiServiceServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindMultiServiceServiceConfig, data)
	return data, nil
}

//
// ========= End multi@service.Service =========

// ========= Begin multi@stream.Handler =========
//

const kindMultiStreamHandlerConfig = "multi@stream.Handler"

// MultiStreamHandlerConfig multi@stream.Handler
type MultiStreamHandlerConfig struct {
	Multi []Handler
}

func init() {
	_ = defTypes.Register(
		kindMultiStreamHandlerConfig,
		func(r *MultiStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (MultiStreamHandlerConfig) isHandler()   {}
func (MultiStreamHandlerConfig) isComponent() {}

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
	NotFound Handler
}

type MuxNetHTTPHandlerRoute struct {
	Prefix  string
	Path    string
	Regexp  string
	Handler Handler
}

func init() {
	_ = defTypes.Register(
		kindMuxNetHTTPHandlerConfig,
		func(r *MuxNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (MuxNetHTTPHandlerConfig) isHandler()   {}
func (MuxNetHTTPHandlerConfig) isComponent() {}

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
	NotFound Handler
}

type MuxStreamHandlerRoute struct {
	Pattern string
	Regexp  string
	Prefix  string
	Handler Handler
}

func init() {
	_ = defTypes.Register(
		kindMuxStreamHandlerConfig,
		func(r *MuxStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (MuxStreamHandlerConfig) isHandler()   {}
func (MuxStreamHandlerConfig) isComponent() {}

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

const kindNetworkStreamDialerDialerConfig = "network@stream/dialer.Dialer"

// NetworkStreamDialerDialerConfig network@stream/dialer.Dialer
type NetworkStreamDialerDialerConfig struct {
	Network NetworkStreamDialerDialerNetworkEnum
	Address string
}

type NetworkStreamDialerDialerNetworkEnum string

const (
	NetworkStreamDialerDialerNetworkEnumEnumUnix NetworkStreamDialerDialerNetworkEnum = "unix"
	NetworkStreamDialerDialerNetworkEnumEnumTCP6 NetworkStreamDialerDialerNetworkEnum = "tcp6"
	NetworkStreamDialerDialerNetworkEnumEnumTCP4 NetworkStreamDialerDialerNetworkEnum = "tcp4"
	NetworkStreamDialerDialerNetworkEnumEnumTCP  NetworkStreamDialerDialerNetworkEnum = "tcp"
)

func init() {
	_ = defTypes.Register(
		kindNetworkStreamDialerDialerConfig,
		func(r *NetworkStreamDialerDialerConfig) Dialer {
			return r
		},
	)
}

func (NetworkStreamDialerDialerConfig) isDialer()    {}
func (NetworkStreamDialerDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NetworkStreamDialerDialerConfig) MarshalJSON() ([]byte, error) {
	type t NetworkStreamDialerDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindNetworkStreamDialerDialerConfig, data)
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
		func(r *NetworkStreamListenerListenConfigConfig) ListenConfig {
			return r
		},
	)
}

func (NetworkStreamListenerListenConfigConfig) isListenConfig() {}
func (NetworkStreamListenerListenConfigConfig) isComponent()    {}

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

const kindNoneOnceOnceConfig = "none@once.Once"

// NoneOnceOnceConfig none@once.Once
type NoneOnceOnceConfig struct {
	Any Component
}

func init() {
	_ = defTypes.Register(
		kindNoneOnceOnceConfig,
		func(r *NoneOnceOnceConfig) Once {
			return r
		},
	)
}

func (NoneOnceOnceConfig) isOnce()      {}
func (NoneOnceOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneOnceOnceConfig) MarshalJSON() ([]byte, error) {
	type t NoneOnceOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindNoneOnceOnceConfig, data)
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
	Handlers []Handler
}

type PollerNetHTTPHandlerPollerEnum string

const (
	PollerNetHTTPHandlerPollerEnumEnumRoundRobin PollerNetHTTPHandlerPollerEnum = "round_robin"
	PollerNetHTTPHandlerPollerEnumEnumRandom     PollerNetHTTPHandlerPollerEnum = "random"
)

func init() {
	_ = defTypes.Register(
		kindPollerNetHTTPHandlerConfig,
		func(r *PollerNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (PollerNetHTTPHandlerConfig) isHandler()   {}
func (PollerNetHTTPHandlerConfig) isComponent() {}

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
	Handlers []Handler
}

type PollerStreamHandlerPollerEnum string

const (
	PollerStreamHandlerPollerEnumEnumRoundRobin PollerStreamHandlerPollerEnum = "round_robin"
	PollerStreamHandlerPollerEnumEnumRandom     PollerStreamHandlerPollerEnum = "random"
)

func init() {
	_ = defTypes.Register(
		kindPollerStreamHandlerConfig,
		func(r *PollerStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (PollerStreamHandlerConfig) isHandler()   {}
func (PollerStreamHandlerConfig) isComponent() {}

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

const kindPollerStreamDialerDialerConfig = "poller@stream/dialer.Dialer"

// PollerStreamDialerDialerConfig poller@stream/dialer.Dialer
type PollerStreamDialerDialerConfig struct {
	Poller  PollerStreamDialerDialerPollerEnum
	Dialers []Dialer
}

type PollerStreamDialerDialerPollerEnum string

const (
	PollerStreamDialerDialerPollerEnumEnumRoundRobin PollerStreamDialerDialerPollerEnum = "round_robin"
	PollerStreamDialerDialerPollerEnumEnumRandom     PollerStreamDialerDialerPollerEnum = "random"
)

func init() {
	_ = defTypes.Register(
		kindPollerStreamDialerDialerConfig,
		func(r *PollerStreamDialerDialerConfig) Dialer {
			return r
		},
	)
}

func (PollerStreamDialerDialerConfig) isDialer()    {}
func (PollerStreamDialerDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m PollerStreamDialerDialerConfig) MarshalJSON() ([]byte, error) {
	type t PollerStreamDialerDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindPollerStreamDialerDialerConfig, data)
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
		func(r *PprofNetHTTPHandler) Handler {
			return r
		},
	)
}

func (PprofNetHTTPHandler) isHandler()   {}
func (PprofNetHTTPHandler) isComponent() {}

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
		func(r *RedirectNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (RedirectNetHTTPHandlerConfig) isHandler()   {}
func (RedirectNetHTTPHandlerConfig) isComponent() {}

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
	Def  Decoder
}

func init() {
	_ = defTypes.Register(
		kindRefCodecDecoderConfig,
		func(r *RefCodecDecoderConfig) Decoder {
			return r
		},
	)
}

func (RefCodecDecoderConfig) isDecoder()   {}
func (RefCodecDecoderConfig) isComponent() {}

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
	Def  Encoder
}

func init() {
	_ = defTypes.Register(
		kindRefCodecEncoderConfig,
		func(r *RefCodecEncoderConfig) Encoder {
			return r
		},
	)
}

func (RefCodecEncoderConfig) isEncoder()   {}
func (RefCodecEncoderConfig) isComponent() {}

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
	Def  Marshaler
}

func init() {
	_ = defTypes.Register(
		kindRefCodecMarshalerConfig,
		func(r *RefCodecMarshalerConfig) Marshaler {
			return r
		},
	)
}

func (RefCodecMarshalerConfig) isMarshaler() {}
func (RefCodecMarshalerConfig) isComponent() {}

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
	Def  Unmarshaler
}

func init() {
	_ = defTypes.Register(
		kindRefCodecUnmarshalerConfig,
		func(r *RefCodecUnmarshalerConfig) Unmarshaler {
			return r
		},
	)
}

func (RefCodecUnmarshalerConfig) isUnmarshaler() {}
func (RefCodecUnmarshalerConfig) isComponent()   {}

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
	Def  Reader
}

func init() {
	_ = defTypes.Register(
		kindRefIoReaderConfig,
		func(r *RefIoReaderConfig) Reader {
			return r
		},
	)
}

func (RefIoReaderConfig) isReader()    {}
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
	Def  Writer
}

func init() {
	_ = defTypes.Register(
		kindRefIoWriterConfig,
		func(r *RefIoWriterConfig) Writer {
			return r
		},
	)
}

func (RefIoWriterConfig) isWriter()    {}
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

// ========= Begin ref@net/http.RoundTripper =========
//

const kindRefNetHTTPRoundTripperConfig = "ref@net/http.RoundTripper"

// RefNetHTTPRoundTripperConfig ref@net/http.RoundTripper
type RefNetHTTPRoundTripperConfig struct {
	Name string
	Def  RoundTripper
}

func init() {
	_ = defTypes.Register(
		kindRefNetHTTPRoundTripperConfig,
		func(r *RefNetHTTPRoundTripperConfig) RoundTripper {
			return r
		},
	)
}

func (RefNetHTTPRoundTripperConfig) isRoundTripper() {}
func (RefNetHTTPRoundTripperConfig) isComponent()    {}

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

const kindRefOnceOnceConfig = "ref@once.Once"

// RefOnceOnceConfig ref@once.Once
type RefOnceOnceConfig struct {
	Name string
	Def  Once
}

func init() {
	_ = defTypes.Register(
		kindRefOnceOnceConfig,
		func(r *RefOnceOnceConfig) Once {
			return r
		},
	)
}

func (RefOnceOnceConfig) isOnce()      {}
func (RefOnceOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefOnceOnceConfig) MarshalJSON() ([]byte, error) {
	type t RefOnceOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefOnceOnceConfig, data)
	return data, nil
}

//
// ========= End ref@once.Once =========

// ========= Begin ref@service.Service =========
//

const kindRefServiceServiceConfig = "ref@service.Service"

// RefServiceServiceConfig ref@service.Service
type RefServiceServiceConfig struct {
	Name string
	Def  Service
}

func init() {
	_ = defTypes.Register(
		kindRefServiceServiceConfig,
		func(r *RefServiceServiceConfig) Service {
			return r
		},
	)
}

func (RefServiceServiceConfig) isService()   {}
func (RefServiceServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefServiceServiceConfig) MarshalJSON() ([]byte, error) {
	type t RefServiceServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefServiceServiceConfig, data)
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
	Def  Handler
}

func init() {
	_ = defTypes.Register(
		kindRefStreamHandlerConfig,
		func(r *RefStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (RefStreamHandlerConfig) isHandler()   {}
func (RefStreamHandlerConfig) isComponent() {}

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

const kindRefStreamDialerDialerConfig = "ref@stream/dialer.Dialer"

// RefStreamDialerDialerConfig ref@stream/dialer.Dialer
type RefStreamDialerDialerConfig struct {
	Name string
	Def  Dialer
}

func init() {
	_ = defTypes.Register(
		kindRefStreamDialerDialerConfig,
		func(r *RefStreamDialerDialerConfig) Dialer {
			return r
		},
	)
}

func (RefStreamDialerDialerConfig) isDialer()    {}
func (RefStreamDialerDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefStreamDialerDialerConfig) MarshalJSON() ([]byte, error) {
	type t RefStreamDialerDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefStreamDialerDialerConfig, data)
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
	Def  ListenConfig
}

func init() {
	_ = defTypes.Register(
		kindRefStreamListenerListenConfigConfig,
		func(r *RefStreamListenerListenConfigConfig) ListenConfig {
			return r
		},
	)
}

func (RefStreamListenerListenConfigConfig) isListenConfig() {}
func (RefStreamListenerListenConfigConfig) isComponent()    {}

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

const kindRefTLSTLSConfig = "ref@tls.TLS"

// RefTLSTLSConfig ref@tls.TLS
type RefTLSTLSConfig struct {
	Name string
	Def  TLS
}

func init() {
	_ = defTypes.Register(
		kindRefTLSTLSConfig,
		func(r *RefTLSTLSConfig) TLS {
			return r
		},
	)
}

func (RefTLSTLSConfig) isTLS()       {}
func (RefTLSTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefTLSTLSConfig) MarshalJSON() ([]byte, error) {
	type t RefTLSTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindRefTLSTLSConfig, data)
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
		func(r *RemoveRequestHeaderNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (RemoveRequestHeaderNetHTTPHandlerConfig) isHandler()   {}
func (RemoveRequestHeaderNetHTTPHandlerConfig) isComponent() {}

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
		func(r *RemoveResponseHeaderNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (RemoveResponseHeaderNetHTTPHandlerConfig) isHandler()   {}
func (RemoveResponseHeaderNetHTTPHandlerConfig) isComponent() {}

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

// ========= Begin self_signed@tls.TLS =========
//

const kindSelfSignedTLSTLS = "self_signed@tls.TLS"

// SelfSignedTLSTLS self_signed@tls.TLS
type SelfSignedTLSTLS struct {
}

func init() {
	_ = defTypes.Register(
		kindSelfSignedTLSTLS,
		func(r *SelfSignedTLSTLS) TLS {
			return r
		},
	)
}

func (SelfSignedTLSTLS) isTLS()       {}
func (SelfSignedTLSTLS) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m SelfSignedTLSTLS) MarshalJSON() ([]byte, error) {
	type t SelfSignedTLSTLS
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindSelfSignedTLSTLS, data)
	return data, nil
}

//
// ========= End self_signed@tls.TLS =========

// ========= Begin service@once.Once =========
//

const kindServiceOnceOnceConfig = "service@once.Once"

// ServiceOnceOnceConfig service@once.Once
type ServiceOnceOnceConfig struct {
	Service Service
}

func init() {
	_ = defTypes.Register(
		kindServiceOnceOnceConfig,
		func(r *ServiceOnceOnceConfig) Once {
			return r
		},
	)
}

func (ServiceOnceOnceConfig) isOnce()      {}
func (ServiceOnceOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceOnceOnceConfig) MarshalJSON() ([]byte, error) {
	type t ServiceOnceOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindServiceOnceOnceConfig, data)
	return data, nil
}

//
// ========= End service@once.Once =========

// ========= Begin stream@service.Service =========
//

const kindStreamServiceServiceConfig = "stream@service.Service"

// StreamServiceServiceConfig stream@service.Service
type StreamServiceServiceConfig struct {
	Listener ListenConfig
	Handler  Handler
}

func init() {
	_ = defTypes.Register(
		kindStreamServiceServiceConfig,
		func(r *StreamServiceServiceConfig) Service {
			return r
		},
	)
}

func (StreamServiceServiceConfig) isService()   {}
func (StreamServiceServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamServiceServiceConfig) MarshalJSON() ([]byte, error) {
	type t StreamServiceServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindStreamServiceServiceConfig, data)
	return data, nil
}

//
// ========= End stream@service.Service =========

// ========= Begin tls@stream/dialer.Dialer =========
//

const kindTLSStreamDialerDialerConfig = "tls@stream/dialer.Dialer"

// TLSStreamDialerDialerConfig tls@stream/dialer.Dialer
type TLSStreamDialerDialerConfig struct {
	Dialer Dialer
	TLS    TLS
}

func init() {
	_ = defTypes.Register(
		kindTLSStreamDialerDialerConfig,
		func(r *TLSStreamDialerDialerConfig) Dialer {
			return r
		},
	)
}

func (TLSStreamDialerDialerConfig) isDialer()    {}
func (TLSStreamDialerDialerConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSStreamDialerDialerConfig) MarshalJSON() ([]byte, error) {
	type t TLSStreamDialerDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = appendKV(kindKey, kindTLSStreamDialerDialerConfig, data)
	return data, nil
}

//
// ========= End tls@stream/dialer.Dialer =========

// ========= Begin tls@stream/listener.ListenConfig =========
//

const kindTLSStreamListenerListenConfigConfig = "tls@stream/listener.ListenConfig"

// TLSStreamListenerListenConfigConfig tls@stream/listener.ListenConfig
type TLSStreamListenerListenConfigConfig struct {
	ListenConfig ListenConfig
	TLS          TLS
}

func init() {
	_ = defTypes.Register(
		kindTLSStreamListenerListenConfigConfig,
		func(r *TLSStreamListenerListenConfigConfig) ListenConfig {
			return r
		},
	)
}

func (TLSStreamListenerListenConfigConfig) isListenConfig() {}
func (TLSStreamListenerListenConfigConfig) isComponent()    {}

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
	Handler Handler
	TLS     TLS
}

func init() {
	_ = defTypes.Register(
		kindTLSDownStreamHandlerConfig,
		func(r *TLSDownStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (TLSDownStreamHandlerConfig) isHandler()   {}
func (TLSDownStreamHandlerConfig) isComponent() {}

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
	Handler Handler
	TLS     TLS
}

func init() {
	_ = defTypes.Register(
		kindTLSUpStreamHandlerConfig,
		func(r *TLSUpStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (TLSUpStreamHandlerConfig) isHandler()   {}
func (TLSUpStreamHandlerConfig) isComponent() {}

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
		func(r *TransportNetHTTPRoundTripperConfig) RoundTripper {
			return r
		},
	)
}

func (TransportNetHTTPRoundTripperConfig) isRoundTripper() {}
func (TransportNetHTTPRoundTripperConfig) isComponent()    {}

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
	Handler Handler
}

func init() {
	_ = defTypes.Register(
		kindWeightedNetHTTPHandlerConfig,
		func(r *WeightedNetHTTPHandlerConfig) Handler {
			return r
		},
	)
}

func (WeightedNetHTTPHandlerConfig) isHandler()   {}
func (WeightedNetHTTPHandlerConfig) isComponent() {}

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
	Handler Handler
}

func init() {
	_ = defTypes.Register(
		kindWeightedStreamHandlerConfig,
		func(r *WeightedStreamHandlerConfig) Handler {
			return r
		},
	)
}

func (WeightedStreamHandlerConfig) isHandler()   {}
func (WeightedStreamHandlerConfig) isComponent() {}

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
		func(r *YamlCodecMarshaler) Marshaler {
			return r
		},
	)
}

func (YamlCodecMarshaler) isMarshaler() {}
func (YamlCodecMarshaler) isComponent() {}

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
		func(r *YamlCodecUnmarshaler) Unmarshaler {
			return r
		},
	)
}

func (YamlCodecUnmarshaler) isUnmarshaler() {}
func (YamlCodecUnmarshaler) isComponent()   {}

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
