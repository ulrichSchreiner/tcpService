package tcpService

import (
	"fmt"
	"net"
	"strconv"
)

type ConnectionHandler func(net.Conn)

func Service(port int, h ConnectionHandler) {
	fmt.Printf("Starting the server on port %d\n", port)

	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go h(conn)
	}
}
