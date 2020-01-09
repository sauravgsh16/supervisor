package supervisor

import "net"

// Server struct
type Server struct {
	conns  map[int64]*Connection
	domain *domain
}

// NewServer returns a new server
func NewServer() *Server {
	return &Server{
		conns: make(map[int64]*Connection),
	}
}

// Open a new connection with the server
func (s *Server) Open(conn net.Conn, istype, depends string) {
	c := newConnection(s, conn, istype, depends)
	s.conns[c.id] = c
	c.openConnection()
}

func (s *Server) handleOutgoing() {

	_ = func(id int64, c *Connection) {

	}

	for range s.conns {

	}
}
