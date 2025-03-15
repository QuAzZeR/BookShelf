package main

import (
	"github.com/quazzer/BookShelf/api/api/server"
)

func main() {
	server.NewHTTPServer()
}

func Add(a int, b int) int {
	return a + b
}
