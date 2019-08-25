package client

import (
	"bufio"
	"fmt"
	"middleware-sockets/util"
	"net"
	"os"
)

func StartClient() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:7474")
	defer conn.Close()
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		text = util.TrimString(text)
		if text == "exit" {
			return
		}
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = util.TrimString(message)
		fmt.Println("Message from server: "+ message)
	}
}