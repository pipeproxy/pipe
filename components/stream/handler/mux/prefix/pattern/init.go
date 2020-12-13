package pattern

import (
	"github.com/pipeproxy/pipe/components/stream/handler/mux/prefix"
)

func init() {
	for _, pattern := range patterns {
		prefix.RegisterRegexp(pattern[0], pattern[1])
	}
}

// RegisterRegexp pattern
var patterns = [][2]string{

	// tls
	// 0       1       2       3       4       5       6       7       8
	// +-------+-------+-------+-------+-------+-------+-------+-------+-------+
	// |record |    version    |                   ...                         |
	// +-------+---------------+---------------+-------------------------------+
	{"tls", "^\u0016\u0003(\u0000|\u0001|\u0002|\u0003)"},

	// socks
	// 0       1       2       3       4       5       6       7       8
	// +-------+-------+-------+-------+-------+-------+-------+-------+-------+
	// |version|command|                       ...                             |
	// +-------+-------+-------------------------------------------------------+
	{"socks4", "^\u0004(\u0001|\u0002)"},
	{"socks5", "^\u0005(\u0001|\u0002|\u0003)"},

	{"http1", "^(GET|HEAD|POST|PUT|PATCH|DELETE|CONNECT|OPTIONS|TRACE) "},
	{"http2", "^PRI \\* HTTP/2\\.0"},
	{"ssh", "^SSH-"},
}
