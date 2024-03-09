package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalln("error happend when connecting to server becuse: ", err)
	}
	defer connection.Close()
	fmt.Printf("Local client connection %v , Remote client connection %v \n", connection.LocalAddr(), connection.RemoteAddr())

	connection.Write([]byte(`ali said "I love fateme"`))

	var dataResponse = make([]byte, 1024)

	_, err = connection.Read(dataResponse)
	if err != nil {
		log.Println("your Response don't reseived becuse: ", err)
	}
	fmt.Println("Response: ", string(dataResponse))
}
