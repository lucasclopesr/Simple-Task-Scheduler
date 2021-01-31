package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/api"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/jobhandler"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/processes"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	errChan := make(chan error, 1)

	s, err := api.NewServer()
	if err != nil {
		panic(err)
	}

	manager := processes.GetProcessManager()
	s.Init(jobhandler.NewJobHandler())

	go func() {
		errChan <- s.Run(ctx, &wg)
	}()

	manager.Run(ctx, &wg)

	done := make(chan bool, 1)

	go func() {
		wg.Wait()
		done <- true
	}()

	for {
		select {
		case err = <-errChan:
			if err != nil {
				cancel()
				log.Println(err)
			}
		case <-done:
			return
		}
	}
}
