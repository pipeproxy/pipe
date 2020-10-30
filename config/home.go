package config

import (
	"net/http"

	"github.com/pipeproxy/pipe/bind"
)

func BuildHomeWithHTTPHandler() bind.HTTPHandler {
	return bind.DirectNetHTTPHandlerConfig{
		Code: http.StatusOK,
		Body: bind.InlineIoReaderConfig{
			Data: `<!DOCTYPE html>
<html>
<head><title>Welcome to Pipe</title></head>
<body>
<center><h1>Welcome to Pipe</h1></center>
<hr><center>Pipe</center>
</body>
</html>`,
		},
	}
}
