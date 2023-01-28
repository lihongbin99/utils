package utils

import (
	"io"
	"net"
)

func Tunnel(conn1, conn2 net.Conn) {
	defer conn1.Close()
	defer conn2.Close()

	errCh := make(chan error)
	go Copy(conn1, conn2, errCh)
	go Copy(conn2, conn1, errCh)

	_ = <-errCh
	_ = <-errCh
}

func Copy(dst io.Writer, src io.Reader, errCh chan error) {
	_, err := io.Copy(dst, src)
	errCh <- err
}

func ReadN(conn net.Conn, buf []byte, n int) error {
	readLen := 0
	for readLen < n {
		l, err := conn.Read(buf[readLen:n])
		if err != nil {
			return err
		}
		readLen += l
	}
	return nil
}

func ReadZeroString(conn net.Conn) (string, error) {
	var err error
	var str []byte
	var buf [1]byte

	for {
		_, err = conn.Read(buf[:])
		if err != nil {
			break
		}
		if buf[0] == 0 {
			break
		}
		str = append(str, buf[0])
	}

	return string(str), err
}
