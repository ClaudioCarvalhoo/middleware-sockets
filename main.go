package main

import (
	"bufio"
	"fmt"
	tcpClient "middleware-sockets/tcp/client"
	tcpServer "middleware-sockets/tcp/server"
	"middleware-sockets/util"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("1. Start TCP Server\n2. Start TCP Client\n3. Start UDP Server\n4. Start UPD Client")
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	input = util.TrimString(input)
	if input == "1" {
		tcpServer.StartServer()
	} else if input == "2" {
		tcpClient.StartClient()
	} else if input == "3" {
		fmt.Println("TODO: Start UDP Server")
	} else if input == "4" {
		fmt.Println("TODO: Start UDP Client")
	} else {
		fmt.Println("Invalid")
	}
}
