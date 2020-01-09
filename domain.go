package supervisor

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var once sync.Once

func init() {
	rand.Seed(time.Now().UnixNano())
	counter = time.Now().UnixNano()
}

func nextID() int64 {
	return atomic.AddInt64(&counter, 1)
}

type event interface{}

type domain struct {
	mux     sync.Mutex
	nodes   map[int64]*Node
	watchCh chan event
}

func newDomain() *domain {
	var d domain
	once.Do(func() {
		d = domain{
			nodes:   make(map[int64]*Node),
			watchCh: make(chan event),
		}
	})

	return &d
}

func (d *domain) add(id int64, n *Node) {
	d.mux.Lock()
	defer d.mux.Unlock()

	d.nodes[id] = n
}
