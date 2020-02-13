package supervisor

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sauravgsh16/can-interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// defaultVbsURL = "tcp://172.24.49.16:19000"
	defaultVbsURL = "tcp://localhost:19000"
	hexchars      = "0123456789abcdef"
)

var (
	canMsgCounter int32  = 10
	msgStr        string = "Xtd 1 18EEFF%d 8 %02d 00 23 04 00 81 00 A0"
)

func nextCanID() int32 {
	return atomic.AddInt32(&canMsgCounter, 1)

}

func encode(src byte) []byte {
	dst := make([]byte, 2)
	dst[0] = hexchars[src>>4]
	dst[1] = hexchars[src&0x0F]
	return dst
}

func fromHexChar(ch byte) (byte, bool) {
	switch {
	case '0' <= ch && ch <= '9':
		return ch - '0', true
	case 'a' <= ch && ch <= 'f':
		return ch - 'a' + 10, true
	case 'A' <= ch && ch <= 'F':
		return ch - 'A' + 10, true
	default:
		return 0, false
	}
}

func getHex(i int) byte {

	sb := encode(byte(i))
	a, ok := fromHexChar(sb[0])
	if !ok {
		fmt.Println("Error")
	}
	b, ok := fromHexChar(sb[1])
	if !ok {
		fmt.Println("Error")
	}

	h := a<<4 | b
	return h
}

type nodeService struct {
	domain     *domain
	leaderDone chan interface{}
	wg         sync.WaitGroup
	can        *can.Can
}

func newNodeService(sdone chan interface{}) *nodeService {
	ldone := make(chan interface{})

	in := make(chan can.DataHolder)
	c, err := can.New(defaultVbsURL, in)
	if err != nil {
		log.Fatalf(err.Error())
	}
	c.Init()

	fmt.Println("Connected to can interface")

	return &nodeService{
		domain:     newDomain(sdone, ldone),
		leaderDone: ldone,
		can:        c,
	}
}

func (s *nodeService) Register(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	n := &nodeCtx{
		req.GetNode(),
		make(chan string),
	}
	id := s.domain.add(n)

	if n.Type == Node_Member {
		s.wg.Add(1)
	}

	go func(id string) {

		i, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalf(err.Error())
		}

		hexI := getHex(i)
		hexCounter := getHex(int(nextCanID()))

		select {

		case s.can.Out <- &can.Message{
			ArbitrationID: []uint8{0x18, 0xee, 0xff, hexI},
			Data:          []uint8{hexCounter, 0x00, 0x20, 0x4, 0x0, 0x81, 0x0, 0xa0},
		}:

		}
	}(id)

	return &RegisterNodeResponse{
		Result: true,
		Id:     id,
	}, nil
}

// WatchLeader checks if leader has been configured, returns leader id
func (s *nodeService) WatchLeader(ctx context.Context, req *LeaderStatusRequest) (*LeaderStatusResponse, error) {
	n := s.domain.get(req.Id)
	if n == nil {
		return nil, status.Error(codes.Internal, "failed to find node")
	}
	ticker := time.NewTicker(1 * time.Second)

	var id string
	defer s.wg.Wait()
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
	ticker.Stop()
	return &LeaderStatusResponse{DependentID: id}, nil
}

// WatchMember checks if member has been configured. Sends id of every new member added
func (s *nodeService) WatchMember(req *MemberStatusRequest, stream Supervise_WatchMemberServer) error {
	n := s.domain.get(req.Id)
	if n == nil {
		return status.Error(codes.Internal, "failed to find node")
	}

	ticker := time.NewTicker(1 * time.Second)

loop:
	for {
		select {
		case <-ticker.C:
			s.domain.watchCh <- n
		case id := <-n.idCh:
			if err := stream.Send(&MemberStatusResponse{
				DependentID: id,
			}); err != nil {
				return err
			}
			s.wg.Done()
		case <-s.leaderDone:
			fmt.Println("Exiting loop ....")
			break loop
		}
	}
	ticker.Stop()
	fmt.Println("Closing can connection ....")
	s.can.Close()
	close(s.domain.watchCh)
	return nil
}
