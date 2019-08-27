package tcp_measure

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"middleware-sockets/util"
)

func StartMeasureClient() {
	times := make([]time.Duration, 10000)
	timesFloated := make([]float64, 10000)
	conn, _ := net.Dial("tcp", "localhost:7474")
	defer conn.Close()
	for i := 0; i < 10000; i++ {
		start := time.Now()
		fmt.Fprintf(conn, "1\n")
		bufio.NewReader(conn).ReadString('\n')
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
