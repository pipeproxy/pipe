package init

import (
	_ "github.com/wzshiming/pipe/codec/json"
	_ "github.com/wzshiming/pipe/http"
	_ "github.com/wzshiming/pipe/http/add_request_header"
	_ "github.com/wzshiming/pipe/http/add_response_header"
	_ "github.com/wzshiming/pipe/http/direct"
	_ "github.com/wzshiming/pipe/http/file"
	_ "github.com/wzshiming/pipe/http/multi"
	_ "github.com/wzshiming/pipe/http/mux"
	_ "github.com/wzshiming/pipe/http/remove_request_header"
	_ "github.com/wzshiming/pipe/http/remove_response_header"
	_ "github.com/wzshiming/pipe/listener/network"
	_ "github.com/wzshiming/pipe/service/multi"
	_ "github.com/wzshiming/pipe/service/server"
	_ "github.com/wzshiming/pipe/stream/close"
	_ "github.com/wzshiming/pipe/stream/forward"
	_ "github.com/wzshiming/pipe/stream/multi"
	_ "github.com/wzshiming/pipe/stream/mux"
	_ "github.com/wzshiming/pipe/tls/from_file"
	_ "github.com/wzshiming/pipe/tls/self_signed"

	_ "github.com/wzshiming/pipe/stream/mux/pattern"
)
