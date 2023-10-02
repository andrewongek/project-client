package main

import (
	"fmt"
	"project-client/server"
	"sync"

	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {

	NewServer, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := NewServer.Start()
		if err != nil {
			panic(err)
		}
		fmt.Println("Server running")

	}()

	wg.Wait()

	return nil
}
