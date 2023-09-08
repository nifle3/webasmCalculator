package polish

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func (pol *Polish) CalculateExpression(expression string) (float64, error) {
	pol.clearAll()

	if !pol.CheckCorrect(expression) {
		return 0, errors.New("Делить на 0 нельзя")
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

func (pol *Polish) CheckCorrect(expression string) bool {
	countBrackets := 0
	for index, value := range expression {
		if string(value) == "(" {
			countBrackets++
		} else if string(value) == ")" {
			countBrackets--
		} else if string(value) == "/" && index != len(expression)-1 && string(expression[index+1]) == "0" {
			return false
		}
	}

	if countBrackets == 0 {
		return true
	} else {
		return false
	}

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
