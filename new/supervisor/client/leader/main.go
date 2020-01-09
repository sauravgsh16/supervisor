package main

import (
	"context"
	"fmt"
	"log"
	"time"

	supervisor "github.com/sauravgsh16/supervisor/server"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127:0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	c := supervisor.NewNodeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	registerReq := &supervisor.RegisterNodeRequest{
		Node: &supervisor.Node{
			Type: supervisor.Node_Leader,
		},
	}

	resp, err := c.Register(ctx, registerReq)

	fmt.Printf("Here %+v\n", resp)

	waitReq := &supervisor.NodeStatusRequest{
		Id: resp.Id,
	}

	t := time.Now()
	respWatch, err := c.Watch(ctx, waitReq)

	now := time.Now()
	fmt.Printf("%v", now.Sub(t))

	fmt.Printf("%t", respWatch.Result)
}
