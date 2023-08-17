package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/zuri03/KubernetesTest/router"
)

func main() {
	fmt.Println("Creating API Server...")

	//set up router
	fmt.Printf("\tIntializing Router\n")
	requestWaitGroup := new(sync.WaitGroup)
	apiRouter := router.New(requestWaitGroup)

	//set up server
	fmt.Printf("\tInitializing Server\n")
	server := &http.Server{
		Addr:        ":9000",
		Handler:     apiRouter,
		IdleTimeout: 60 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Unable to start server due to error, %s\n", err.Error())
			return
		}
	}()

	fmt.Printf("\tServer listening on port 9000\n")

	signaler := make(chan os.Signal)
	signal.Notify(signaler, os.Interrupt)
	signal.Notify(signaler, os.Kill)

	<-signaler

	fmt.Println("Shutdonw signal recieved, shutting down server")
	requestWaitGroup.Wait()

	fmt.Println("Exiting...")
}
