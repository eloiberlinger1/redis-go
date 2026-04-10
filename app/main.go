package main

import (
	"bufio"
	"strings"
	"fmt"
	"net"
	"os"
	"io"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	

	defer l.Close()

	for {

		connection, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleConnection(connection)


	}


	
}

func handleConnection(connection net.Conn) {

	defer connection.Close()
	reader := bufio.NewReader(connection)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF{
				fmt.Println("Client disconnected:", err)
			} else {
				fmt.Println("An error occured !")
			}
			return 
		}


		if line[0] == '*' {

			reader.ReadString('\n')
			rawCommand, _ := reader.ReadString('\n')
			command := strings.TrimSpace(rawCommand)

			if strings.ToUpper(command) == "PING"{
				connection.Write([]byte("+PONG\r\n"))
			}

		} else {
			if strings.Contains(strings.ToUpper(command), "PING") {
				connection.Write([]byte("+PONG\r\n"))
			}
		}


	}

}