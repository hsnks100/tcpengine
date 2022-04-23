package tcpengine

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type TcpHandler interface {
	Recv(conn net.Conn, b []byte)
	OnConnect(conn net.Conn)
	OnClose(conn net.Conn)
}
type TcpEngine struct {
	handler  TcpHandler
	listener net.Listener
	bufSize  int
}

func NewTcpEngine(h TcpHandler) *TcpEngine {
	return &TcpEngine{handler: h}
}

func (t *TcpEngine) Listen(port int) {
	var err error
	t.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if nil != err {
		log.Fatalf("fail to bind address; err: %v", err)
	}
	for {
		conn, err := t.listener.Accept()
		if nil != err {
			log.Printf("fail to accept; err: %v", err)
			continue
		}
		if t.bufSize <= 0 {
			t.bufSize = 256
		}
		t.handler.OnConnect(conn)
		go func() {
			reader := bufio.NewReader(conn)
			for {
				buf := make([]byte, t.bufSize)
				n, err := reader.Read(buf)
				if err != nil {
					t.handler.OnClose(conn)
					return
				}
				t.handler.Recv(conn, buf[:n])
			}
		}()
	}
}
