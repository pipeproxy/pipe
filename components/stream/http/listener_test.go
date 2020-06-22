package http

import (
	"net"
	"reflect"
	"testing"
)

func TestListener(t *testing.T) {
	p1, p2 := net.Pipe()
	_ = p2
	listen := newSingleConnListener(p1)
	if !reflect.DeepEqual(listen.Addr(), p1.LocalAddr()) {
		t.Fail()
	}
	conn, err := listen.Accept()
	if err != nil {
		t.Fatal(err)
	}
	go p2.Write([]byte{1})
	var buf [1]byte
	conn.Read(buf[:])
	if buf[0] != 1 {
		t.Fail()
	}
	{
		conn, err := listen.Accept()
		if err == nil || conn != nil {
			t.Fail()
		}
	}

	conn.Close()
}
