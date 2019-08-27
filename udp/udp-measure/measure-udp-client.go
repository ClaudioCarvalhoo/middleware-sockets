package udp_measure

import (
	"fmt"
	"net"
	"time"
)

func StartClient() {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:4747")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	times := make([]time.Duration, 10000)

	for i := 0; i < 10000; i++ {
		buffer := make([]byte, 1024)
		start := time.Now()
		conn.Write([]byte("1\n"))
		conn.ReadFrom(buffer)
		elapsed := time.Since(start)
		times[i] = elapsed
	}
	sum := int64(0)
	for _, t := range times {
		sum += t.Nanoseconds()
	}
	mean := sum / 10000
	fmt.Print(mean)
	fmt.Println(" nanoseconds")
}
