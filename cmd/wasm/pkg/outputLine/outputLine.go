package outputLine

type OutputLine[T any] struct {
	line []T
}

func NewOutputLine[T any]() OutputLine[T] {
	return OutputLine[T]{}
}

func (out *OutputLine[T]) Clear() {
	out.line = nil
}
