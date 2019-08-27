package tcp_measure

import (
	"bufio"
	"io"
	"net"
)

func StartMeasureServer() {
	listener, _ := net.Listen("tcp", "localhost:7474")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			return
		}
		conn.Write([]byte(message + "\n"))
	}
}
