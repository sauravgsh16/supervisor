package supervisor

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var counter int64

func init() {
	rand.Seed(time.Now().UnixNano())
	counter = time.Now().UnixNano()
}

func nextID() int64 {
	return atomic.AddInt64(&counter, 1)
}

type nodeService struct {
	domain *domain
}

func newNodeService(ch chan interface{}) *nodeService {
	return &nodeService{
		domain: newDomain(ch),
	}
}

func (s *nodeService) Register(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	id := nextID()
	n := &nodeCtx{
		req.GetNode(),
		make(chan interface{}),
	}
	s.domain.add(id, n)
	return &RegisterNodeResponse{Id: id}, nil
}

func (s *nodeService) Watch(ctx context.Context, req *NodeStatusRequest) (*NodeStatusResponse, error) {
	n := s.domain.get(req.Id)
	if n == nil {
		return nil, status.Error(codes.Internal, "failed to find node")
	}
	ticker := time.NewTicker(1 * time.Second)
loop:
	for {
		select {
		case <-ticker.C:
			s.domain.watchCh <- n
		case <-n.done:
			break loop
		}
	}
	ticker.Stop()
	return &NodeStatusResponse{Result: true}, nil
}
