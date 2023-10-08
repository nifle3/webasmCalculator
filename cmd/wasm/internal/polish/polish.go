package polish

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"webasm/cmd/wasm/pkg/expression"
)

var operatorsMap = map[string]interface{}{
	"+": struct{}{},
	"-": struct{}{},
	"/": struct{}{},
	"*": struct{}{},
}

func (pol *Polish) CalculateExpression(expression string) (float64, error) {
	pol.clearAll()

	if err := pol.CheckCorrect(expression); err != nil {
		return 0, err
	}

	pol.expressionToOutlineString(expression)
	result := pol.calculateOutlineString()
	return result, nil
}

func (pol *Polish) clearAll() {
	pol.operatorStack.Clear()
	pol.operandStack.Clear()
	pol.outputLine.Clear()
}

func (pol *Polish) CheckCorrect(outputLine string) error {
	countBrackets := 0
	for index := range outputLine {
		if string(outputLine[index]) == "(" {
			countBrackets++
		} else if string(outputLine[index]) == ")" {
			countBrackets--
		}

		if index == len(outputLine)-1 {
			break
		}

		if string(outputLine[index]) == "/" && string(outputLine[index+1]) == "0" {
			return errors.New(expression.ZeroExpression)
		}

		_, firstIsOperator := operatorsMap[string(outputLine[index])]
		_, secondIsOperator := operatorsMap[string(outputLine[index+1])]

		if firstIsOperator && secondIsOperator {
			return errors.New(expression.OperatorExpression)
		}
	}

	if countBrackets != 0 {
		return errors.New(expression.BracketExpression)
	}

	return nil

}

func (pol *Polish) expressionToOutlineString(expression string) {
	for i := 0; i < len(expression); i++ {
		if unicode.IsDigit(rune(expression[i])) {
			var digitString strings.Builder

			for ; i < len(expression) && unicode.IsDigit(rune(expression[i])); i++ {
				digitString.WriteByte(expression[i])
			}

			pol.outputLine.Enqueue(digitString.String())
		}

		if i == len(expression) {
			break
		}
		current := string(expression[i])

		if current == `(` {
			pol.operatorStack.Append(string(expression[i]))
			continue
		}

		if current == `)` {
			for res := pol.operatorStack.Pop(); res != "("; res = pol.operatorStack.Pop() {
				pol.outputLine.Enqueue(res)
			}
			continue
		}

		for pol.giveMePriority(pol.operatorStack.GiveElem()) >= pol.giveMePriority(current) {
			elem := pol.operatorStack.Pop()
			pol.outputLine.Enqueue(elem)
		}

		pol.operatorStack.Append(current)
	}

	for pol.operatorStack.GiveLen() != 0 {
		pol.outputLine.Enqueue(pol.operatorStack.Pop())
	}
}

func (pol *Polish) giveMePriority(s string) byte {
	switch s {
	case `+`, `-`:
		return 2
	case `*`, `/`:
		return 3
	case `(`:
		return 1
	default:
		return 0
	}
}

func (pol *Polish) calculateOutlineString() float64 {
	for pol.outputLine.GiveLen() != 0 {
		val := pol.outputLine.Dequeue()

		if value, err := strconv.ParseFloat(val, 64); err == nil {
			pol.operandStack.Append(value)
			continue
		}

		a := pol.operandStack.Pop()
		b := pol.operandStack.Pop()

		result := pol.doOperator(a, b, val)
		pol.operandStack.Append(result)
	}

	return pol.operandStack.Pop()
}

func (pol *Polish) doOperator(a, b float64, operator string) float64 {
	switch operator {
	case `+`:
		return b + a
	case `-`:
		return b - a
	case `*`:
		return b * a
	case `/`:
		return b / a
	default:
		return 0
	}
}
