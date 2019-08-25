package server

import(
	"bufio"
	"fmt"
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
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	b := bufio.NewReader(conn)
	bytes, err := b.ReadBytes('\n')
	text := strings.TrimRight(string(bytes), "\n")
	if err != nil { // EOF, or worse
		os.Exit(1)
	}
	// Send a response back to person contacting us.
	conn.Write([]byte(text))
	// Close the connection when you're done with it.
	conn.Close()
}
