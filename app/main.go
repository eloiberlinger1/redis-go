package main

import (
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

	for {
		temp := make([]byte, 1024)
		_, err := connection.Read(temp)

		if err != nil {
			if err == io.EOF{
				fmt.Println("Client disconnected:", err)
			} else {
				fmt.Println("An error occured !")
			}
			return 
		}

		connection.Write([]byte("+PONG\r\n"))
	}

}