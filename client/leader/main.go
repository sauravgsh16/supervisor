package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	supervisor "github.com/sauravgsh16/supervisor/server"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	c := supervisor.NewSuperviseClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id string = "a_test_leader"

	registerReq := &supervisor.RegisterNodeRequest{
		Node: &supervisor.Node{
			Id:   id,
			Type: supervisor.Node_Leader,
		},
	}

	resp, err := c.Register(ctx, registerReq)
	if err != nil || !resp.Result {
		log.Fatalf(err.Error())
	}

	waitReq := &supervisor.MemberStatusRequest{
		Id: id,
	}

	done := make(chan interface{})
	resSteam, err := c.WatchMember(context.Background(), waitReq)
	if err != nil {
		log.Fatalf(err.Error())
	}

	r := make(chan *supervisor.MemberStatusResponse)

	go func(done chan interface{}) {
		for {
			resp, err := resSteam.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf(err.Error())
			}
			select {
			case r <- resp:
			case <-done:
				break
			}
		}
	}(done)

	for resp := range r {
		fmt.Printf("%s\n", resp.DependentID)
	}
}
