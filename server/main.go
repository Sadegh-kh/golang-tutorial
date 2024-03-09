package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		Network = "tcp"
		Address = "127.0.0.1:9999"
	)
	Listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Println("can't listen to this network becuse : ", err)
	}
	defer Listener.Close()
	fmt.Println("server ip listener -> ", Listener.Addr())

	for {

		connection, err := Listener.Accept()
		if err != nil {
			log.Println("can't accept to this connection becuse : ", err)

			continue
		}
		fmt.Printf("Local connection: %v , Remote Connection: %v\n", connection.LocalAddr(), connection.RemoteAddr())

		var data = make([]byte, 1024)
		_, err = connection.Read(data)
		if err != nil {
			log.Println("can't read data from client becouse: ", err)

			continue
		}

		fmt.Println("data from client: ", string(data))
		connection.Write([]byte(`fateme said "I don't care"`))

	}

}
