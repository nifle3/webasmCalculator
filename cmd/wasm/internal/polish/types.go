package polish

type istack[T any] interface {
	Append(item T)
	Pop() T
	GiveElem() T
	Clear()
	GiveLen() int
}

type iline[T any] interface {
	Clear()
	Dequeue() T
	Enqueue(item T)
	GiveLen() int
}

type Polish struct {
	operatorStack istack[string]
	operandStack  istack[float64]
	outputLine    iline[string]
}

func NewPolish(operators istack[string], operands istack[float64], outputline iline[string]) *Polish {
	return &Polish{
		operatorStack: operators,
		operandStack:  operands,
		outputLine:    outputline,
	}
}
