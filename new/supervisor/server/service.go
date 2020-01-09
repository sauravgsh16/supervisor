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

type nodeService struct {
	domain *domain
}

func init() {
	rand.Seed(time.Now().UnixNano())
	counter = time.Now().UnixNano()
}

func nextID() int64 {
	return atomic.AddInt64(&counter, 1)
}

func newNodeService() *nodeService {
	return &nodeService{
		domain: newDomain(),
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
	done := s.domain.getdone(req.Id)
	if done == nil {
		return nil, status.Error(codes.Internal, "failed to find done channel")
	}
loop:
	for {
		select {
		case <-done:
			break loop
		}
	}
	return &NodeStatusResponse{Result: true}, nil
}
