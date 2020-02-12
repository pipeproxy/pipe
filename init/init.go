package init

import (
	_ "github.com/wzshiming/pipe/codec/base32"
	_ "github.com/wzshiming/pipe/codec/base64"
	_ "github.com/wzshiming/pipe/codec/bzip2"
	_ "github.com/wzshiming/pipe/codec/gzip"
	_ "github.com/wzshiming/pipe/codec/hex"
	_ "github.com/wzshiming/pipe/codec/json"
	_ "github.com/wzshiming/pipe/dialer/network"
	_ "github.com/wzshiming/pipe/dialer/poller"
	_ "github.com/wzshiming/pipe/dialer/tls"
	_ "github.com/wzshiming/pipe/http/add_request_header"
	_ "github.com/wzshiming/pipe/http/add_response_header"
	_ "github.com/wzshiming/pipe/http/compress"
	_ "github.com/wzshiming/pipe/http/config_dump"
	_ "github.com/wzshiming/pipe/http/direct"
	_ "github.com/wzshiming/pipe/http/expvar"
	_ "github.com/wzshiming/pipe/http/file"
	_ "github.com/wzshiming/pipe/http/forward"
	_ "github.com/wzshiming/pipe/http/h2c"
	_ "github.com/wzshiming/pipe/http/log"
	_ "github.com/wzshiming/pipe/http/multi"
	_ "github.com/wzshiming/pipe/http/mux"
	_ "github.com/wzshiming/pipe/http/poller"
	_ "github.com/wzshiming/pipe/http/pprof"
	_ "github.com/wzshiming/pipe/http/redirect"
	_ "github.com/wzshiming/pipe/http/remove_request_header"
	_ "github.com/wzshiming/pipe/http/remove_response_header"
	_ "github.com/wzshiming/pipe/http/weighted"
	_ "github.com/wzshiming/pipe/input/file"
	_ "github.com/wzshiming/pipe/input/inline"
	_ "github.com/wzshiming/pipe/listener/multi"
	_ "github.com/wzshiming/pipe/listener/network"
	_ "github.com/wzshiming/pipe/listener/tls"
	_ "github.com/wzshiming/pipe/once/message"
	_ "github.com/wzshiming/pipe/output/file"
	_ "github.com/wzshiming/pipe/service/multi"
	_ "github.com/wzshiming/pipe/service/server"
	_ "github.com/wzshiming/pipe/stream/forward"
	_ "github.com/wzshiming/pipe/stream/http"
	_ "github.com/wzshiming/pipe/stream/multi"
	_ "github.com/wzshiming/pipe/stream/mux"
	_ "github.com/wzshiming/pipe/stream/poller"
	_ "github.com/wzshiming/pipe/stream/tls"
	_ "github.com/wzshiming/pipe/stream/weighted"
	_ "github.com/wzshiming/pipe/tls/acme"
	_ "github.com/wzshiming/pipe/tls/from"
	_ "github.com/wzshiming/pipe/tls/self_signed"

	_ "github.com/wzshiming/pipe/stream/mux/pattern"
)
