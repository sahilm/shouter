package shouter

import (
	"net"
	"time"
)

type Conn struct {
	net.Conn

	IdleTimeout time.Duration
}

func (c *Conn) Write(p []byte) (n int, err error) {
	c.updateDeadline()
	n, err = c.Conn.Write(p)
	return
}

func (c *Conn) Read(b []byte) (n int, err error) {
	c.updateDeadline()
	n, err = c.Conn.Read(b)
	return
}

func (c *Conn) Close() (err error) {
	err = c.Conn.Close()
	return
}

func (c *Conn) updateDeadline() {
	idleDeadline := time.Now().Add(c.IdleTimeout)
	c.Conn.SetDeadline(idleDeadline)
}