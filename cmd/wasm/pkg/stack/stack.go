package stack

type Stack[T any] struct {
	array []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (arr *Stack[T]) Pop() T {
	var result T
	n := len(arr.array)

	if n != 0 {
		result = arr.array[n-1]
		arr.array = arr.array[:n-1]
	}

	return result
}

func (arr *Stack[T]) GiveElem() T {
	var result T
	n := len(arr.array)

	if n != 0 {
		result = arr.array[n-1]
	}

	return result
}

func (arr *Stack[T]) Append(item T) {
	arr.array = append(arr.array, item)
}

func (arr *Stack[T]) Clear() {
	arr.array = nil
}
