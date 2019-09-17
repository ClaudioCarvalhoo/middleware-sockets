package client

import (
	"bufio"
	"fmt"
	"middleware-sockets/util"
	"net"
	"os"
)

const address = "localhost:7474"

func StartClient() {
	conn, _ := net.Dial("tcp", address)
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		text = util.TrimString(text)
		if text == "exit" {
			return
		}
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = util.TrimString(message)
		fmt.Println("Message from server: " + message)
	}
}
