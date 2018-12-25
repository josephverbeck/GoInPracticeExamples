package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside of goroutes")
	go func() {
		fmt.Println("Inside of goroutes")
	}()
	fmt.Println("Outside Again.")

	runtime.Gosched()
}
