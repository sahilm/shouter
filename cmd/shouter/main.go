package main

import (
	"log"

	"time"

	"github.com/sahilm/shouter"
)

func main() {
	srv := shouter.Server{
		Addr:        ":8080",
		IdleTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
