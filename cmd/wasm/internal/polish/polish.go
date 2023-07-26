package polish

func (pol *Polish) CalculateExpression(expression string) float64 {
	pol.clearAll()

	if !pol.CheckCorrect(expression) {
		return 0
	}

	pol.ExpressionToString(expression)
	result := pol.CalculateString()

	return result
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

func (pol *Polish) ExpressionToString(expression string) {
}

func (pol *Polish) giveMePrioritet(s string) byte {
	switch s {
	case `+`, `-`:
		return 2
	case `*`, `/`:
		return 3
	case `)`:
		return 4
	default:
		return 0
	}
}

func (pol *Polish) CalculateString() float64 {
	return 0
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
