package shouter_test

import (
	"testing"
	"time"

	"bufio"
	"net"

	"github.com/sahilm/shouter"
)

// These are manual tests

func TestServerProtectsAgaintSlowloris(t *testing.T) {
	srv := shouter.Server{
		Addr:         ":8080",
		IdleTimeout:  5 * time.Second,
		MaxReadBytes: 1000,
	}
	go srv.ListenAndServe()

	time.Sleep(1 * time.Second) // hack to wait for server to start
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		t.Fatal(err)
	}
	// We slowly write to simulate a Slowloris. We should fail in one second
	// because we don't satisfy the application level requirement of sending a complete request (with newlines)
	// within 1 second
	for {
		w := bufio.NewWriter(conn)
		w.WriteString(".")
		time.Sleep(200 * time.Millisecond)
		w.Flush()
	}
}
