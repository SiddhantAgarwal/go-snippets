package options_pattern

import "time"

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
}

type Option func(*Server)

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func NewServer(options ...Option) *Server {
	s := &Server{}

	for _, option := range options {
		option(s)
	}

	return s
}
