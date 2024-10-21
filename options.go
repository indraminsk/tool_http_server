package server

import (
	"net"
	"time"
)

type Options func(*Server)

func Addr(host, port string) Options {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort(host, port)
	}
}

func ReadTimeout(readTimeout time.Duration) Options {
	return func(s *Server) {
		s.server.ReadTimeout = readTimeout
	}
}

func WriteTimeout(writeTimeout time.Duration) Options {
	return func(s *Server) {
		s.server.WriteTimeout = writeTimeout
	}
}

func ShutdownTimeout(shutdownTimeout time.Duration) Options {
	return func(s *Server) {
		s.shutdownTimeout = shutdownTimeout
	}
}
