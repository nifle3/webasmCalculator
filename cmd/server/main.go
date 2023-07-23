package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
	path = "D:\\webasm\\assets"
)

func main() {
	log.Printf("server is listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(path))))
}
