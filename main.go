package main

import (
	"log"
	"os"

	supervisor "github.com/sauravgsh16/supervisor/server"
)

func main() {
	s := supervisor.NewServer("9090")
	if err := s.Run(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
