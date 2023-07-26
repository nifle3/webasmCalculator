package main

import (
	"fmt"
	"webasm/cmd/wasm/internal/polish"
	"webasm/cmd/wasm/pkg/queue"
	"webasm/cmd/wasm/pkg/stack"
)

func main() {
	var polska = polish.NewPolish(stack.NewStack[string](), stack.NewStack[float64](), queue.NewQueue[string]())
	result := polska.CalculateExpression("3*(2+3)")
	fmt.Println(result)
	/*
		c := make(chan int)
		fmt.Println("WASM Go Initialized")
		registration.RegistrHandler()
		<-c
	*/
}
