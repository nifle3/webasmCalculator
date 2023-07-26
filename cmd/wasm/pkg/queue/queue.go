package queue

type Queue[T any] struct {
	line []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Clear() {
	q.line = nil
}

func (q *Queue[T]) Dequeue() T {
	var result T
	n := len(q.line)

	if n != 0 {
		result = q.line[0]
		q.line = q.line[1:]
	}

	return result
}

func (q *Queue[T]) Enqueue(item T) {
	q.line = append(q.line, item)
}

func (q *Queue[T]) GiveLen() int {
	return len(q.line)
}
