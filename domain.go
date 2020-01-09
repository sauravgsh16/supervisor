package supervisor

import (
	"sync"
	"sync/atomic"
)

var once sync.Once

type domain struct {
	wg       sync.WaitGroup
	mux      sync.Mutex
	services map[int64]*service
	done     chan *service
	count    int32
}

func newDomain() *domain {
	var d domain
	once.Do(func() {
		d = domain{
			services: make(map[int64]*service),
			done:     make(chan *service),
		}
	})

	return &d
}

func (d *domain) add(id int64, s *service) {
	d.mux.Lock()
	defer d.mux.Unlock()

	d.services[id] = s
	atomic.AddInt32(&d.count, 1)
}

func (d *domain) num() int32 {
	return atomic.LoadInt32(&d.count)
}

func (d *domain) multiplex(s *service) {
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		select {
		case <-s.done:
			d.done <- s
		}
	}()
}
