package tests

import (
	"testing"
	"webasm/cmd/wasm/internal/polish"
	"webasm/cmd/wasm/pkg/outputLine"
	"webasm/cmd/wasm/pkg/stack"
)

var polska = polish.NewPolish(stack.NewStack[string](), stack.NewStack[float64](), outputLine.NewOutputLine())

func TestCheckCorrect(t *testing.T) {
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

func TestExpressionToString(t *testing.T) {

}

func TestCalculateString(t *testing.T) {

}
