package udp_measure

import (
	"net"
)

func StartServer() {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:4747")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		_, addr, _ := conn.ReadFromUDP(buffer)
		go conn.WriteTo([]byte("1\n"), addr)
	}
}
