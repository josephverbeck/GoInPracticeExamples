package main

import (
	"log"
	"net"
	"time"
)

func main() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		panic("Failed to connect.")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "Example ", f)
	logger.Println("This is a normal log")
	logger.Panicf("This is a panic")
}
