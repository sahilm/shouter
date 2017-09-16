package shouter

import (
	"io"
	"net"
	"time"
)

type conn struct {
	net.Conn

	IdleTimeout   time.Duration
	MaxReadBuffer int64
}

func (c *conn) Write(p []byte) (n int, err error) {
	c.updateDeadline()
	n, err = c.Conn.Write(p)
	return
}

func (c *conn) Read(b []byte) (n int, err error) {
	c.updateDeadline()
	r := io.LimitReader(c.Conn, c.MaxReadBuffer)
	n, err = r.Read(b)
	return
}

func (c *conn) Close() (err error) {
	err = c.Conn.Close()
	return
}

func (c *conn) updateDeadline() {
	idleDeadline := time.Now().Add(c.IdleTimeout)
	c.Conn.SetDeadline(idleDeadline)
}
