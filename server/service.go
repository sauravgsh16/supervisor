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
	domain     *domain
	leaderDone chan interface{}
}

func newNodeService(ch chan interface{}) *nodeService {
	l := make(chan interface{})
	return &nodeService{
		domain:     newDomain(ch, l),
		leaderDone: l,
	}
}

func (s *nodeService) Register(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	id := req.GetNode().GetId()
	n := &nodeCtx{
		req.GetNode(),
		make(chan string),
	}
	s.domain.add(id, n)
	return &RegisterNodeResponse{Result: true}, nil
}

// WatchLeader checks if leader has been configured, returns leader id
func (s *nodeService) WatchLeader(ctx context.Context, req *LeaderStatusRequest) (*LeaderStatusResponse, error) {
	n := s.domain.get(req.Id)
	if n == nil {
		return nil, status.Error(codes.Internal, "failed to find node")
	}
	ticker := time.NewTicker(1 * time.Millisecond)

	var id string
loop:
	for {
		select {
		case <-ticker.C:
			if s.domain.closed {
				return nil, nil
			}
			s.domain.watchCh <- n
		case id = <-n.idCh:
			break loop
		}
	}
	close(n.idCh)
	ticker.Stop()
	return &LeaderStatusResponse{DependentID: id}, nil
}

// WatchMember checks if member has been configured. Sends id of every new member added
func (s *nodeService) WatchMember(req *MemberStatusRequest, stream Supervise_WatchMemberServer) error {
	n := s.domain.get(req.Id)
	if n == nil {
		return status.Error(codes.Internal, "failed to find node")
	}

	ticker := time.NewTicker(1 * time.Millisecond)

loop:
	for {
		select {
		case <-ticker.C:
			s.domain.watchCh <- n
		case id := <-n.idCh:
			stream.Send(&MemberStatusResponse{
				DependentID: id,
			})
		case <-s.leaderDone:
			break loop
		}
	}
	ticker.Stop()
	return nil
}
