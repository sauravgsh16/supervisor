package supervisor

import (
	"sync"
)

type service struct {
	id        int64
	is        string
	depends   string
	initiated bool
	mux       sync.RWMutex
	done      chan interface{}
}

func newService(id int64, is, d string) *service {
	return &service{
		id:      id,
		is:      is,
		depends: d,
		done:    make(chan interface{}),
	}
}

func (s *service) initiate() {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.initiated = true
}

func (s *service) status() bool {
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.initiated
}
