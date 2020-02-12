package main

import (
	"context"
	"fmt"
	"log"
	"sync"
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

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			registerReq := &supervisor.RegisterNodeRequest{
				Node: &supervisor.Node{
					Type: supervisor.Node_Member,
				},
			}

			resp, err := c.Register(ctx, registerReq)
			if err != nil || !resp.Result {
				log.Fatalf(err.Error())
			}

			fmt.Printf("ID is %s\n", resp.Id)

			waitReq := &supervisor.LeaderStatusRequest{
				Id: resp.Id,
			}

			respWatch, err := c.WatchLeader(context.Background(), waitReq)
			if err != nil {
				log.Fatalf(err.Error())
			}

			fmt.Printf("%s\n", respWatch.DependentID)
		}(i)
	}

	wg.Wait()
	fmt.Println("All done")
	/*
		registerReq := &supervisor.RegisterNodeRequest{
			Node: &supervisor.Node{
				Type: supervisor.Node_Member,
				Id:   id,
			},
		}

		resp, err := c.Register(ctx, registerReq)
		if err != nil || !resp.Result {
			log.Fatalf(err.Error())
		}

		waitReq := &supervisor.LeaderStatusRequest{
			Id: id,
		}

		respWatch, err := c.WatchLeader(context.Background(), waitReq)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Printf("%s\n", respWatch.DependentID)
	*/
}
