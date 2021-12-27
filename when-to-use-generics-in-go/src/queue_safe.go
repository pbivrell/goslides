package main

type Queue [T any]struct {
	items chan []T  // non-empty slices only
	empty chan bool // holds true if the queue is empty
}

func NewQueue[T any]() *Queue[T] {
	items := make(chan []T, 1)
	empty := make(chan bool, 1)
	empty <- true
	return &Queue[T]{items, empty}
}

func (q *Queue[T]) Get() T {
	items := <-q.items
	item := items[0]
	items = items[1:]
	if len(items) == 0 {
		q.empty <- true
	} else {
		q.items <- items
	}
	return item
}

func (q *Queue[T]) Put(item T) {
	var items []T
	select {
	case items = <-q.items:
	case <-q.empty:
	}
	items = append(items, item)
	q.items <- items
}

func main() {
	
}
