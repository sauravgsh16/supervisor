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
}

// NewServer returns a new server
func NewServer(ctx context.Context, port string) Server {
	return &server{
		port:    port,
		service: newNodeService(),
		ctx:     ctx,
		grpc:    grpc.NewServer(),
	}
}

func (s *server) registerService() error {
	if s.grpc == nil || s.service == nil {
		return errors.New("register called before definition")
	}
	RegisterNodeServiceServer(s.grpc, s.service)
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
			s.grpc.GracefulStop()
			<-s.ctx.Done()
		}
	}()
	return s.grpc.Serve(s.listener)
}

func (s *server) Run() error {
	if err := s.registerService(); err != nil {
		return err
	}

	if err := s.startListener(); err != nil {
		return err
	}

	log.Printf("supervisor server starting on port: %s", s.port)
	return s.run()
}
