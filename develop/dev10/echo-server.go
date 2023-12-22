package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Println(buffer[:n])
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error while starting server :", err.Error())
	}
	defer listen.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection : ", err.Error())
		}

		fmt.Println("Accepted connection from : ", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
