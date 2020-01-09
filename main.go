package main

import (
	"context"
	"log"
	"os"

	supervisor "github.com/sauravgsh16/supervisor/server"
)

func main() {
	ctx := context.Background()
	s := supervisor.NewServer(ctx, "9090")
	if err := s.Run(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
