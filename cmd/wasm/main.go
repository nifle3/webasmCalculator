package main

import (
	"fmt"
	"webasm/cmd/wasm/internal/polish"
	"webasm/cmd/wasm/internal/registration"
	"webasm/cmd/wasm/pkg/queue"
	"webasm/cmd/wasm/pkg/stack"
)

func main() {
	c := make(chan int)
	fmt.Println("WASM Go Initialized")

	polska := polish.NewPolish(stack.NewStack[string](), stack.NewStack[float64](), queue.NewQueue[string]())
	wasm := registration.NewWasmHandler(polska)
	wasm.Handler()
	
	<-c
}
