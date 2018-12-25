package main

import (
	"log"
)

func main() {
	log.Println("This is a regular message.")
	log.Fatalln("This is a fatal error.")
	log.Println("End.")
}
