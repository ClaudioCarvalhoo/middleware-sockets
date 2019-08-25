package client

import "net"
import "fmt"
import "bufio"
import "os"

func StartClient() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:3333")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message from server: "+message)
	}
}