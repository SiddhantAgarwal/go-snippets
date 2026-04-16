package options_pattern

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	testCases := []struct {
		name  string
		setup func() *Server
		check func(s *Server)
	}{
		{
			name: "default, no options",
			setup: func() *Server {
				return NewServer()
			},
			check: func(s *Server) {
				assert.Equal(t, s.Host, "")
				assert.Equal(t, s.Port, 0)
				assert.Equal(t, s.Timeout, time.Duration(0))
			},
		},
		{
			name: "with host",
			setup: func() *Server {
				return NewServer(
					WithHost("127.0.0.1"),
				)
			},
			check: func(s *Server) {
				assert.Equal(t, s.Host, "127.0.0.1")
				assert.Equal(t, s.Port, 0)
				assert.Equal(t, s.Timeout, time.Duration(0))
			},
		},
		{
			name: "with host, port",
			setup: func() *Server {
				return NewServer(
					WithHost("127.0.0.1"),
					WithPort(8000),
				)
			},
			check: func(s *Server) {
				assert.Equal(t, s.Host, "127.0.0.1")
				assert.Equal(t, s.Port, 8000)
				assert.Equal(t, s.Timeout, time.Duration(0))
			},
		},
		{
			name: "with all options",
			setup: func() *Server {
				return NewServer(
					WithHost("127.0.0.1"),
					WithPort(8000),
					WithTimeout(time.Minute),
				)
			},
			check: func(s *Server) {
				assert.Equal(t, s.Host, "127.0.0.1")
				assert.Equal(t, s.Port, 8000)
				assert.Equal(t, s.Timeout, time.Minute)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := tc.setup()

			tc.check(s)
		})
	}
}
