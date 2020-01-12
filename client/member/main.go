package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
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

	var id string = "a_test_member"
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int, id string) {
			defer wg.Done()

			s := strconv.Itoa(i)
			id = id + s
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

			respWatch, err := c.WatchLeader(ctx, waitReq)
			if err != nil {
				log.Fatalf(err.Error())
			}

			fmt.Printf("%s\n", respWatch.DependentID)
		}(i, id)
	}

	wg.Wait()
	fmt.Println("All done")
}
