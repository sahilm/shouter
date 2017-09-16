package main

import (
	"log"

	"github.com/sahilm/shouter"
)

func main() {
	srv := shouter.Server{}
	log.Fatal(srv.ListenAndServe())
}
