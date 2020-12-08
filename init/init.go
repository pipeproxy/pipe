// Code generated. DO NOT EDIT!
package init

import (
	_ "github.com/pipeproxy/pipe/components/balance"
	_ "github.com/pipeproxy/pipe/components/balance/none"
	_ "github.com/pipeproxy/pipe/components/balance/random"
	_ "github.com/pipeproxy/pipe/components/balance/round_robin"
	_ "github.com/pipeproxy/pipe/components/common/load"
	_ "github.com/pipeproxy/pipe/components/once"
	_ "github.com/pipeproxy/pipe/components/once/components"
	_ "github.com/pipeproxy/pipe/components/once/message"
	_ "github.com/pipeproxy/pipe/components/once/multi"
	_ "github.com/pipeproxy/pipe/components/once/service"
	_ "github.com/pipeproxy/pipe/components/packet"
	_ "github.com/pipeproxy/pipe/components/packet/handler/http3"
	_ "github.com/pipeproxy/pipe/components/packet/handler/quic"
	_ "github.com/pipeproxy/pipe/components/packet/listener/listener"
	_ "github.com/pipeproxy/pipe/components/protocol"
	_ "github.com/pipeproxy/pipe/components/protocol/http"
	_ "github.com/pipeproxy/pipe/components/protocol/http/add_request_header"
	_ "github.com/pipeproxy/pipe/components/protocol/http/add_response_header"
	_ "github.com/pipeproxy/pipe/components/protocol/http/compress"
	_ "github.com/pipeproxy/pipe/components/protocol/http/config_dump"
	_ "github.com/pipeproxy/pipe/components/protocol/http/direct"
	_ "github.com/pipeproxy/pipe/components/protocol/http/expvar"
	_ "github.com/pipeproxy/pipe/components/protocol/http/file"
	_ "github.com/pipeproxy/pipe/components/protocol/http/forward"
	_ "github.com/pipeproxy/pipe/components/protocol/http/hosts"
	_ "github.com/pipeproxy/pipe/components/protocol/http/lb"
	_ "github.com/pipeproxy/pipe/components/protocol/http/log"
	_ "github.com/pipeproxy/pipe/components/protocol/http/method"
	_ "github.com/pipeproxy/pipe/components/protocol/http/metrics"
	_ "github.com/pipeproxy/pipe/components/protocol/http/multi"
	_ "github.com/pipeproxy/pipe/components/protocol/http/mux"
	_ "github.com/pipeproxy/pipe/components/protocol/http/pprof"
	_ "github.com/pipeproxy/pipe/components/protocol/http/quit"
	_ "github.com/pipeproxy/pipe/components/protocol/http/redirect"
	_ "github.com/pipeproxy/pipe/components/protocol/http/remove_request_header"
	_ "github.com/pipeproxy/pipe/components/protocol/http/remove_response_header"
	_ "github.com/pipeproxy/pipe/components/protocol/http/strip_prefix"
	_ "github.com/pipeproxy/pipe/components/service"
	_ "github.com/pipeproxy/pipe/components/service/multi"
	_ "github.com/pipeproxy/pipe/components/service/packet"
	_ "github.com/pipeproxy/pipe/components/service/stream"
	_ "github.com/pipeproxy/pipe/components/service/tags"
	_ "github.com/pipeproxy/pipe/components/service/wait"
	_ "github.com/pipeproxy/pipe/components/stdio/input"
	_ "github.com/pipeproxy/pipe/components/stdio/input/env"
	_ "github.com/pipeproxy/pipe/components/stdio/input/file"
	_ "github.com/pipeproxy/pipe/components/stdio/input/http"
	_ "github.com/pipeproxy/pipe/components/stdio/input/inline"
	_ "github.com/pipeproxy/pipe/components/stdio/output"
	_ "github.com/pipeproxy/pipe/components/stdio/output/file"
	_ "github.com/pipeproxy/pipe/components/stream"
	_ "github.com/pipeproxy/pipe/components/stream/dialer/dialer"
	_ "github.com/pipeproxy/pipe/components/stream/dialer/lb"
	_ "github.com/pipeproxy/pipe/components/stream/dialer/tls"
	_ "github.com/pipeproxy/pipe/components/stream/handler/forward"
	_ "github.com/pipeproxy/pipe/components/stream/handler/http1"
	_ "github.com/pipeproxy/pipe/components/stream/handler/http2"
	_ "github.com/pipeproxy/pipe/components/stream/handler/lb"
	_ "github.com/pipeproxy/pipe/components/stream/handler/log"
	_ "github.com/pipeproxy/pipe/components/stream/handler/multi"
	_ "github.com/pipeproxy/pipe/components/stream/handler/mux"
	_ "github.com/pipeproxy/pipe/components/stream/handler/mux/pattern"
	_ "github.com/pipeproxy/pipe/components/stream/handler/tls"
	_ "github.com/pipeproxy/pipe/components/stream/listener/listener"
	_ "github.com/pipeproxy/pipe/components/stream/listener/tls"
	_ "github.com/pipeproxy/pipe/components/tls"
	_ "github.com/pipeproxy/pipe/components/tls/acme"
	_ "github.com/pipeproxy/pipe/components/tls/from"
	_ "github.com/pipeproxy/pipe/components/tls/merge"
	_ "github.com/pipeproxy/pipe/components/tls/self_signed"
	_ "github.com/pipeproxy/pipe/components/tls/validation"

	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/balance/policy"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/once/once"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/packet/handler"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/packet/listenconfig"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/protocol/handler"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/service/service"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/stream/dialer"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/stream/handler"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/stream/listenconfig"
	_ "github.com/pipeproxy/pipe/components/common/gen/github.com/pipeproxy/pipe/components/tls/tls"
	_ "github.com/pipeproxy/pipe/components/common/gen/io/reader"
	_ "github.com/pipeproxy/pipe/components/common/gen/io/writer"
	_ "github.com/pipeproxy/pipe/components/common/gen/net/conn"
	_ "github.com/pipeproxy/pipe/components/common/gen/net/http/handler"
	_ "github.com/pipeproxy/pipe/components/common/gen/net/packetconn"
)
