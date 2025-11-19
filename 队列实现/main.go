package main

func main() {

}

type queue[T any] struct {
	length      int
	queueStruct []T
}

func (q *queue[T]) Push(item T) {
	q.queueStruct = append(q.queueStruct, item)
	q.length++
}
func (q *queue[T]) Pop() {
	q.queueStruct = q.queueStruct[1:]
	q.length--
}
func (q *queue[T]) Front() T {
	return q.queueStruct[0]
}
func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		length:      0,
		queueStruct: make([]T, 0, 16),
	}
}
