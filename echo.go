package main
 
import (
	"fmt"
	"io"
	"net"
)
 
func main() {
	fmt.Println("Starting the server")
 
	listener, err := net.Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic (err)
		}
		go io.Copy (conn, conn)
	}
}
