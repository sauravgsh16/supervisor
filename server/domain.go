package supervisor

import (
	"log"
	"sync"
	"sync/atomic"
)

var once sync.Once

type nodeCtx struct {
	*Node
	done chan interface{}
}

type domain struct {
	mux     sync.RWMutex
	nodes   map[int64]*nodeCtx
	watchCh chan *nodeCtx
	member  int64
	leader  int64
}

func newDomain(done chan interface{}) *domain {
	var d domain
	once.Do(func() {
		d = domain{
			nodes:   make(map[int64]*nodeCtx),
			watchCh: make(chan *nodeCtx),
		}
	})

	go d.watcher(done)

	return &d
}

func (d *domain) addMember() {
	atomic.AddInt64(&d.member, 1)
}

func (d *domain) addLeader() {
	atomic.AddInt64(&d.leader, 1)
}

func (d *domain) memberexists() bool {
	return atomic.LoadInt64(&d.member) > 0
}

func (d *domain) leaderexists() bool {
	return atomic.LoadInt64(&d.leader) > 0
}

func (d *domain) add(id int64, n *nodeCtx) {
	d.mux.Lock()
	d.nodes[id] = n
	d.mux.Unlock()

	switch n.Type {
	case Node_Leader:
		d.addLeader()
	case Node_Member:
		d.addMember()
	}
}

func (d *domain) get(id int64) *nodeCtx {
	d.mux.RLock()
	defer d.mux.RUnlock()

	n, ok := d.nodes[id]
	if !ok {
		return nil
	}
	return n
}

func (d *domain) watcher(done chan interface{}) {
loop:
	for {
		select {
		case n := <-d.watchCh:
			switch n.Type {
			case Node_Leader:
				if !d.memberexists() {
					continue
				}
			case Node_Member:
				if !d.leaderexists() {
					continue
				}
			}
			select {
			case n.done <- true:
			}
		case <-done:
			log.Printf("closing domain")
			break loop
		}
	}
	close(d.watchCh)
}
