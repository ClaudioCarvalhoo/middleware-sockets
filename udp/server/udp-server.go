package server

import (
	"fmt"
	"middleware-sockets/util"
	"net"
	"os"
	"strings"
	"time"
)

const address = "localhost:4747"

func StartServer() {
	addr, _ := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Listening on " + "address")

	buffer := make([]byte, 1024)
	go func() {
		for {
			size, addr, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("disconnected?")
				return
			}

			text := util.TrimString(string(buffer[:size]))

			fmt.Printf("Message Received from %s: %s\n", addr, text)
			text = strings.ToUpper(text)

			conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
			_, err = conn.WriteTo([]byte(text), addr)
			if err != nil {
				fmt.Println("disconnected?")
				return
			}
		}
	}()

	time.Sleep(2 * time.Hour)
}
