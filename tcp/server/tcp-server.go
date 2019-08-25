package server

import (
	"bufio"
	"fmt"
	"middleware-sockets/util"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7474"
	CONN_TYPE = "tcp"
)

func StartServer() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	id := 1
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Accepted connection #%d\n", id)
		// Handle connections in a new goroutine.
		go handleRequest(conn, id)
		id++
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, id int) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		message = util.TrimString(message)
		if err != nil {
			fmt.Printf("Connection #%d disconnected\n", id)
			return
		}
		fmt.Println("Message Received:", message)
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))
	}
}
