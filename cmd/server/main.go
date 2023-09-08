package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const (
	port = ":8080"
)

func main() {
	var dir string

	if len(os.Args) > 2 && (os.Args[1] == "--dir" || os.Args[1] == "-D") {
		dir = os.Args[2]
	} else {
		dir, _ = os.Getwd()
		if runtime.GOOS == "windows" {
			dir = strings.TrimSuffix(dir, "\\cmd\\server")
			dir = dir + "\\assets"
		} else {
			dir = strings.TrimSuffix(dir, "/cmd/server")
			dir = dir + "/assets"
		}
	}

	log.Printf("server is listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(dir))))
}
