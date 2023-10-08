package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"webasm/cmd/wasm/internal/polish"
	"webasm/cmd/wasm/pkg/expression"
	"webasm/cmd/wasm/pkg/queue"
	"webasm/cmd/wasm/pkg/stack"
)

var (
	polishVariable = polish.NewPolish(stack.NewStack[string](), stack.NewStack[float64](), queue.NewQueue[string]())
)

func TestCheckCorrect(t *testing.T) {
	data := map[string]error{
		"1+2*((2-3)-1)":    nil,
		"1/0":              errors.New(expression.ZeroExpression),
		"":                 nil,
		"1+2*(2-3":         errors.New(expression.BracketExpression),
		"1+2*3*(1+4-(5-3)": errors.New(expression.BracketExpression),
		"1---1-2-2-2--2":   errors.New(expression.OperatorExpression),
	}

	for key, value := range data {
		t.Run("test", func(t *testing.T) {
			result := polishVariable.CheckCorrect(key)
			assert.Equal(t, value, result)
		})
	}
}

func TestFullCalculate(t *testing.T) {
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
		t.Run("testFullCalculate", func(t *testing.T) {
			result, _ := polishVariable.CalculateExpression(key)
			assert.Equal(t, value, result)
		})
	}
}
