package supervisor

import (
	"sync"
)

var once sync.Once

type nodeCtx struct {
	*Node
	done chan interface{}
}

type domain struct {
	mux     sync.RWMutex
	leader  map[int64]*nodeCtx
	member  map[int64]*nodeCtx
	doneCh  map[int64]chan interface{}
	watchCh chan *nodeCtx
}

func newDomain() *domain {
	var d domain
	once.Do(func() {
		d = domain{
			leader:  make(map[int64]*nodeCtx),
			member:  make(map[int64]*nodeCtx),
			doneCh:  make(map[int64]chan interface{}),
			watchCh: make(chan *nodeCtx),
		}
	})

	go d.watcher()

	return &d
}

func (d *domain) add(id int64, n *nodeCtx) {
	d.mux.Lock()
	defer d.mux.Unlock()

	switch n.Type {
	case Node_Leader:
		d.leader[id] = n
	case Node_Member:
		d.member[id] = n
	}

	d.doneCh[id] = n.done

	d.watchCh <- n
}

func (d *domain) getdone(id int64) chan interface{} {
	d.mux.RLock()
	defer d.mux.Unlock()

	ch, ok := d.doneCh[id]
	if !ok {
		return nil
	}
	return ch
}

func (d *domain) watcher() {
	for {
		select {
		case n := <-d.watchCh:
			switch n.Type {
			case Node_Leader:
				if len(d.member) <= 0 {
					continue
				}
			case Node_Member:
				if len(d.leader) <= 0 {
					continue
				}
			}
			select {
			case n.done <- true:
			}
		}
	}
}
