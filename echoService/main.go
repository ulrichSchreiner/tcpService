package main

import (
	"fmt"
	"github.com/ulrichSchreiner/tcpService"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("please specify port\n")
		os.Exit(1)
	}
	p, _ := strconv.Atoi(os.Args[1])
	tcpService.Service(p, func(c net.Conn) {
		io.Copy(c, c)
	})
}
