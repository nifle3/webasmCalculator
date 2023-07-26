package main

import (
	"fmt"
	"webasm/cmd/wasm/internal/registration"
)

func main() {
	c := make(chan int)
	fmt.Println("WASM Go Initialized")
	registration.RegistrHandler()
	<-c
}
