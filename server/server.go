package supervisor

import (
	"errors"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// Server interface
type Server interface {
	Run() error
}

type server struct {
	port     string
	service  *nodeService
	grpc     *grpc.Server
	listener net.Listener
	done     chan interface{}
}

// NewServer returns a new server
func NewServer(port string) Server {
	return &server{
		port: port,
		done: make(chan interface{}),
	}
}

func (s *server) registerService() error {
	s.service = newNodeService(s.done)
	s.grpc = grpc.NewServer()
	if s.service == nil {
		return errors.New("register called before definition")
	}
	RegisterSuperviseServer(s.grpc, s.service)
	return nil
}

func (s *server) startListener() error {
	var err error

	s.listener, err = net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) run() error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		for range ch {
			log.Println("shutting down supervisor server....")
			s.done <- true
			break
		}
		close(s.done)
		s.grpc.GracefulStop()
	}()
	return s.grpc.Serve(s.listener)
}

func (s *server) Run() error {
	if err := s.startListener(); err != nil {
		return err
	}

	if err := s.registerService(); err != nil {
		return err
	}

	log.Printf("supervisor server starting on port: %s\n", s.port)
	return s.run()
}
