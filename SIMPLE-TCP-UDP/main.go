package main

import (
	"bufio"
	"fmt"
	tcpClient "middleware-sockets/tcp/client"
	tcpServer "middleware-sockets/tcp/server"
	tcpMeasure "middleware-sockets/tcp/tcp-measure"
	udpClient "middleware-sockets/udp/client"
	udpServer "middleware-sockets/udp/server"
	udpMeasure "middleware-sockets/udp/udp-measure"
	"middleware-sockets/util"
	"os"
	"strconv"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press 1 for normal options or 2 for time measuring options.")
	input, _ := reader.ReadString('\n')
	input = util.TrimString(input)
	switch input {
	case "1":
		fmt.Println("+-------------------+\n|--------TCP--------|\n+-------------------+\n|1. Start TCP Server|\n|2. Start TCP Client|\n+-------------------+\n\n+-------------------+\n|--------UDP--------|\n+-------------------+\n|3. Start UDP Server|\n|4. Start UPD Client|\n+-------------------+")
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = util.TrimString(input)
		switch input {
		case "1":
			tcpServer.StartServer()
			break
		case "2":
			tcpClient.StartClient()
			break
		case "3":
			udpServer.StartServer()
			break
		case "4":
			udpClient.StartClient()
			break
		}
	case "2":
		fmt.Println("+-------------------+\n|--------TCP--------|\n+-------------------+\n|1. Start TCP Server|\n|2. Start TCP Client|\n+-------------------+\n\n+-------------------+\n|--------UDP--------|\n+-------------------+\n|3. Start UDP Server|\n|4. Start UPD Client|\n+-------------------+")
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = util.TrimString(input)
		switch input {
		case "1":
			tcpMeasure.StartMeasureServer()
			break
		case "2":
			fmt.Println("How many?")
			input, _ := reader.ReadString('\n')
			input = util.TrimString(input)
			k, _ := strconv.Atoi(input)
			for i := 0; i < k; i++ {
				go tcpMeasure.StartMeasureClient()
			}
			break
		case "3":
			udpMeasure.StartServer()
			break
		case "4":
			fmt.Println("How many?")
			input, _ := reader.ReadString('\n')
			input = util.TrimString(input)
			k, _ := strconv.Atoi(input)
			for i := 0; i < k; i++ {
				go udpMeasure.StartClient()
			}
			break
		}
		time.Sleep(2 * time.Second)
	}
}
