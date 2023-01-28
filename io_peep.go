package utils

import (
	"net"
	"time"
)

type PeepIo struct {
	net.Conn
	readn int
	buf   []byte
}

func (t *PeepIo) PeepN(n int) ([]byte, error) {
	if t.buf == nil {
		t.buf = make([]byte, 0)
	}

	// 读取需要偷窥的数据
	buf := make([]byte, n)
	if err := ReadN(t.Conn, buf, n); err != nil {
		return nil, err
	}

	// 保存偷窥的数据
	t.buf = append(t.buf, buf...)

	return buf, nil
}

func (t *PeepIo) Read(buf []byte) (n int, err error) {
	if t.readn >= len(t.buf) {
		return t.Conn.Read(buf)
	}

	// 返回偷窥的数据
	dataLen := copy(buf, t.buf[t.readn:])
	t.readn += dataLen
	return dataLen, nil
}

func (t *PeepIo) Write(b []byte) (n int, err error) {
	return t.Conn.Write(b)
}

func (t *PeepIo) Close() error {
	return t.Conn.Close()
}

func (t *PeepIo) LocalAddr() net.Addr {
	return t.Conn.LocalAddr()
}

func (t *PeepIo) RemoteAddr() net.Addr {
	return t.Conn.RemoteAddr()
}

func (t *PeepIo) SetDeadline(t1 time.Time) error {
	return t.Conn.SetDeadline(t1)
}

func (t *PeepIo) SetReadDeadline(t1 time.Time) error {
	return t.Conn.SetReadDeadline(t1)
}

func (t *PeepIo) SetWriteDeadline(t1 time.Time) error {
	return t.Conn.SetWriteDeadline(t1)
}
