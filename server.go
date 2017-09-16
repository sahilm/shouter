package shouter

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

type Server struct {
	Addr        string
	IdleTimeout time.Duration
}

func (srv Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("starting server on %v\n", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		newConn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection %v", err)
		}
		log.Printf("accepted connection from %v", newConn.RemoteAddr())

		conn := &Conn{
			Conn:        newConn,
			IdleTimeout: srv.IdleTimeout,
		}
		conn.SetDeadline(time.Now().Add(conn.IdleTimeout))
		go handle(conn)
	}
}

func handle(conn net.Conn) error {
	defer func() {
		log.Printf("closing connection from %v", conn.RemoteAddr())
		conn.Close()
	}()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	scanr := bufio.NewScanner(r)
	for {
		scanned := scanr.Scan()
		if !scanned {
			if err := scanr.Err(); err != nil {
				return err
			}
			break
		}
		w.WriteString(strings.ToUpper(scanr.Text()) + "\n")
		w.Flush()
	}
	return nil
}
