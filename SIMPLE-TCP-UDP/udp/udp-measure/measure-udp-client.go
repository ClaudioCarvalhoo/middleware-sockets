package udp_measure

import (
	"fmt"
	"net"
	"time"
	"middleware-sockets/util"
)

func StartClient() {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:4747")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	times := make([]time.Duration, 10000)
	timesFloated := make([]float64, 10000)

	for i := 0; i < 10000; i++ {
		buffer := make([]byte, 1024)
		start := time.Now()
		conn.Write([]byte("1\n"))
		conn.ReadFrom(buffer)
		elapsed := time.Since(start)
		times[i] = elapsed
		timesFloated[i] = float64(elapsed.Nanoseconds())
	}
	sum := int64(0)
	for _, t := range times {
		sum += t.Nanoseconds()
	}
	mean := sum / 10000
	std := util.StdDev(timesFloated, float64(mean))
	fmt.Print("Mean: ")
	fmt.Print(mean)
	fmt.Println(" nanoseconds")
	fmt.Print("Stdev: ")
	fmt.Print(std)
	fmt.Println(" nanoseconds")
}
