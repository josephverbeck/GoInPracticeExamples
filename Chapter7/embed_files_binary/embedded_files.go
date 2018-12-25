package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
)

func main() {
	box := rice.MustFindBox("../files/")
	httpbox := box.HTTPBox()
	http.ListenAndServe(":8080", http.FileServer(httpbox))
}
