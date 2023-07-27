package tests

import (
	"log"
	"testing"
	"webasm/cmd/wasm/internal/polish"
	"webasm/cmd/wasm/pkg/queue"
	"webasm/cmd/wasm/pkg/stack"
)

var (
	polska = polish.NewPolish(stack.NewStack[string](), stack.NewStack[float64](), queue.NewQueue[string]())
)

func TestCheckCorrect(t *testing.T) {
	log.Printf("start TestCheckCorrect\n")
	data := map[string]bool{
		"1+2*((2-3)-1)":    true,
		"1/0":              false,
		"":                 true,
		"1+2*(2-3":         false,
		"1+2*3*(1+4-(5-3)": false,
	}

	for key, value := range data {
		result := polska.CheckCorrect(key)
		if result != value {
			t.Errorf("on %s :expected = %t, received = %t\n", key, value, result)
		}
	}
}

func TestFullCalculate(t *testing.T) {
	log.Printf("start TestFullCalculate\n")
	data := map[string]float64{
		"2+2":                 4,
		"1-1":                 0,
		"1*(2+3)":             5,
		"3*(2+3)":             15,
		"3*(2+3)+3":           18,
		"2+(1+2+10+111+90)-1": 215,
		"3/5":                 0.6,
		"2*(1-2*(2-6/(5-2)))": 2,
	}

	for key, value := range data {
		result := polska.CalculateExpression(key)
		if result != value {
			t.Errorf("on %s: expected = %g, received = %g", key, value, result)
		}
	}
}
