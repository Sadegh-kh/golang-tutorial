package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		Network = "tcp"
		Address = "127.0.0.1:8000"
	)
	Listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalln("can't listen to this network becuse : ", err)
	}
	fmt.Println("my ip address listener -> ", Listener.Addr())

	// connection, err := Listener.Accept()
	// if err != nil {
	// 	log.Fatalln("can't accept to this connection becuse : ", err)
	// }

	// fmt.Printf("Local connection: %v , Remote Connection: %v\n", connection.LocalAddr(), connection.RemoteAddr())

}
