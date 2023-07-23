package pkg

type Stack[T any] struct {
	array []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (arr *Stack[T]) Pop() T {
	n := len(arr.array) - 1
	element := arr.array[n]
	arr.array = arr.array[:n]
	return element
}

func (arr *Stack[T]) GiveElem() T {
	return arr.array[len(arr.array)-1]
}

func (arr *Stack[T]) Append(items []T) {
	for _, value := range items {
		arr.array = append(arr.array, value)
	}
}

func (arr *Stack[T]) AppendOne(items T) {
	arr.array = append(arr.array, items)
}
