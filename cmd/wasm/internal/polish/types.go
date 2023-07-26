package polish

type istack[T any] interface {
	Append(item T)
	Pop() T
	GiveElem() T
	Clear()
}

type iline interface {
	Clear()
}

type Polish struct {
	operatorStack istack[string]
	operandStack  istack[float64]
	outputLine    iline
}

func NewPolish(operators istack[string], operands istack[float64], outputline iline) Polish {
	return Polish{
		operatorStack: operators,
		operandStack:  operands,
		outputLine:    outputline,
	}
}
