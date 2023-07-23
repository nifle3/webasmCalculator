package main

import (
	"strconv"
	"webasm/cmd/wasm/pkg"
)

func main() {
	ArrayToString("1+2+3+4+5+6-2/3/1")
}

func ArrayToString(expression string) string {
	result := ""
	stack := pkg.NewStack[string]()

	for _, value := range expression {
		s := string(value)
		_, err := strconv.ParseFloat(s, 64)
		if err != nil {
			result += s
			continue
		}
		
	}

	return result
}

func Calculate(polska string) float64 {
	return 0
}
