package main

import (
	"bufio"
	"fmt"
	tcpClient "middleware-sockets/tcp/client"
	tcpServer "middleware-sockets/tcp/server"
	udpClient "middleware-sockets/udp/client"
	udpServer "middleware-sockets/udp/server"
	"middleware-sockets/util"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("+-------------------+\n|--------TCP--------|\n+-------------------+\n|1. Start TCP Server|\n|2. Start TCP Client|\n+-------------------+\n\n+-------------------+\n|--------UDP--------|\n+-------------------+\n|3. Start UDP Server|\n|4. Start UPD Client|\n+-------------------+")
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	input = util.TrimString(input)
	if input == "1" {
		tcpServer.StartServer()
	} else if input == "2" {
		tcpClient.StartClient()
	} else if input == "3" {
		udpServer.StartServer()
	} else if input == "4" {
		udpClient.StartClient()
	} else {
		fmt.Println("Invalid")
	}
}
