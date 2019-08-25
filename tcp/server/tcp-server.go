package server

import (
	"bufio"
	"fmt"
	"io"
	"middleware-sockets/util"
	"net"
	"os"
	"strings"
)

const address = "localhost:7474"

func StartServer() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on " + address)

	id := 1
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Accepted connection #%d\n", id)
		go handleConnection(conn, id)
		id++
	}
}

func handleConnection(conn net.Conn, id int) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		message = util.TrimString(message)
		if err == io.EOF {
			fmt.Printf("Connection #%d disconnected\n", id)
			return
		}
		fmt.Println("Message Received:", message)
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))
	}
}
