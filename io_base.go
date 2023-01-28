package utils

import "net"

type BaseIO struct {
	net.Conn
}

func (t BaseIO) ReadN(buf []byte, n int) error {
	return ReadN(t.Conn, buf, n)
}
