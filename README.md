# tcpengine

Supports writing structured tcp server/client.

# server example

```go
package main

import (
	"fmt"
	"net"

	"github.com/hsnks100/tcpengine"
)

type kHandler struct {
}

func (k *kHandler) Recv(conn net.Conn, b []byte) {
	fmt.Println("recv", b)
}
func (k *kHandler) OnConnect(conn net.Conn) {
	fmt.Println("onConnect")

}
func (k *kHandler) OnClose(conn net.Conn) {
	fmt.Println("onClose")
}
func main() {
	te := tcpengine.NewTcpEngine(&kHandler{})
	te.Listen(3333)
}
```

# client 
not net.