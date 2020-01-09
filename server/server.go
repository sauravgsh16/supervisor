package supervisor

import (
	"context"
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
	ctx      context.Context
	grpc     *grpc.Server
	listener net.Listener
	done     chan interface{}
}

// NewServer returns a new server
func NewServer(ctx context.Context, port string) Server {
	return &server{
		port: port,
		ctx:  ctx,
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
			s.grpc.GracefulStop()
			s.done <- true
			<-s.ctx.Done()
		}
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

	log.Printf("supervisor server starting on port: %s", s.port)
	return s.run()
}
