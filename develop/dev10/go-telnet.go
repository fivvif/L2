package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Timeout for connection")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	address := net.JoinHostPort(host, port)

	// Подключение к серверу с таймаутом
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		log.Fatalln("Error connecting to server : ", err.Error())
	}
	fmt.Println("Connected to ", conn.RemoteAddr())

	defer conn.Close()

	// Горутина для чтения данных из сокета и вывода их в STDOUT
	go func() {
		readBuffer := make([]byte, 1024)
		for {
			i, err := conn.Read(readBuffer)
			if err != nil {
				log.Fatalln("Server closed connection :", err.Error())
			}
			fmt.Print("<- " + string(readBuffer[:i]))
		}
	}()

	// Запись данных в сокет
	writeBuffer := make([]byte, 1024)
	for {
		i, err := os.Stdin.Read(writeBuffer)
		if err != nil {
			conn.Close()
			fmt.Println("Closing connection...")
			return
		}
		_, err = conn.Write(writeBuffer[:i])
		if err != nil {
			log.Fatalln("Error while writing to server : ", err.Error())
		}
	}
}
