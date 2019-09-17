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

	ips := make(map[string]int)
	new := 1

	for {
		buffer := make([]byte, 1024)
		//Joga entrada no buffer
		size, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		id, ok := ips[addr.String()]
		if !ok {
			id = new
			ips[addr.String()] = id
			new++
		}
		go handleConnection(conn, buffer, size, addr, id)
	}
}

func handleConnection(conn *net.UDPConn, buffer []byte, size int, addr *net.UDPAddr, id int) {
	//Lê do buffer
	text := util.TrimString(string(buffer[:size]))

	fmt.Printf("Message Received from #%d: %s\n", id, text)
	text = strings.ToUpper(text)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteTo([]byte(text), addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
