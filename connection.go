package supervisor

import (
	"math/rand"
	"net"
	"sync/atomic"
	"time"
)

var counter int64

func init() {
	rand.Seed(time.Now().UnixNano())
	counter = time.Now().UnixNano()
}

func nextID() int64 {
	return atomic.AddInt64(&counter, 1)
}

// Connection struct
type Connection struct {
	id      int64
	server  *Server
	service *service
	network net.Conn
}

// newConnection returns a new connection
func newConnection(s *Server, n net.Conn, t, depends string) *Connection {
	id := nextID()
	c := &Connection{
		id:      id,
		server:  s,
		network: n,
		service: newService(id, t, depends),
	}
	return c
}

func (c *Connection) openConnection() {

}
