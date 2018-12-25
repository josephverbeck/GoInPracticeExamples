package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()
	logger := log.New(logfile, "Example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a reg message")
	logger.Fatalln("This is a fatal error.")
	logger.Println("End.")
}
