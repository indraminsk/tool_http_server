package server

import (
	"context"
	"net"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second

	_defaultHost = ""
	_defaultPort = ":8000"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, opts ...Options) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		Addr:         net.JoinHostPort(_defaultHost, _defaultPort),
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
	}

	server := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(server)
	}

	server.start()

	return server
}

func (r *Server) start() {
	go func() {
		r.notify <- r.server.ListenAndServe()

		close(r.notify)
	}()
}

func (r *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), r.shutdownTimeout)
	defer cancel()

	return r.server.Shutdown(ctx)
}

func (r *Server) Notify() <-chan error {
	return r.notify
}
