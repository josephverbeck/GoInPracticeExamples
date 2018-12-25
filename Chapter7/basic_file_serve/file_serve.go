package main

import (
	"net/http"
)

func main() {
	dir := http.Dir("./..")
	http.ListenAndServe(":8080", http.FileServer(dir))
}
