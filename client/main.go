package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("error happend when connecting to server becuse: ", err)
	}
	fmt.Printf("Local client connection %v , Remote client connection %v \n", connection.LocalAddr(), connection.RemoteAddr())
}
