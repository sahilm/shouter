package main

import (
	"time"

	"github.com/sahilm/shouter"
)

func main() {
	srv := shouter.Server{
		Addr:         ":8080",
		IdleTimeout:  10 * time.Second,
		MaxReadBytes: 1000,
	}
	go srv.ListenAndServe()
	time.Sleep(10 * time.Second)
	srv.Shutdown()
}
