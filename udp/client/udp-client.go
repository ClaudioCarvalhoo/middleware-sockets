package client

import (
	"bufio"
	"fmt"
	"middleware-sockets/util"
	"net"
	"os"
	"time"
)

const address = "localhost:4747"

func StartClient() {
	addr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return
	}

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

		buffer := make([]byte, 1024)
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		size, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("só deus sabe")
			return
		}

		received := util.TrimString(string(buffer[:size]))
		fmt.Printf("Message from server in %s: %s\n", addr, received)
	}
}
