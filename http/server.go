package http

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/wzshiming/pipe"
	"github.com/wzshiming/pipe/stream"
)

type server struct {
	ch       chan net.Conn
	listener *Listener
	handler  http.Handler
	tls      *tls.Config
	once     sync.Once
	svc      *http.Server
}

func NewServer(handler http.Handler, tls *tls.Config) *server {
	s := &server{
		handler:  handler,
		listener: NewListener(),
		tls:      tls,
	}
	s.svc = &http.Server{
		Handler: s,
	}
	return s
}

func (s *server) Close() {
	s.listener.Close()
	s.svc.Close()
}

func (s *server) start() {
	go s.run()
}

func (s *server) run() {
	if s.tls == nil {
		err := s.svc.Serve(s.listener)
		if err != nil {
			log.Println("[ERROR] [http]", err)
		}
	} else {
		s.svc.TLSConfig = s.tls
		err := s.svc.ServeTLS(s.listener, "", "")
		if err != nil {
			log.Println("[ERROR] [http]", err)
		}
	}
}

func (s *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Server", pipe.Name)
	s.handler.ServeHTTP(rw, r)
}

func (s *server) ServeStream(ctx context.Context, stm stream.Stream) {
	s.once.Do(s.start)
	s.listener.Send(stm)
}
