package supervisor

import (
	"log"
	"sync"
	"sync/atomic"
)

var once sync.Once

type nodeCtx struct {
	*Node
	idCh chan string
}

type domain struct {
	mux        sync.RWMutex
	nodes      map[string]*nodeCtx
	watchCh    chan *nodeCtx
	ldone      chan interface{}
	member     int64
	prevMember int64
	leaderID   string
	memberID   []string
	closed     bool
}

func newDomain(done, ldone chan interface{}) *domain {
	var d domain
	once.Do(func() {
		d = domain{
			nodes:    make(map[string]*nodeCtx),
			watchCh:  make(chan *nodeCtx),
			memberID: make([]string, 0),
			ldone:    ldone,
		}
	})

	go d.watcher(done)

	return &d
}

func (d *domain) addMember() {
	atomic.AddInt64(&d.member, 1)
}

func (d *domain) addprev() {
	atomic.AddInt64(&d.prevMember, 1)
}

func (d *domain) memberCount() int64 {
	return atomic.LoadInt64(&d.member)
}

func (d *domain) prevCount() int64 {
	return atomic.LoadInt64(&d.prevMember)
}

func (d *domain) add(id string, n *nodeCtx) {
	d.mux.Lock()
	defer d.mux.Unlock()

	d.nodes[id] = n

	switch n.Type {
	case Node_Leader:
		d.leaderID = id
	case Node_Member:
		d.addMember()
		d.memberID = append(d.memberID, id)
	}
}

func (d *domain) get(id string) *nodeCtx {
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
				if d.memberCount() > d.prevCount() {
					var i int64
					for i = 0; i < (d.memberCount() - d.prevCount()); i++ {
						select {
						case n.idCh <- d.memberID[d.prevCount()]:
						}
						d.addprev()
					}
				}
			case Node_Member:
				if d.leaderID != "" {
					select {
					case n.idCh <- d.leaderID:
					}
					close(n.idCh)
				}
			}
		case <-done:
			log.Printf("closing domain")
			select {
			case d.ldone <- true:
			default:
			}
			break loop
		}
	}
	close(d.watchCh)
	d.closed = true
}
