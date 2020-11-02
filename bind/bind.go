// DO NOT EDIT! Code generated.

package bind

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wzshiming/funcfg/types"
	"github.com/wzshiming/funcfg/unmarshaler"
)

// ========= Begin Common =========
//

var kindKey = `@Kind`

var provider = types.NewEmptyProvider()

// Unmarshal parses the encoded data and stores the result
func Unmarshal(config []byte, v interface{}) error {
	u := unmarshaler.Unmarshaler{
		Ctx:      context.Background(),
		Provider: provider,
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

func prepend(k, v string, data []byte) []byte {
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf(`{%q:%q}`, k, v))
		} else {
			data = append([]byte(fmt.Sprintf(`{%q:%q,`, k, v)), data[1:]...)
		}
	}
	return data
}

//
// ========= End Common =========

// ========= Begin acme@tls.TLS type =========
//

const kindAcmeTLSConfig = `acme@tls.TLS`

// AcmeTLSConfig acme@tls.TLS
type AcmeTLSConfig struct {
	Domains  []string
	CacheDir string
}

func init() {
	_ = provider.Register(
		kindAcmeTLSConfig,
		func(r *AcmeTLSConfig) TLS { return r },
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
	data = prepend(kindKey, kindAcmeTLSConfig, data)
	return data, nil
}

//
// ========= End acme@tls.TLS type =========

// ========= Begin add_request_header@net/http.Handler type =========
//

const kindAddRequestHeaderNetHTTPHandlerConfig = `add_request_header@net/http.Handler`

// AddRequestHeaderNetHTTPHandlerConfig add_request_header@net/http.Handler
type AddRequestHeaderNetHTTPHandlerConfig struct {
	Key   string
	Value string
}

func init() {
	_ = provider.Register(
		kindAddRequestHeaderNetHTTPHandlerConfig,
		func(r *AddRequestHeaderNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindAddRequestHeaderNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End add_request_header@net/http.Handler type =========

// ========= Begin add_response_header@net/http.Handler type =========
//

const kindAddResponseHeaderNetHTTPHandlerConfig = `add_response_header@net/http.Handler`

// AddResponseHeaderNetHTTPHandlerConfig add_response_header@net/http.Handler
type AddResponseHeaderNetHTTPHandlerConfig struct {
	Key   string
	Value string
}

func init() {
	_ = provider.Register(
		kindAddResponseHeaderNetHTTPHandlerConfig,
		func(r *AddResponseHeaderNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindAddResponseHeaderNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End add_response_header@net/http.Handler type =========

// ========= Begin components@once.Once type =========
//

const kindComponentsOnceConfig = `components@once.Once`

// ComponentsOnceConfig components@once.Once
type ComponentsOnceConfig struct {
	Components []Component
}

func init() {
	_ = provider.Register(
		kindComponentsOnceConfig,
		func(r *ComponentsOnceConfig) Once { return r },
	)
}

func (ComponentsOnceConfig) isOnce()      {}
func (ComponentsOnceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ComponentsOnceConfig) MarshalJSON() ([]byte, error) {
	type t ComponentsOnceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindComponentsOnceConfig, data)
	return data, nil
}

//
// ========= End components@once.Once type =========

// ========= Begin compress@net/http.Handler type =========
//

const kindCompressNetHTTPHandlerConfig = `compress@net/http.Handler`

// CompressNetHTTPHandlerConfig compress@net/http.Handler
type CompressNetHTTPHandlerConfig struct {
	Level   int
	Handler HTTPHandler
}

func init() {
	_ = provider.Register(
		kindCompressNetHTTPHandlerConfig,
		func(r *CompressNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindCompressNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End compress@net/http.Handler type =========

// ========= Begin config_dump@net/http.Handler type =========
//

const kindConfigDumpNetHTTPHandlerConfig = `config_dump@net/http.Handler`

// ConfigDumpNetHTTPHandlerConfig config_dump@net/http.Handler
type ConfigDumpNetHTTPHandlerConfig struct {
	ReadOnly bool
}

func init() {
	_ = provider.Register(
		kindConfigDumpNetHTTPHandlerConfig,
		func(r *ConfigDumpNetHTTPHandlerConfig) HTTPHandler { return r },
	)
}

func (ConfigDumpNetHTTPHandlerConfig) isHTTPHandler() {}
func (ConfigDumpNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ConfigDumpNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t ConfigDumpNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindConfigDumpNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End config_dump@net/http.Handler type =========

// ========= Begin def@io.Reader type =========
//

const kindDefIoReaderConfig = `def@io.Reader`

// DefIoReaderConfig def@io.Reader
type DefIoReaderConfig struct {
	Name string
	Def  IoReader `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefIoReaderConfig,
		func(r *DefIoReaderConfig) IoReader { return r },
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
	data = prepend(kindKey, kindDefIoReaderConfig, data)
	return data, nil
}

//
// ========= End def@io.Reader type =========

// ========= Begin def@io.Writer type =========
//

const kindDefIoWriterConfig = `def@io.Writer`

// DefIoWriterConfig def@io.Writer
type DefIoWriterConfig struct {
	Name string
	Def  IoWriter `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefIoWriterConfig,
		func(r *DefIoWriterConfig) IoWriter { return r },
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
	data = prepend(kindKey, kindDefIoWriterConfig, data)
	return data, nil
}

//
// ========= End def@io.Writer type =========

// ========= Begin def@net.Conn type =========
//

const kindDefNetConnConfig = `def@net.Conn`

// DefNetConnConfig def@net.Conn
type DefNetConnConfig struct {
	Name string
	Def  NetConn `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefNetConnConfig,
		func(r *DefNetConnConfig) NetConn { return r },
	)
}

func (DefNetConnConfig) isNetConn()   {}
func (DefNetConnConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefNetConnConfig) MarshalJSON() ([]byte, error) {
	type t DefNetConnConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindDefNetConnConfig, data)
	return data, nil
}

//
// ========= End def@net.Conn type =========

// ========= Begin def@net.PacketConn type =========
//

const kindDefNetPacketConnConfig = `def@net.PacketConn`

// DefNetPacketConnConfig def@net.PacketConn
type DefNetPacketConnConfig struct {
	Name string
	Def  NetPacketConn `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefNetPacketConnConfig,
		func(r *DefNetPacketConnConfig) NetPacketConn { return r },
	)
}

func (DefNetPacketConnConfig) isNetPacketConn() {}
func (DefNetPacketConnConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefNetPacketConnConfig) MarshalJSON() ([]byte, error) {
	type t DefNetPacketConnConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindDefNetPacketConnConfig, data)
	return data, nil
}

//
// ========= End def@net.PacketConn type =========

// ========= Begin def@net/http.Handler type =========
//

const kindDefNetHTTPHandlerConfig = `def@net/http.Handler`

// DefNetHTTPHandlerConfig def@net/http.Handler
type DefNetHTTPHandlerConfig struct {
	Name string
	Def  HTTPHandler `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefNetHTTPHandlerConfig,
		func(r *DefNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindDefNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End def@net/http.Handler type =========

// ========= Begin def@once.Once type =========
//

const kindDefOnceConfig = `def@once.Once`

// DefOnceConfig def@once.Once
type DefOnceConfig struct {
	Name string
	Def  Once `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefOnceConfig,
		func(r *DefOnceConfig) Once { return r },
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
	data = prepend(kindKey, kindDefOnceConfig, data)
	return data, nil
}

//
// ========= End def@once.Once type =========

// ========= Begin def@packet.Handler type =========
//

const kindDefPacketHandlerConfig = `def@packet.Handler`

// DefPacketHandlerConfig def@packet.Handler
type DefPacketHandlerConfig struct {
	Name string
	Def  PacketHandler `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefPacketHandlerConfig,
		func(r *DefPacketHandlerConfig) PacketHandler { return r },
	)
}

func (DefPacketHandlerConfig) isPacketHandler() {}
func (DefPacketHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefPacketHandlerConfig) MarshalJSON() ([]byte, error) {
	type t DefPacketHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindDefPacketHandlerConfig, data)
	return data, nil
}

//
// ========= End def@packet.Handler type =========

// ========= Begin def@packet.ListenConfig type =========
//

const kindDefPacketListenConfigConfig = `def@packet.ListenConfig`

// DefPacketListenConfigConfig def@packet.ListenConfig
type DefPacketListenConfigConfig struct {
	Name string
	Def  PacketListenConfig `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefPacketListenConfigConfig,
		func(r *DefPacketListenConfigConfig) PacketListenConfig { return r },
	)
}

func (DefPacketListenConfigConfig) isPacketListenConfig() {}
func (DefPacketListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefPacketListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t DefPacketListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindDefPacketListenConfigConfig, data)
	return data, nil
}

//
// ========= End def@packet.ListenConfig type =========

// ========= Begin def@protocol.Handler type =========
//

const kindDefProtocolHandlerConfig = `def@protocol.Handler`

// DefProtocolHandlerConfig def@protocol.Handler
type DefProtocolHandlerConfig struct {
	Name string
	Def  ProtocolHandler `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefProtocolHandlerConfig,
		func(r *DefProtocolHandlerConfig) ProtocolHandler { return r },
	)
}

func (DefProtocolHandlerConfig) isProtocolHandler() {}
func (DefProtocolHandlerConfig) isComponent()       {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefProtocolHandlerConfig) MarshalJSON() ([]byte, error) {
	type t DefProtocolHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindDefProtocolHandlerConfig, data)
	return data, nil
}

//
// ========= End def@protocol.Handler type =========

// ========= Begin def@service.Service type =========
//

const kindDefServiceConfig = `def@service.Service`

// DefServiceConfig def@service.Service
type DefServiceConfig struct {
	Name string
	Def  Service `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefServiceConfig,
		func(r *DefServiceConfig) Service { return r },
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
	data = prepend(kindKey, kindDefServiceConfig, data)
	return data, nil
}

//
// ========= End def@service.Service type =========

// ========= Begin def@stream.Dialer type =========
//

const kindDefStreamDialerConfig = `def@stream.Dialer`

// DefStreamDialerConfig def@stream.Dialer
type DefStreamDialerConfig struct {
	Name string
	Def  StreamDialer `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefStreamDialerConfig,
		func(r *DefStreamDialerConfig) StreamDialer { return r },
	)
}

func (DefStreamDialerConfig) isStreamDialer() {}
func (DefStreamDialerConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t DefStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindDefStreamDialerConfig, data)
	return data, nil
}

//
// ========= End def@stream.Dialer type =========

// ========= Begin def@stream.Handler type =========
//

const kindDefStreamHandlerConfig = `def@stream.Handler`

// DefStreamHandlerConfig def@stream.Handler
type DefStreamHandlerConfig struct {
	Name string
	Def  StreamHandler `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefStreamHandlerConfig,
		func(r *DefStreamHandlerConfig) StreamHandler { return r },
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
	data = prepend(kindKey, kindDefStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End def@stream.Handler type =========

// ========= Begin def@stream.ListenConfig type =========
//

const kindDefStreamListenConfigConfig = `def@stream.ListenConfig`

// DefStreamListenConfigConfig def@stream.ListenConfig
type DefStreamListenConfigConfig struct {
	Name string
	Def  StreamListenConfig `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefStreamListenConfigConfig,
		func(r *DefStreamListenConfigConfig) StreamListenConfig { return r },
	)
}

func (DefStreamListenConfigConfig) isStreamListenConfig() {}
func (DefStreamListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DefStreamListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t DefStreamListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindDefStreamListenConfigConfig, data)
	return data, nil
}

//
// ========= End def@stream.ListenConfig type =========

// ========= Begin def@tls.TLS type =========
//

const kindDefTLSConfig = `def@tls.TLS`

// DefTLSConfig def@tls.TLS
type DefTLSConfig struct {
	Name string
	Def  TLS `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindDefTLSConfig,
		func(r *DefTLSConfig) TLS { return r },
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
	data = prepend(kindKey, kindDefTLSConfig, data)
	return data, nil
}

//
// ========= End def@tls.TLS type =========

// ========= Begin dialer@stream.Dialer type =========
//

const kindDialerStreamDialerConfig = `dialer@stream.Dialer`

// DialerStreamDialerConfig dialer@stream.Dialer
type DialerStreamDialerConfig struct {
	Network  DialerStreamDialerDialerNetworkEnum
	Address  string
	Original bool `json:",omitempty"`
	Virtual  bool `json:",omitempty"`
}

type DialerStreamDialerDialerNetworkEnum string

const (
	DialerStreamDialerDialerNetworkEnumEnumUnix DialerStreamDialerDialerNetworkEnum = "unix"
	DialerStreamDialerDialerNetworkEnumEnumTCP6 DialerStreamDialerDialerNetworkEnum = "tcp6"
	DialerStreamDialerDialerNetworkEnumEnumTCP4 DialerStreamDialerDialerNetworkEnum = "tcp4"
	DialerStreamDialerDialerNetworkEnumEnumTCP  DialerStreamDialerDialerNetworkEnum = "tcp"
)

func init() {
	_ = provider.Register(
		kindDialerStreamDialerConfig,
		func(r *DialerStreamDialerConfig) StreamDialer { return r },
	)
}

func (DialerStreamDialerConfig) isStreamDialer() {}
func (DialerStreamDialerConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DialerStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t DialerStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindDialerStreamDialerConfig, data)
	return data, nil
}

//
// ========= End dialer@stream.Dialer type =========

// ========= Begin direct@net/http.Handler type =========
//

const kindDirectNetHTTPHandlerConfig = `direct@net/http.Handler`

// DirectNetHTTPHandlerConfig direct@net/http.Handler
type DirectNetHTTPHandlerConfig struct {
	Code int
	Body IoReader
}

func init() {
	_ = provider.Register(
		kindDirectNetHTTPHandlerConfig,
		func(r *DirectNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindDirectNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End direct@net/http.Handler type =========

// ========= Begin env@io.Reader type =========
//

const kindEnvIoReaderConfig = `env@io.Reader`

// EnvIoReaderConfig env@io.Reader
type EnvIoReaderConfig struct {
	Name string
}

func init() {
	_ = provider.Register(
		kindEnvIoReaderConfig,
		func(r *EnvIoReaderConfig) IoReader { return r },
	)
}

func (EnvIoReaderConfig) isIoReader()  {}
func (EnvIoReaderConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m EnvIoReaderConfig) MarshalJSON() ([]byte, error) {
	type t EnvIoReaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindEnvIoReaderConfig, data)
	return data, nil
}

//
// ========= End env@io.Reader type =========

// ========= Begin expvar@net/http.Handler type =========
//

const kindExpvarNetHTTPHandler = `expvar@net/http.Handler`

// ExpvarNetHTTPHandler expvar@net/http.Handler
type ExpvarNetHTTPHandler struct {
}

func init() {
	_ = provider.Register(
		kindExpvarNetHTTPHandler,
		func(r *ExpvarNetHTTPHandler) HTTPHandler { return r },
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
	data = prepend(kindKey, kindExpvarNetHTTPHandler, data)
	return data, nil
}

//
// ========= End expvar@net/http.Handler type =========

// ========= Begin file@io.Reader type =========
//

const kindFileIoReaderConfig = `file@io.Reader`

// FileIoReaderConfig file@io.Reader
type FileIoReaderConfig struct {
	Path string
}

func init() {
	_ = provider.Register(
		kindFileIoReaderConfig,
		func(r *FileIoReaderConfig) IoReader { return r },
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
	data = prepend(kindKey, kindFileIoReaderConfig, data)
	return data, nil
}

//
// ========= End file@io.Reader type =========

// ========= Begin file@io.Writer type =========
//

const kindFileIoWriterConfig = `file@io.Writer`

// FileIoWriterConfig file@io.Writer
type FileIoWriterConfig struct {
	Path string
}

func init() {
	_ = provider.Register(
		kindFileIoWriterConfig,
		func(r *FileIoWriterConfig) IoWriter { return r },
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
	data = prepend(kindKey, kindFileIoWriterConfig, data)
	return data, nil
}

//
// ========= End file@io.Writer type =========

// ========= Begin file@net/http.Handler type =========
//

const kindFileNetHTTPHandlerConfig = `file@net/http.Handler`

// FileNetHTTPHandlerConfig file@net/http.Handler
type FileNetHTTPHandlerConfig struct {
	Root string
}

func init() {
	_ = provider.Register(
		kindFileNetHTTPHandlerConfig,
		func(r *FileNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindFileNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End file@net/http.Handler type =========

// ========= Begin forward@net/http.Handler type =========
//

const kindForwardNetHTTPHandlerConfig = `forward@net/http.Handler`

// ForwardNetHTTPHandlerConfig forward@net/http.Handler
type ForwardNetHTTPHandlerConfig struct {
	Dialer StreamDialer `json:",omitempty"`
	URL    string
}

func init() {
	_ = provider.Register(
		kindForwardNetHTTPHandlerConfig,
		func(r *ForwardNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindForwardNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End forward@net/http.Handler type =========

// ========= Begin forward@stream.Handler type =========
//

const kindForwardStreamHandlerConfig = `forward@stream.Handler`

// ForwardStreamHandlerConfig forward@stream.Handler
type ForwardStreamHandlerConfig struct {
	Dialer StreamDialer
}

func init() {
	_ = provider.Register(
		kindForwardStreamHandlerConfig,
		func(r *ForwardStreamHandlerConfig) StreamHandler { return r },
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
	data = prepend(kindKey, kindForwardStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End forward@stream.Handler type =========

// ========= Begin from@tls.TLS type =========
//

const kindFromTLSConfig = `from@tls.TLS`

// FromTLSConfig from@tls.TLS
type FromTLSConfig struct {
	Domain string
	Cert   IoReader
	Key    IoReader
}

func init() {
	_ = provider.Register(
		kindFromTLSConfig,
		func(r *FromTLSConfig) TLS { return r },
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
	data = prepend(kindKey, kindFromTLSConfig, data)
	return data, nil
}

//
// ========= End from@tls.TLS type =========

// ========= Begin hosts@net/http.Handler type =========
//

const kindHostsNetHTTPHandlerConfig = `hosts@net/http.Handler`

// HostsNetHTTPHandlerConfig hosts@net/http.Handler
type HostsNetHTTPHandlerConfig struct {
	Hosts    []HostsNetHTTPHandlerRoute
	NotFound HTTPHandler `json:",omitempty"`
}

type HostsNetHTTPHandlerRoute struct {
	Domains []string
	Handler HTTPHandler
}

func init() {
	_ = provider.Register(
		kindHostsNetHTTPHandlerConfig,
		func(r *HostsNetHTTPHandlerConfig) HTTPHandler { return r },
	)
}

func (HostsNetHTTPHandlerConfig) isHTTPHandler() {}
func (HostsNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HostsNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t HostsNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindHostsNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End hosts@net/http.Handler type =========

// ========= Begin http1@stream.Handler type =========
//

const kindHTTP1StreamHandlerConfig = `http1@stream.Handler`

// HTTP1StreamHandlerConfig http1@stream.Handler
type HTTP1StreamHandlerConfig struct {
	Handler HTTPHandler
}

func init() {
	_ = provider.Register(
		kindHTTP1StreamHandlerConfig,
		func(r *HTTP1StreamHandlerConfig) StreamHandler { return r },
	)
}

func (HTTP1StreamHandlerConfig) isStreamHandler() {}
func (HTTP1StreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTP1StreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t HTTP1StreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindHTTP1StreamHandlerConfig, data)
	return data, nil
}

//
// ========= End http1@stream.Handler type =========

// ========= Begin http2@stream.Handler type =========
//

const kindHTTP2StreamHandlerConfig = `http2@stream.Handler`

// HTTP2StreamHandlerConfig http2@stream.Handler
type HTTP2StreamHandlerConfig struct {
	Handler HTTPHandler
	TLS     TLS `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindHTTP2StreamHandlerConfig,
		func(r *HTTP2StreamHandlerConfig) StreamHandler { return r },
	)
}

func (HTTP2StreamHandlerConfig) isStreamHandler() {}
func (HTTP2StreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTP2StreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t HTTP2StreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindHTTP2StreamHandlerConfig, data)
	return data, nil
}

//
// ========= End http2@stream.Handler type =========

// ========= Begin http3@packet.Handler type =========
//

const kindHTTP3PacketHandlerConfig = `http3@packet.Handler`

// HTTP3PacketHandlerConfig http3@packet.Handler
type HTTP3PacketHandlerConfig struct {
	Handler HTTPHandler
	TLS     TLS
}

func init() {
	_ = provider.Register(
		kindHTTP3PacketHandlerConfig,
		func(r *HTTP3PacketHandlerConfig) PacketHandler { return r },
	)
}

func (HTTP3PacketHandlerConfig) isPacketHandler() {}
func (HTTP3PacketHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTP3PacketHandlerConfig) MarshalJSON() ([]byte, error) {
	type t HTTP3PacketHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindHTTP3PacketHandlerConfig, data)
	return data, nil
}

//
// ========= End http3@packet.Handler type =========

// ========= Begin http@io.Reader type =========
//

const kindHTTPIoReaderConfig = `http@io.Reader`

// HTTPIoReaderConfig http@io.Reader
type HTTPIoReaderConfig struct {
	Dialer StreamDialer `json:",omitempty"`
	URL    string
}

func init() {
	_ = provider.Register(
		kindHTTPIoReaderConfig,
		func(r *HTTPIoReaderConfig) IoReader { return r },
	)
}

func (HTTPIoReaderConfig) isIoReader()  {}
func (HTTPIoReaderConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPIoReaderConfig) MarshalJSON() ([]byte, error) {
	type t HTTPIoReaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindHTTPIoReaderConfig, data)
	return data, nil
}

//
// ========= End http@io.Reader type =========

// ========= Begin inline@io.Reader type =========
//

const kindInlineIoReaderConfig = `inline@io.Reader`

// InlineIoReaderConfig inline@io.Reader
type InlineIoReaderConfig struct {
	Data string
}

func init() {
	_ = provider.Register(
		kindInlineIoReaderConfig,
		func(r *InlineIoReaderConfig) IoReader { return r },
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
	data = prepend(kindKey, kindInlineIoReaderConfig, data)
	return data, nil
}

//
// ========= End inline@io.Reader type =========

// ========= Begin lb@net/http.Handler type =========
//

const kindLbNetHTTPHandlerConfig = `lb@net/http.Handler`

// LbNetHTTPHandlerConfig lb@net/http.Handler
type LbNetHTTPHandlerConfig struct {
	Policy   LbNetHTTPHandlerLoadBalancePolicyEnum `json:",omitempty"`
	Handlers []LbNetHTTPHandlerWeight
}

type LbNetHTTPHandlerLoadBalancePolicyEnum string

const (
	LbNetHTTPHandlerLoadBalancePolicyEnumEnumRoundRobin LbNetHTTPHandlerLoadBalancePolicyEnum = "round_robin"
	LbNetHTTPHandlerLoadBalancePolicyEnumEnumRandom     LbNetHTTPHandlerLoadBalancePolicyEnum = "random"
)

type LbNetHTTPHandlerWeight struct {
	Weight  uint `json:",omitempty"`
	Handler HTTPHandler
}

func init() {
	_ = provider.Register(
		kindLbNetHTTPHandlerConfig,
		func(r *LbNetHTTPHandlerConfig) HTTPHandler { return r },
	)
}

func (LbNetHTTPHandlerConfig) isHTTPHandler() {}
func (LbNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LbNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t LbNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLbNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End lb@net/http.Handler type =========

// ========= Begin lb@stream.Dialer type =========
//

const kindLbStreamDialerConfig = `lb@stream.Dialer`

// LbStreamDialerConfig lb@stream.Dialer
type LbStreamDialerConfig struct {
	Policy  LbStreamDialerLoadBalancePolicyEnum `json:",omitempty"`
	Dialers []LbStreamDialerWeight
}

type LbStreamDialerLoadBalancePolicyEnum string

const (
	LbStreamDialerLoadBalancePolicyEnumEnumRoundRobin LbStreamDialerLoadBalancePolicyEnum = "round_robin"
	LbStreamDialerLoadBalancePolicyEnumEnumRandom     LbStreamDialerLoadBalancePolicyEnum = "random"
)

type LbStreamDialerWeight struct {
	Weight uint `json:",omitempty"`
	Dialer StreamDialer
}

func init() {
	_ = provider.Register(
		kindLbStreamDialerConfig,
		func(r *LbStreamDialerConfig) StreamDialer { return r },
	)
}

func (LbStreamDialerConfig) isStreamDialer() {}
func (LbStreamDialerConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LbStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t LbStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLbStreamDialerConfig, data)
	return data, nil
}

//
// ========= End lb@stream.Dialer type =========

// ========= Begin lb@stream.Handler type =========
//

const kindLbStreamHandlerConfig = `lb@stream.Handler`

// LbStreamHandlerConfig lb@stream.Handler
type LbStreamHandlerConfig struct {
	Policy   LbStreamHandlerLoadBalancePolicyEnum `json:",omitempty"`
	Handlers []LbStreamHandlerWeight
}

type LbStreamHandlerLoadBalancePolicyEnum string

const (
	LbStreamHandlerLoadBalancePolicyEnumEnumRoundRobin LbStreamHandlerLoadBalancePolicyEnum = "round_robin"
	LbStreamHandlerLoadBalancePolicyEnumEnumRandom     LbStreamHandlerLoadBalancePolicyEnum = "random"
)

type LbStreamHandlerWeight struct {
	Weight  uint `json:",omitempty"`
	Handler StreamHandler
}

func init() {
	_ = provider.Register(
		kindLbStreamHandlerConfig,
		func(r *LbStreamHandlerConfig) StreamHandler { return r },
	)
}

func (LbStreamHandlerConfig) isStreamHandler() {}
func (LbStreamHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LbStreamHandlerConfig) MarshalJSON() ([]byte, error) {
	type t LbStreamHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLbStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End lb@stream.Handler type =========

// ========= Begin listener@packet.ListenConfig type =========
//

const kindListenerPacketListenConfigConfig = `listener@packet.ListenConfig`

// ListenerPacketListenConfigConfig listener@packet.ListenConfig
type ListenerPacketListenConfigConfig struct {
	Network ListenerPacketListenConfigListenerNetworkEnum
	Address string
}

type ListenerPacketListenConfigListenerNetworkEnum string

const (
	ListenerPacketListenConfigListenerNetworkEnumEnumUnixPacket ListenerPacketListenConfigListenerNetworkEnum = "unixpacket"
	ListenerPacketListenConfigListenerNetworkEnumEnumUDP6       ListenerPacketListenConfigListenerNetworkEnum = "udp6"
	ListenerPacketListenConfigListenerNetworkEnumEnumUDP4       ListenerPacketListenConfigListenerNetworkEnum = "udp4"
	ListenerPacketListenConfigListenerNetworkEnumEnumUDP        ListenerPacketListenConfigListenerNetworkEnum = "udp"
)

func init() {
	_ = provider.Register(
		kindListenerPacketListenConfigConfig,
		func(r *ListenerPacketListenConfigConfig) PacketListenConfig { return r },
	)
}

func (ListenerPacketListenConfigConfig) isPacketListenConfig() {}
func (ListenerPacketListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ListenerPacketListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t ListenerPacketListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindListenerPacketListenConfigConfig, data)
	return data, nil
}

//
// ========= End listener@packet.ListenConfig type =========

// ========= Begin listener@stream.ListenConfig type =========
//

const kindListenerStreamListenConfigConfig = `listener@stream.ListenConfig`

// ListenerStreamListenConfigConfig listener@stream.ListenConfig
type ListenerStreamListenConfigConfig struct {
	Network ListenerStreamListenConfigListenerNetworkEnum
	Address string
	Virtual bool `json:",omitempty"`
}

type ListenerStreamListenConfigListenerNetworkEnum string

const (
	ListenerStreamListenConfigListenerNetworkEnumEnumUnix ListenerStreamListenConfigListenerNetworkEnum = "unix"
	ListenerStreamListenConfigListenerNetworkEnumEnumTCP6 ListenerStreamListenConfigListenerNetworkEnum = "tcp6"
	ListenerStreamListenConfigListenerNetworkEnumEnumTCP4 ListenerStreamListenConfigListenerNetworkEnum = "tcp4"
	ListenerStreamListenConfigListenerNetworkEnumEnumTCP  ListenerStreamListenConfigListenerNetworkEnum = "tcp"
)

func init() {
	_ = provider.Register(
		kindListenerStreamListenConfigConfig,
		func(r *ListenerStreamListenConfigConfig) StreamListenConfig { return r },
	)
}

func (ListenerStreamListenConfigConfig) isStreamListenConfig() {}
func (ListenerStreamListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ListenerStreamListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t ListenerStreamListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindListenerStreamListenConfigConfig, data)
	return data, nil
}

//
// ========= End listener@stream.ListenConfig type =========

// ========= Begin load@io.Reader type =========
//

const kindLoadIoReaderConfig = `load@io.Reader`

// LoadIoReaderConfig load@io.Reader
type LoadIoReaderConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadIoReaderConfig,
		func(r *LoadIoReaderConfig) IoReader { return r },
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
	data = prepend(kindKey, kindLoadIoReaderConfig, data)
	return data, nil
}

//
// ========= End load@io.Reader type =========

// ========= Begin load@io.Writer type =========
//

const kindLoadIoWriterConfig = `load@io.Writer`

// LoadIoWriterConfig load@io.Writer
type LoadIoWriterConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadIoWriterConfig,
		func(r *LoadIoWriterConfig) IoWriter { return r },
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
	data = prepend(kindKey, kindLoadIoWriterConfig, data)
	return data, nil
}

//
// ========= End load@io.Writer type =========

// ========= Begin load@net.Conn type =========
//

const kindLoadNetConnConfig = `load@net.Conn`

// LoadNetConnConfig load@net.Conn
type LoadNetConnConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadNetConnConfig,
		func(r *LoadNetConnConfig) NetConn { return r },
	)
}

func (LoadNetConnConfig) isNetConn()   {}
func (LoadNetConnConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadNetConnConfig) MarshalJSON() ([]byte, error) {
	type t LoadNetConnConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLoadNetConnConfig, data)
	return data, nil
}

//
// ========= End load@net.Conn type =========

// ========= Begin load@net.PacketConn type =========
//

const kindLoadNetPacketConnConfig = `load@net.PacketConn`

// LoadNetPacketConnConfig load@net.PacketConn
type LoadNetPacketConnConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadNetPacketConnConfig,
		func(r *LoadNetPacketConnConfig) NetPacketConn { return r },
	)
}

func (LoadNetPacketConnConfig) isNetPacketConn() {}
func (LoadNetPacketConnConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadNetPacketConnConfig) MarshalJSON() ([]byte, error) {
	type t LoadNetPacketConnConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLoadNetPacketConnConfig, data)
	return data, nil
}

//
// ========= End load@net.PacketConn type =========

// ========= Begin load@net/http.Handler type =========
//

const kindLoadNetHTTPHandlerConfig = `load@net/http.Handler`

// LoadNetHTTPHandlerConfig load@net/http.Handler
type LoadNetHTTPHandlerConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadNetHTTPHandlerConfig,
		func(r *LoadNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindLoadNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End load@net/http.Handler type =========

// ========= Begin load@once.Once type =========
//

const kindLoadOnceConfig = `load@once.Once`

// LoadOnceConfig load@once.Once
type LoadOnceConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadOnceConfig,
		func(r *LoadOnceConfig) Once { return r },
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
	data = prepend(kindKey, kindLoadOnceConfig, data)
	return data, nil
}

//
// ========= End load@once.Once type =========

// ========= Begin load@packet.Handler type =========
//

const kindLoadPacketHandlerConfig = `load@packet.Handler`

// LoadPacketHandlerConfig load@packet.Handler
type LoadPacketHandlerConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadPacketHandlerConfig,
		func(r *LoadPacketHandlerConfig) PacketHandler { return r },
	)
}

func (LoadPacketHandlerConfig) isPacketHandler() {}
func (LoadPacketHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadPacketHandlerConfig) MarshalJSON() ([]byte, error) {
	type t LoadPacketHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLoadPacketHandlerConfig, data)
	return data, nil
}

//
// ========= End load@packet.Handler type =========

// ========= Begin load@packet.ListenConfig type =========
//

const kindLoadPacketListenConfigConfig = `load@packet.ListenConfig`

// LoadPacketListenConfigConfig load@packet.ListenConfig
type LoadPacketListenConfigConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadPacketListenConfigConfig,
		func(r *LoadPacketListenConfigConfig) PacketListenConfig { return r },
	)
}

func (LoadPacketListenConfigConfig) isPacketListenConfig() {}
func (LoadPacketListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadPacketListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t LoadPacketListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLoadPacketListenConfigConfig, data)
	return data, nil
}

//
// ========= End load@packet.ListenConfig type =========

// ========= Begin load@protocol.Handler type =========
//

const kindLoadProtocolHandlerConfig = `load@protocol.Handler`

// LoadProtocolHandlerConfig load@protocol.Handler
type LoadProtocolHandlerConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadProtocolHandlerConfig,
		func(r *LoadProtocolHandlerConfig) ProtocolHandler { return r },
	)
}

func (LoadProtocolHandlerConfig) isProtocolHandler() {}
func (LoadProtocolHandlerConfig) isComponent()       {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadProtocolHandlerConfig) MarshalJSON() ([]byte, error) {
	type t LoadProtocolHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLoadProtocolHandlerConfig, data)
	return data, nil
}

//
// ========= End load@protocol.Handler type =========

// ========= Begin load@service.Service type =========
//

const kindLoadServiceConfig = `load@service.Service`

// LoadServiceConfig load@service.Service
type LoadServiceConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadServiceConfig,
		func(r *LoadServiceConfig) Service { return r },
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
	data = prepend(kindKey, kindLoadServiceConfig, data)
	return data, nil
}

//
// ========= End load@service.Service type =========

// ========= Begin load@stream.Dialer type =========
//

const kindLoadStreamDialerConfig = `load@stream.Dialer`

// LoadStreamDialerConfig load@stream.Dialer
type LoadStreamDialerConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadStreamDialerConfig,
		func(r *LoadStreamDialerConfig) StreamDialer { return r },
	)
}

func (LoadStreamDialerConfig) isStreamDialer() {}
func (LoadStreamDialerConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t LoadStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLoadStreamDialerConfig, data)
	return data, nil
}

//
// ========= End load@stream.Dialer type =========

// ========= Begin load@stream.Handler type =========
//

const kindLoadStreamHandlerConfig = `load@stream.Handler`

// LoadStreamHandlerConfig load@stream.Handler
type LoadStreamHandlerConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadStreamHandlerConfig,
		func(r *LoadStreamHandlerConfig) StreamHandler { return r },
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
	data = prepend(kindKey, kindLoadStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End load@stream.Handler type =========

// ========= Begin load@stream.ListenConfig type =========
//

const kindLoadStreamListenConfigConfig = `load@stream.ListenConfig`

// LoadStreamListenConfigConfig load@stream.ListenConfig
type LoadStreamListenConfigConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadStreamListenConfigConfig,
		func(r *LoadStreamListenConfigConfig) StreamListenConfig { return r },
	)
}

func (LoadStreamListenConfigConfig) isStreamListenConfig() {}
func (LoadStreamListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m LoadStreamListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t LoadStreamListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindLoadStreamListenConfigConfig, data)
	return data, nil
}

//
// ========= End load@stream.ListenConfig type =========

// ========= Begin load@tls.TLS type =========
//

const kindLoadTLSConfig = `load@tls.TLS`

// LoadTLSConfig load@tls.TLS
type LoadTLSConfig struct {
	Load IoReader
}

func init() {
	_ = provider.Register(
		kindLoadTLSConfig,
		func(r *LoadTLSConfig) TLS { return r },
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
	data = prepend(kindKey, kindLoadTLSConfig, data)
	return data, nil
}

//
// ========= End load@tls.TLS type =========

// ========= Begin log@net/http.Handler type =========
//

const kindLogNetHTTPHandlerConfig = `log@net/http.Handler`

// LogNetHTTPHandlerConfig log@net/http.Handler
type LogNetHTTPHandlerConfig struct {
	Output  IoWriter
	Handler HTTPHandler
}

func init() {
	_ = provider.Register(
		kindLogNetHTTPHandlerConfig,
		func(r *LogNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindLogNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End log@net/http.Handler type =========

// ========= Begin merge@tls.TLS type =========
//

const kindMergeTLSConfig = `merge@tls.TLS`

// MergeTLSConfig merge@tls.TLS
type MergeTLSConfig struct {
	Merge []TLS
}

func init() {
	_ = provider.Register(
		kindMergeTLSConfig,
		func(r *MergeTLSConfig) TLS { return r },
	)
}

func (MergeTLSConfig) isTLS()       {}
func (MergeTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MergeTLSConfig) MarshalJSON() ([]byte, error) {
	type t MergeTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindMergeTLSConfig, data)
	return data, nil
}

//
// ========= End merge@tls.TLS type =========

// ========= Begin message@once.Once type =========
//

const kindMessageOnceConfig = `message@once.Once`

// MessageOnceConfig message@once.Once
type MessageOnceConfig struct {
	Message string
}

func init() {
	_ = provider.Register(
		kindMessageOnceConfig,
		func(r *MessageOnceConfig) Once { return r },
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
	data = prepend(kindKey, kindMessageOnceConfig, data)
	return data, nil
}

//
// ========= End message@once.Once type =========

// ========= Begin method@net/http.Handler type =========
//

const kindMethodNetHTTPHandlerConfig = `method@net/http.Handler`

// MethodNetHTTPHandlerConfig method@net/http.Handler
type MethodNetHTTPHandlerConfig struct {
	Methods  []MethodNetHTTPHandlerRoute
	NotFound HTTPHandler
}

type MethodNetHTTPHandlerRoute struct {
	Method  MethodNetHTTPHandlerMethodEnum
	Handler HTTPHandler
}

type MethodNetHTTPHandlerMethodEnum string

const (
	MethodNetHTTPHandlerMethodEnumMethodTrace   MethodNetHTTPHandlerMethodEnum = "TRACE"
	MethodNetHTTPHandlerMethodEnumMethodPut     MethodNetHTTPHandlerMethodEnum = "PUT"
	MethodNetHTTPHandlerMethodEnumMethodPost    MethodNetHTTPHandlerMethodEnum = "POST"
	MethodNetHTTPHandlerMethodEnumMethodPatch   MethodNetHTTPHandlerMethodEnum = "PATCH"
	MethodNetHTTPHandlerMethodEnumMethodOptions MethodNetHTTPHandlerMethodEnum = "OPTIONS"
	MethodNetHTTPHandlerMethodEnumMethodHead    MethodNetHTTPHandlerMethodEnum = "HEAD"
	MethodNetHTTPHandlerMethodEnumMethodGet     MethodNetHTTPHandlerMethodEnum = "GET"
	MethodNetHTTPHandlerMethodEnumMethodDelete  MethodNetHTTPHandlerMethodEnum = "DELETE"
	MethodNetHTTPHandlerMethodEnumMethodConnect MethodNetHTTPHandlerMethodEnum = "CONNECT"
)

func init() {
	_ = provider.Register(
		kindMethodNetHTTPHandlerConfig,
		func(r *MethodNetHTTPHandlerConfig) HTTPHandler { return r },
	)
}

func (MethodNetHTTPHandlerConfig) isHTTPHandler() {}
func (MethodNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MethodNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t MethodNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindMethodNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End method@net/http.Handler type =========

// ========= Begin metrics@net/http.Handler type =========
//

const kindMetricsNetHTTPHandler = `metrics@net/http.Handler`

// MetricsNetHTTPHandler metrics@net/http.Handler
type MetricsNetHTTPHandler struct {
}

func init() {
	_ = provider.Register(
		kindMetricsNetHTTPHandler,
		func(r *MetricsNetHTTPHandler) HTTPHandler { return r },
	)
}

func (MetricsNetHTTPHandler) isHTTPHandler() {}
func (MetricsNetHTTPHandler) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m MetricsNetHTTPHandler) MarshalJSON() ([]byte, error) {
	type t MetricsNetHTTPHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindMetricsNetHTTPHandler, data)
	return data, nil
}

//
// ========= End metrics@net/http.Handler type =========

// ========= Begin multi@net/http.Handler type =========
//

const kindMultiNetHTTPHandlerConfig = `multi@net/http.Handler`

// MultiNetHTTPHandlerConfig multi@net/http.Handler
type MultiNetHTTPHandlerConfig struct {
	Multi []HTTPHandler
}

func init() {
	_ = provider.Register(
		kindMultiNetHTTPHandlerConfig,
		func(r *MultiNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindMultiNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End multi@net/http.Handler type =========

// ========= Begin multi@once.Once type =========
//

const kindMultiOnceConfig = `multi@once.Once`

// MultiOnceConfig multi@once.Once
type MultiOnceConfig struct {
	Multi []Once
}

func init() {
	_ = provider.Register(
		kindMultiOnceConfig,
		func(r *MultiOnceConfig) Once { return r },
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
	data = prepend(kindKey, kindMultiOnceConfig, data)
	return data, nil
}

//
// ========= End multi@once.Once type =========

// ========= Begin multi@service.Service type =========
//

const kindMultiServiceConfig = `multi@service.Service`

// MultiServiceConfig multi@service.Service
type MultiServiceConfig struct {
	Multi []Service
}

func init() {
	_ = provider.Register(
		kindMultiServiceConfig,
		func(r *MultiServiceConfig) Service { return r },
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
	data = prepend(kindKey, kindMultiServiceConfig, data)
	return data, nil
}

//
// ========= End multi@service.Service type =========

// ========= Begin multi@stream.Handler type =========
//

const kindMultiStreamHandlerConfig = `multi@stream.Handler`

// MultiStreamHandlerConfig multi@stream.Handler
type MultiStreamHandlerConfig struct {
	Multi []StreamHandler
}

func init() {
	_ = provider.Register(
		kindMultiStreamHandlerConfig,
		func(r *MultiStreamHandlerConfig) StreamHandler { return r },
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
	data = prepend(kindKey, kindMultiStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End multi@stream.Handler type =========

// ========= Begin mux@net/http.Handler type =========
//

const kindMuxNetHTTPHandlerConfig = `mux@net/http.Handler`

// MuxNetHTTPHandlerConfig mux@net/http.Handler
type MuxNetHTTPHandlerConfig struct {
	Routes   []MuxNetHTTPHandlerRoute
	NotFound HTTPHandler `json:",omitempty"`
}

type MuxNetHTTPHandlerRoute struct {
	Prefix  string `json:",omitempty"`
	Path    string `json:",omitempty"`
	Regexp  string `json:",omitempty"`
	Handler HTTPHandler
}

func init() {
	_ = provider.Register(
		kindMuxNetHTTPHandlerConfig,
		func(r *MuxNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindMuxNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End mux@net/http.Handler type =========

// ========= Begin mux@stream.Handler type =========
//

const kindMuxStreamHandlerConfig = `mux@stream.Handler`

// MuxStreamHandlerConfig mux@stream.Handler
type MuxStreamHandlerConfig struct {
	Routes   []MuxStreamHandlerRoute
	NotFound StreamHandler
}

type MuxStreamHandlerRoute struct {
	Pattern string `json:",omitempty"`
	Regexp  string `json:",omitempty"`
	Prefix  string `json:",omitempty"`
	Handler StreamHandler
}

func init() {
	_ = provider.Register(
		kindMuxStreamHandlerConfig,
		func(r *MuxStreamHandlerConfig) StreamHandler { return r },
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
	data = prepend(kindKey, kindMuxStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End mux@stream.Handler type =========

// ========= Begin none@io.Reader type =========
//

const kindNoneIoReader = `none@io.Reader`

// NoneIoReader none@io.Reader
type NoneIoReader struct {
}

func init() {
	_ = provider.Register(
		kindNoneIoReader,
		func(r *NoneIoReader) IoReader { return r },
	)
}

func (NoneIoReader) isIoReader()  {}
func (NoneIoReader) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneIoReader) MarshalJSON() ([]byte, error) {
	type t NoneIoReader
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneIoReader, data)
	return data, nil
}

//
// ========= End none@io.Reader type =========

// ========= Begin none@io.Writer type =========
//

const kindNoneIoWriter = `none@io.Writer`

// NoneIoWriter none@io.Writer
type NoneIoWriter struct {
}

func init() {
	_ = provider.Register(
		kindNoneIoWriter,
		func(r *NoneIoWriter) IoWriter { return r },
	)
}

func (NoneIoWriter) isIoWriter()  {}
func (NoneIoWriter) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneIoWriter) MarshalJSON() ([]byte, error) {
	type t NoneIoWriter
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneIoWriter, data)
	return data, nil
}

//
// ========= End none@io.Writer type =========

// ========= Begin none@net.Conn type =========
//

const kindNoneNetConn = `none@net.Conn`

// NoneNetConn none@net.Conn
type NoneNetConn struct {
}

func init() {
	_ = provider.Register(
		kindNoneNetConn,
		func(r *NoneNetConn) NetConn { return r },
	)
}

func (NoneNetConn) isNetConn()   {}
func (NoneNetConn) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneNetConn) MarshalJSON() ([]byte, error) {
	type t NoneNetConn
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneNetConn, data)
	return data, nil
}

//
// ========= End none@net.Conn type =========

// ========= Begin none@net.PacketConn type =========
//

const kindNoneNetPacketConn = `none@net.PacketConn`

// NoneNetPacketConn none@net.PacketConn
type NoneNetPacketConn struct {
}

func init() {
	_ = provider.Register(
		kindNoneNetPacketConn,
		func(r *NoneNetPacketConn) NetPacketConn { return r },
	)
}

func (NoneNetPacketConn) isNetPacketConn() {}
func (NoneNetPacketConn) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneNetPacketConn) MarshalJSON() ([]byte, error) {
	type t NoneNetPacketConn
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneNetPacketConn, data)
	return data, nil
}

//
// ========= End none@net.PacketConn type =========

// ========= Begin none@net/http.Handler type =========
//

const kindNoneNetHTTPHandler = `none@net/http.Handler`

// NoneNetHTTPHandler none@net/http.Handler
type NoneNetHTTPHandler struct {
}

func init() {
	_ = provider.Register(
		kindNoneNetHTTPHandler,
		func(r *NoneNetHTTPHandler) HTTPHandler { return r },
	)
}

func (NoneNetHTTPHandler) isHTTPHandler() {}
func (NoneNetHTTPHandler) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneNetHTTPHandler) MarshalJSON() ([]byte, error) {
	type t NoneNetHTTPHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneNetHTTPHandler, data)
	return data, nil
}

//
// ========= End none@net/http.Handler type =========

// ========= Begin none@once.Once type =========
//

const kindNoneOnce = `none@once.Once`

// NoneOnce none@once.Once
type NoneOnce struct {
}

func init() {
	_ = provider.Register(
		kindNoneOnce,
		func(r *NoneOnce) Once { return r },
	)
}

func (NoneOnce) isOnce()      {}
func (NoneOnce) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneOnce) MarshalJSON() ([]byte, error) {
	type t NoneOnce
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneOnce, data)
	return data, nil
}

//
// ========= End none@once.Once type =========

// ========= Begin none@packet.Handler type =========
//

const kindNonePacketHandler = `none@packet.Handler`

// NonePacketHandler none@packet.Handler
type NonePacketHandler struct {
}

func init() {
	_ = provider.Register(
		kindNonePacketHandler,
		func(r *NonePacketHandler) PacketHandler { return r },
	)
}

func (NonePacketHandler) isPacketHandler() {}
func (NonePacketHandler) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NonePacketHandler) MarshalJSON() ([]byte, error) {
	type t NonePacketHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNonePacketHandler, data)
	return data, nil
}

//
// ========= End none@packet.Handler type =========

// ========= Begin none@packet.ListenConfig type =========
//

const kindNonePacketListenConfig = `none@packet.ListenConfig`

// NonePacketListenConfig none@packet.ListenConfig
type NonePacketListenConfig struct {
}

func init() {
	_ = provider.Register(
		kindNonePacketListenConfig,
		func(r *NonePacketListenConfig) PacketListenConfig { return r },
	)
}

func (NonePacketListenConfig) isPacketListenConfig() {}
func (NonePacketListenConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NonePacketListenConfig) MarshalJSON() ([]byte, error) {
	type t NonePacketListenConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNonePacketListenConfig, data)
	return data, nil
}

//
// ========= End none@packet.ListenConfig type =========

// ========= Begin none@protocol.Handler type =========
//

const kindNoneProtocolHandler = `none@protocol.Handler`

// NoneProtocolHandler none@protocol.Handler
type NoneProtocolHandler struct {
}

func init() {
	_ = provider.Register(
		kindNoneProtocolHandler,
		func(r *NoneProtocolHandler) ProtocolHandler { return r },
	)
}

func (NoneProtocolHandler) isProtocolHandler() {}
func (NoneProtocolHandler) isComponent()       {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneProtocolHandler) MarshalJSON() ([]byte, error) {
	type t NoneProtocolHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneProtocolHandler, data)
	return data, nil
}

//
// ========= End none@protocol.Handler type =========

// ========= Begin none@service.Service type =========
//

const kindNoneService = `none@service.Service`

// NoneService none@service.Service
type NoneService struct {
}

func init() {
	_ = provider.Register(
		kindNoneService,
		func(r *NoneService) Service { return r },
	)
}

func (NoneService) isService()   {}
func (NoneService) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneService) MarshalJSON() ([]byte, error) {
	type t NoneService
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneService, data)
	return data, nil
}

//
// ========= End none@service.Service type =========

// ========= Begin none@stream.Dialer type =========
//

const kindNoneStreamDialer = `none@stream.Dialer`

// NoneStreamDialer none@stream.Dialer
type NoneStreamDialer struct {
}

func init() {
	_ = provider.Register(
		kindNoneStreamDialer,
		func(r *NoneStreamDialer) StreamDialer { return r },
	)
}

func (NoneStreamDialer) isStreamDialer() {}
func (NoneStreamDialer) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneStreamDialer) MarshalJSON() ([]byte, error) {
	type t NoneStreamDialer
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneStreamDialer, data)
	return data, nil
}

//
// ========= End none@stream.Dialer type =========

// ========= Begin none@stream.Handler type =========
//

const kindNoneStreamHandler = `none@stream.Handler`

// NoneStreamHandler none@stream.Handler
type NoneStreamHandler struct {
}

func init() {
	_ = provider.Register(
		kindNoneStreamHandler,
		func(r *NoneStreamHandler) StreamHandler { return r },
	)
}

func (NoneStreamHandler) isStreamHandler() {}
func (NoneStreamHandler) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneStreamHandler) MarshalJSON() ([]byte, error) {
	type t NoneStreamHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneStreamHandler, data)
	return data, nil
}

//
// ========= End none@stream.Handler type =========

// ========= Begin none@stream.ListenConfig type =========
//

const kindNoneStreamListenConfig = `none@stream.ListenConfig`

// NoneStreamListenConfig none@stream.ListenConfig
type NoneStreamListenConfig struct {
}

func init() {
	_ = provider.Register(
		kindNoneStreamListenConfig,
		func(r *NoneStreamListenConfig) StreamListenConfig { return r },
	)
}

func (NoneStreamListenConfig) isStreamListenConfig() {}
func (NoneStreamListenConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneStreamListenConfig) MarshalJSON() ([]byte, error) {
	type t NoneStreamListenConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneStreamListenConfig, data)
	return data, nil
}

//
// ========= End none@stream.ListenConfig type =========

// ========= Begin none@tls.TLS type =========
//

const kindNoneTLS = `none@tls.TLS`

// NoneTLS none@tls.TLS
type NoneTLS struct {
}

func init() {
	_ = provider.Register(
		kindNoneTLS,
		func(r *NoneTLS) TLS { return r },
	)
}

func (NoneTLS) isTLS()       {}
func (NoneTLS) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m NoneTLS) MarshalJSON() ([]byte, error) {
	type t NoneTLS
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindNoneTLS, data)
	return data, nil
}

//
// ========= End none@tls.TLS type =========

// ========= Begin packet@service.Service type =========
//

const kindPacketServiceConfig = `packet@service.Service`

// PacketServiceConfig packet@service.Service
type PacketServiceConfig struct {
	Listener PacketListenConfig
	Handler  PacketHandler
}

func init() {
	_ = provider.Register(
		kindPacketServiceConfig,
		func(r *PacketServiceConfig) Service { return r },
	)
}

func (PacketServiceConfig) isService()   {}
func (PacketServiceConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m PacketServiceConfig) MarshalJSON() ([]byte, error) {
	type t PacketServiceConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindPacketServiceConfig, data)
	return data, nil
}

//
// ========= End packet@service.Service type =========

// ========= Begin pprof@net/http.Handler type =========
//

const kindPprofNetHTTPHandler = `pprof@net/http.Handler`

// PprofNetHTTPHandler pprof@net/http.Handler
type PprofNetHTTPHandler struct {
}

func init() {
	_ = provider.Register(
		kindPprofNetHTTPHandler,
		func(r *PprofNetHTTPHandler) HTTPHandler { return r },
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
	data = prepend(kindKey, kindPprofNetHTTPHandler, data)
	return data, nil
}

//
// ========= End pprof@net/http.Handler type =========

// ========= Begin quic@stream.ListenConfig type =========
//

const kindQuicStreamListenConfigConfig = `quic@stream.ListenConfig`

// QuicStreamListenConfigConfig quic@stream.ListenConfig
type QuicStreamListenConfigConfig struct {
	Packet NetPacketConn
	TLS    TLS
}

func init() {
	_ = provider.Register(
		kindQuicStreamListenConfigConfig,
		func(r *QuicStreamListenConfigConfig) StreamListenConfig { return r },
	)
}

func (QuicStreamListenConfigConfig) isStreamListenConfig() {}
func (QuicStreamListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m QuicStreamListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t QuicStreamListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindQuicStreamListenConfigConfig, data)
	return data, nil
}

//
// ========= End quic@stream.ListenConfig type =========

// ========= Begin quit@net/http.Handler type =========
//

const kindQuitNetHTTPHandler = `quit@net/http.Handler`

// QuitNetHTTPHandler quit@net/http.Handler
type QuitNetHTTPHandler struct {
}

func init() {
	_ = provider.Register(
		kindQuitNetHTTPHandler,
		func(r *QuitNetHTTPHandler) HTTPHandler { return r },
	)
}

func (QuitNetHTTPHandler) isHTTPHandler() {}
func (QuitNetHTTPHandler) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m QuitNetHTTPHandler) MarshalJSON() ([]byte, error) {
	type t QuitNetHTTPHandler
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindQuitNetHTTPHandler, data)
	return data, nil
}

//
// ========= End quit@net/http.Handler type =========

// ========= Begin redirect@net/http.Handler type =========
//

const kindRedirectNetHTTPHandlerConfig = `redirect@net/http.Handler`

// RedirectNetHTTPHandlerConfig redirect@net/http.Handler
type RedirectNetHTTPHandlerConfig struct {
	Code     int
	Location string
}

func init() {
	_ = provider.Register(
		kindRedirectNetHTTPHandlerConfig,
		func(r *RedirectNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindRedirectNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End redirect@net/http.Handler type =========

// ========= Begin ref@io.Reader type =========
//

const kindRefIoReaderConfig = `ref@io.Reader`

// RefIoReaderConfig ref@io.Reader
type RefIoReaderConfig struct {
	Name string
	Def  IoReader `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefIoReaderConfig,
		func(r *RefIoReaderConfig) IoReader { return r },
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
	data = prepend(kindKey, kindRefIoReaderConfig, data)
	return data, nil
}

//
// ========= End ref@io.Reader type =========

// ========= Begin ref@io.Writer type =========
//

const kindRefIoWriterConfig = `ref@io.Writer`

// RefIoWriterConfig ref@io.Writer
type RefIoWriterConfig struct {
	Name string
	Def  IoWriter `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefIoWriterConfig,
		func(r *RefIoWriterConfig) IoWriter { return r },
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
	data = prepend(kindKey, kindRefIoWriterConfig, data)
	return data, nil
}

//
// ========= End ref@io.Writer type =========

// ========= Begin ref@net.Conn type =========
//

const kindRefNetConnConfig = `ref@net.Conn`

// RefNetConnConfig ref@net.Conn
type RefNetConnConfig struct {
	Name string
	Def  NetConn `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefNetConnConfig,
		func(r *RefNetConnConfig) NetConn { return r },
	)
}

func (RefNetConnConfig) isNetConn()   {}
func (RefNetConnConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefNetConnConfig) MarshalJSON() ([]byte, error) {
	type t RefNetConnConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindRefNetConnConfig, data)
	return data, nil
}

//
// ========= End ref@net.Conn type =========

// ========= Begin ref@net.PacketConn type =========
//

const kindRefNetPacketConnConfig = `ref@net.PacketConn`

// RefNetPacketConnConfig ref@net.PacketConn
type RefNetPacketConnConfig struct {
	Name string
	Def  NetPacketConn `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefNetPacketConnConfig,
		func(r *RefNetPacketConnConfig) NetPacketConn { return r },
	)
}

func (RefNetPacketConnConfig) isNetPacketConn() {}
func (RefNetPacketConnConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefNetPacketConnConfig) MarshalJSON() ([]byte, error) {
	type t RefNetPacketConnConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindRefNetPacketConnConfig, data)
	return data, nil
}

//
// ========= End ref@net.PacketConn type =========

// ========= Begin ref@net/http.Handler type =========
//

const kindRefNetHTTPHandlerConfig = `ref@net/http.Handler`

// RefNetHTTPHandlerConfig ref@net/http.Handler
type RefNetHTTPHandlerConfig struct {
	Name string
	Def  HTTPHandler `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefNetHTTPHandlerConfig,
		func(r *RefNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindRefNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End ref@net/http.Handler type =========

// ========= Begin ref@once.Once type =========
//

const kindRefOnceConfig = `ref@once.Once`

// RefOnceConfig ref@once.Once
type RefOnceConfig struct {
	Name string
	Def  Once `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefOnceConfig,
		func(r *RefOnceConfig) Once { return r },
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
	data = prepend(kindKey, kindRefOnceConfig, data)
	return data, nil
}

//
// ========= End ref@once.Once type =========

// ========= Begin ref@packet.Handler type =========
//

const kindRefPacketHandlerConfig = `ref@packet.Handler`

// RefPacketHandlerConfig ref@packet.Handler
type RefPacketHandlerConfig struct {
	Name string
	Def  PacketHandler `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefPacketHandlerConfig,
		func(r *RefPacketHandlerConfig) PacketHandler { return r },
	)
}

func (RefPacketHandlerConfig) isPacketHandler() {}
func (RefPacketHandlerConfig) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefPacketHandlerConfig) MarshalJSON() ([]byte, error) {
	type t RefPacketHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindRefPacketHandlerConfig, data)
	return data, nil
}

//
// ========= End ref@packet.Handler type =========

// ========= Begin ref@packet.ListenConfig type =========
//

const kindRefPacketListenConfigConfig = `ref@packet.ListenConfig`

// RefPacketListenConfigConfig ref@packet.ListenConfig
type RefPacketListenConfigConfig struct {
	Name string
	Def  PacketListenConfig `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefPacketListenConfigConfig,
		func(r *RefPacketListenConfigConfig) PacketListenConfig { return r },
	)
}

func (RefPacketListenConfigConfig) isPacketListenConfig() {}
func (RefPacketListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefPacketListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t RefPacketListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindRefPacketListenConfigConfig, data)
	return data, nil
}

//
// ========= End ref@packet.ListenConfig type =========

// ========= Begin ref@protocol.Handler type =========
//

const kindRefProtocolHandlerConfig = `ref@protocol.Handler`

// RefProtocolHandlerConfig ref@protocol.Handler
type RefProtocolHandlerConfig struct {
	Name string
	Def  ProtocolHandler `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefProtocolHandlerConfig,
		func(r *RefProtocolHandlerConfig) ProtocolHandler { return r },
	)
}

func (RefProtocolHandlerConfig) isProtocolHandler() {}
func (RefProtocolHandlerConfig) isComponent()       {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefProtocolHandlerConfig) MarshalJSON() ([]byte, error) {
	type t RefProtocolHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindRefProtocolHandlerConfig, data)
	return data, nil
}

//
// ========= End ref@protocol.Handler type =========

// ========= Begin ref@service.Service type =========
//

const kindRefServiceConfig = `ref@service.Service`

// RefServiceConfig ref@service.Service
type RefServiceConfig struct {
	Name string
	Def  Service `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefServiceConfig,
		func(r *RefServiceConfig) Service { return r },
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
	data = prepend(kindKey, kindRefServiceConfig, data)
	return data, nil
}

//
// ========= End ref@service.Service type =========

// ========= Begin ref@stream.Dialer type =========
//

const kindRefStreamDialerConfig = `ref@stream.Dialer`

// RefStreamDialerConfig ref@stream.Dialer
type RefStreamDialerConfig struct {
	Name string
	Def  StreamDialer `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefStreamDialerConfig,
		func(r *RefStreamDialerConfig) StreamDialer { return r },
	)
}

func (RefStreamDialerConfig) isStreamDialer() {}
func (RefStreamDialerConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t RefStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindRefStreamDialerConfig, data)
	return data, nil
}

//
// ========= End ref@stream.Dialer type =========

// ========= Begin ref@stream.Handler type =========
//

const kindRefStreamHandlerConfig = `ref@stream.Handler`

// RefStreamHandlerConfig ref@stream.Handler
type RefStreamHandlerConfig struct {
	Name string
	Def  StreamHandler `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefStreamHandlerConfig,
		func(r *RefStreamHandlerConfig) StreamHandler { return r },
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
	data = prepend(kindKey, kindRefStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End ref@stream.Handler type =========

// ========= Begin ref@stream.ListenConfig type =========
//

const kindRefStreamListenConfigConfig = `ref@stream.ListenConfig`

// RefStreamListenConfigConfig ref@stream.ListenConfig
type RefStreamListenConfigConfig struct {
	Name string
	Def  StreamListenConfig `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefStreamListenConfigConfig,
		func(r *RefStreamListenConfigConfig) StreamListenConfig { return r },
	)
}

func (RefStreamListenConfigConfig) isStreamListenConfig() {}
func (RefStreamListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RefStreamListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t RefStreamListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindRefStreamListenConfigConfig, data)
	return data, nil
}

//
// ========= End ref@stream.ListenConfig type =========

// ========= Begin ref@tls.TLS type =========
//

const kindRefTLSConfig = `ref@tls.TLS`

// RefTLSConfig ref@tls.TLS
type RefTLSConfig struct {
	Name string
	Def  TLS `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindRefTLSConfig,
		func(r *RefTLSConfig) TLS { return r },
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
	data = prepend(kindKey, kindRefTLSConfig, data)
	return data, nil
}

//
// ========= End ref@tls.TLS type =========

// ========= Begin remove_request_header@net/http.Handler type =========
//

const kindRemoveRequestHeaderNetHTTPHandlerConfig = `remove_request_header@net/http.Handler`

// RemoveRequestHeaderNetHTTPHandlerConfig remove_request_header@net/http.Handler
type RemoveRequestHeaderNetHTTPHandlerConfig struct {
	Key string
}

func init() {
	_ = provider.Register(
		kindRemoveRequestHeaderNetHTTPHandlerConfig,
		func(r *RemoveRequestHeaderNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindRemoveRequestHeaderNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End remove_request_header@net/http.Handler type =========

// ========= Begin remove_response_header@net/http.Handler type =========
//

const kindRemoveResponseHeaderNetHTTPHandlerConfig = `remove_response_header@net/http.Handler`

// RemoveResponseHeaderNetHTTPHandlerConfig remove_response_header@net/http.Handler
type RemoveResponseHeaderNetHTTPHandlerConfig struct {
	Key string
}

func init() {
	_ = provider.Register(
		kindRemoveResponseHeaderNetHTTPHandlerConfig,
		func(r *RemoveResponseHeaderNetHTTPHandlerConfig) HTTPHandler { return r },
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
	data = prepend(kindKey, kindRemoveResponseHeaderNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End remove_response_header@net/http.Handler type =========

// ========= Begin self_signed@tls.TLS type =========
//

const kindSelfSignedTLS = `self_signed@tls.TLS`

// SelfSignedTLS self_signed@tls.TLS
type SelfSignedTLS struct {
}

func init() {
	_ = provider.Register(
		kindSelfSignedTLS,
		func(r *SelfSignedTLS) TLS { return r },
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
	data = prepend(kindKey, kindSelfSignedTLS, data)
	return data, nil
}

//
// ========= End self_signed@tls.TLS type =========

// ========= Begin service@once.Once type =========
//

const kindServiceOnceConfig = `service@once.Once`

// ServiceOnceConfig service@once.Once
type ServiceOnceConfig struct {
	Service Service
}

func init() {
	_ = provider.Register(
		kindServiceOnceConfig,
		func(r *ServiceOnceConfig) Once { return r },
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
	data = prepend(kindKey, kindServiceOnceConfig, data)
	return data, nil
}

//
// ========= End service@once.Once type =========

// ========= Begin stream@service.Service type =========
//

const kindStreamServiceConfig = `stream@service.Service`

// StreamServiceConfig stream@service.Service
type StreamServiceConfig struct {
	Listener          StreamListenConfig
	Handler           StreamHandler
	DisconnectOnClose bool `json:",omitempty"`
}

func init() {
	_ = provider.Register(
		kindStreamServiceConfig,
		func(r *StreamServiceConfig) Service { return r },
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
	data = prepend(kindKey, kindStreamServiceConfig, data)
	return data, nil
}

//
// ========= End stream@service.Service type =========

// ========= Begin strip_prefix@net/http.Handler type =========
//

const kindStripPrefixNetHTTPHandlerConfig = `strip_prefix@net/http.Handler`

// StripPrefixNetHTTPHandlerConfig strip_prefix@net/http.Handler
type StripPrefixNetHTTPHandlerConfig struct {
	Prefix  string
	Handler HTTPHandler
}

func init() {
	_ = provider.Register(
		kindStripPrefixNetHTTPHandlerConfig,
		func(r *StripPrefixNetHTTPHandlerConfig) HTTPHandler { return r },
	)
}

func (StripPrefixNetHTTPHandlerConfig) isHTTPHandler() {}
func (StripPrefixNetHTTPHandlerConfig) isComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StripPrefixNetHTTPHandlerConfig) MarshalJSON() ([]byte, error) {
	type t StripPrefixNetHTTPHandlerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindStripPrefixNetHTTPHandlerConfig, data)
	return data, nil
}

//
// ========= End strip_prefix@net/http.Handler type =========

// ========= Begin tls@stream.Dialer type =========
//

const kindTLSStreamDialerConfig = `tls@stream.Dialer`

// TLSStreamDialerConfig tls@stream.Dialer
type TLSStreamDialerConfig struct {
	Dialer StreamDialer
	TLS    TLS
}

func init() {
	_ = provider.Register(
		kindTLSStreamDialerConfig,
		func(r *TLSStreamDialerConfig) StreamDialer { return r },
	)
}

func (TLSStreamDialerConfig) isStreamDialer() {}
func (TLSStreamDialerConfig) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSStreamDialerConfig) MarshalJSON() ([]byte, error) {
	type t TLSStreamDialerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindTLSStreamDialerConfig, data)
	return data, nil
}

//
// ========= End tls@stream.Dialer type =========

// ========= Begin tls@stream.ListenConfig type =========
//

const kindTLSStreamListenConfigConfig = `tls@stream.ListenConfig`

// TLSStreamListenConfigConfig tls@stream.ListenConfig
type TLSStreamListenConfigConfig struct {
	ListenConfig StreamListenConfig
	TLS          TLS
}

func init() {
	_ = provider.Register(
		kindTLSStreamListenConfigConfig,
		func(r *TLSStreamListenConfigConfig) StreamListenConfig { return r },
	)
}

func (TLSStreamListenConfigConfig) isStreamListenConfig() {}
func (TLSStreamListenConfigConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSStreamListenConfigConfig) MarshalJSON() ([]byte, error) {
	type t TLSStreamListenConfigConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindTLSStreamListenConfigConfig, data)
	return data, nil
}

//
// ========= End tls@stream.ListenConfig type =========

// ========= Begin tls_down@stream.Handler type =========
//

const kindTLSDownStreamHandlerConfig = `tls_down@stream.Handler`

// TLSDownStreamHandlerConfig tls_down@stream.Handler
type TLSDownStreamHandlerConfig struct {
	Handler StreamHandler
	TLS     TLS
}

func init() {
	_ = provider.Register(
		kindTLSDownStreamHandlerConfig,
		func(r *TLSDownStreamHandlerConfig) StreamHandler { return r },
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
	data = prepend(kindKey, kindTLSDownStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End tls_down@stream.Handler type =========

// ========= Begin tls_up@stream.Handler type =========
//

const kindTLSUpStreamHandlerConfig = `tls_up@stream.Handler`

// TLSUpStreamHandlerConfig tls_up@stream.Handler
type TLSUpStreamHandlerConfig struct {
	Handler StreamHandler
	TLS     TLS
}

func init() {
	_ = provider.Register(
		kindTLSUpStreamHandlerConfig,
		func(r *TLSUpStreamHandlerConfig) StreamHandler { return r },
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
	data = prepend(kindKey, kindTLSUpStreamHandlerConfig, data)
	return data, nil
}

//
// ========= End tls_up@stream.Handler type =========

// ========= Begin validation@tls.TLS type =========
//

const kindValidationTLSConfig = `validation@tls.TLS`

// ValidationTLSConfig validation@tls.TLS
type ValidationTLSConfig struct {
	Ca IoReader
}

func init() {
	_ = provider.Register(
		kindValidationTLSConfig,
		func(r *ValidationTLSConfig) TLS { return r },
	)
}

func (ValidationTLSConfig) isTLS()       {}
func (ValidationTLSConfig) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ValidationTLSConfig) MarshalJSON() ([]byte, error) {
	type t ValidationTLSConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindValidationTLSConfig, data)
	return data, nil
}

//
// ========= End validation@tls.TLS type =========

// ========= Begin wait@service.Service type =========
//

const kindWaitService = `wait@service.Service`

// WaitService wait@service.Service
type WaitService struct {
}

func init() {
	_ = provider.Register(
		kindWaitService,
		func(r *WaitService) Service { return r },
	)
}

func (WaitService) isService()   {}
func (WaitService) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m WaitService) MarshalJSON() ([]byte, error) {
	type t WaitService
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	data = prepend(kindKey, kindWaitService, data)
	return data, nil
}

//
// ========= End wait@service.Service type =========

// ========= Begin tls.TLS interface =========
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
// ========= End tls.TLS interface =========

// ========= Begin http.Handler interface =========
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
// ========= End http.Handler interface =========

// ========= Begin once.Once interface =========
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
// ========= End once.Once interface =========

// ========= Begin io.Reader interface =========
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
// ========= End io.Reader interface =========

// ========= Begin io.Writer interface =========
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
// ========= End io.Writer interface =========

// ========= Begin net.Conn interface =========
//

// NetConn net.Conn
type NetConn interface {
	isNetConn()
	Component
}

// RawNetConn is store raw bytes of NetConn
type RawNetConn []byte

func (RawNetConn) isNetConn()   {}
func (RawNetConn) isComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawNetConn) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawNetConn) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawNetConn: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End net.Conn interface =========

// ========= Begin net.PacketConn interface =========
//

// NetPacketConn net.PacketConn
type NetPacketConn interface {
	isNetPacketConn()
	Component
}

// RawNetPacketConn is store raw bytes of NetPacketConn
type RawNetPacketConn []byte

func (RawNetPacketConn) isNetPacketConn() {}
func (RawNetPacketConn) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawNetPacketConn) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawNetPacketConn) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawNetPacketConn: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End net.PacketConn interface =========

// ========= Begin packet.Handler interface =========
//

// PacketHandler packet.Handler
type PacketHandler interface {
	isPacketHandler()
	Component
}

// RawPacketHandler is store raw bytes of PacketHandler
type RawPacketHandler []byte

func (RawPacketHandler) isPacketHandler() {}
func (RawPacketHandler) isComponent()     {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawPacketHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawPacketHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawPacketHandler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End packet.Handler interface =========

// ========= Begin packet.ListenConfig interface =========
//

// PacketListenConfig packet.ListenConfig
type PacketListenConfig interface {
	isPacketListenConfig()
	Component
}

// RawPacketListenConfig is store raw bytes of PacketListenConfig
type RawPacketListenConfig []byte

func (RawPacketListenConfig) isPacketListenConfig() {}
func (RawPacketListenConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawPacketListenConfig) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawPacketListenConfig) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawPacketListenConfig: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End packet.ListenConfig interface =========

// ========= Begin protocol.Handler interface =========
//

// ProtocolHandler protocol.Handler
type ProtocolHandler interface {
	isProtocolHandler()
	Component
}

// RawProtocolHandler is store raw bytes of ProtocolHandler
type RawProtocolHandler []byte

func (RawProtocolHandler) isProtocolHandler() {}
func (RawProtocolHandler) isComponent()       {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawProtocolHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawProtocolHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawProtocolHandler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End protocol.Handler interface =========

// ========= Begin service.Service interface =========
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
// ========= End service.Service interface =========

// ========= Begin stream.Dialer interface =========
//

// StreamDialer stream.Dialer
type StreamDialer interface {
	isStreamDialer()
	Component
}

// RawStreamDialer is store raw bytes of StreamDialer
type RawStreamDialer []byte

func (RawStreamDialer) isStreamDialer() {}
func (RawStreamDialer) isComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawStreamDialer) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawStreamDialer) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawStreamDialer: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End stream.Dialer interface =========

// ========= Begin stream.Handler interface =========
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
// ========= End stream.Handler interface =========

// ========= Begin stream.ListenConfig interface =========
//

// StreamListenConfig stream.ListenConfig
type StreamListenConfig interface {
	isStreamListenConfig()
	Component
}

// RawStreamListenConfig is store raw bytes of StreamListenConfig
type RawStreamListenConfig []byte

func (RawStreamListenConfig) isStreamListenConfig() {}
func (RawStreamListenConfig) isComponent()          {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawStreamListenConfig) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawStreamListenConfig) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawStreamListenConfig: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

//
// ========= End stream.ListenConfig interface =========
